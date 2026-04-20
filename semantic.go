package main

import (
	"math"
	"strconv"
)

type SymbolKind int

const (
	SymbolParam SymbolKind = iota
	SymbolFunction
)

type Symbol struct {
	Name string
	Type string
	Kind SymbolKind
	Line int
	Col  int
}

type SymbolTable struct {
	parent  *SymbolTable
	symbols map[string]*Symbol
	order   []string
}

func NewSymbolTable(parent *SymbolTable) *SymbolTable {
	return &SymbolTable{parent: parent, symbols: make(map[string]*Symbol)}
}

func (s *SymbolTable) Declare(sym *Symbol) (existing *Symbol, ok bool) {
	if prev, exists := s.symbols[sym.Name]; exists {
		return prev, false
	}
	s.symbols[sym.Name] = sym
	s.order = append(s.order, sym.Name)
	return nil, true
}

func (s *SymbolTable) Lookup(name string) *Symbol {
	if sym, ok := s.symbols[name]; ok {
		return sym
	}
	if s.parent != nil {
		return s.parent.Lookup(name)
	}
	return nil
}

type AstBuilder struct {
	tokens  []Token
	pos     int
	symbols *SymbolTable
	errors  []AnalyzerError
}

func NewAstBuilder(tokens []Token) *AstBuilder {
	filtered := make([]Token, 0, len(tokens))
	for _, t := range tokens {
		if t.Code != CodeSpace {
			filtered = append(filtered, t)
		}
	}
	return &AstBuilder{tokens: filtered, symbols: NewSymbolTable(nil)}
}

func (b *AstBuilder) peek() int {
	if b.pos < len(b.tokens) {
		return b.tokens[b.pos].Code
	}
	return codeEOF
}

func (b *AstBuilder) advance() Token {
	if b.pos >= len(b.tokens) {
		return Token{Code: codeEOF}
	}
	tok := b.tokens[b.pos]
	b.pos++
	return tok
}

func (b *AstBuilder) addError(line, col int, fragment, key string, params map[string]string) {
	b.errors = append(b.errors, AnalyzerError{
		Line:          line,
		Col:           col,
		Fragment:      fragment,
		MessageKey:    key,
		MessageParams: params,
	})
}

func (b *AstBuilder) Build() (*AstNode, []AnalyzerError) {
	if len(b.tokens) == 0 {
		return nil, b.errors
	}
	root := b.parseFunction()
	return root, b.errors
}

func (b *AstBuilder) parseFunction() *AstNode {
	startTok := b.tokens[b.pos]
	b.advance()

	nameTok := b.advance()
	node := newNode("FunctionDeclNode", startTok.Line, startTok.StartPos)
	node.addAttr("name", quoted(nameTok.Lexeme))

	b.advance()

	paramsNode := newNode("ParamListNode", nameTok.Line, nameTok.StartPos)
	paramScope := NewSymbolTable(b.symbols)
	if b.peek() != CodeRParen {
		b.parseParamList(paramsNode, paramScope)
	}
	b.symbols = paramScope
	node.addChild("params", paramsNode)

	b.advance()
	b.advance()

	retTypeTok := b.advance()
	retTypeNode := newNode("TypeNode", retTypeTok.Line, retTypeTok.StartPos)
	retTypeNode.addAttr("name", quoted(retTypeTok.Lexeme))
	node.addChild("returnType", retTypeNode)

	b.advance()

	bodyNode := b.parseReturnStmt(retTypeTok.Lexeme)
	node.addChild("body", bodyNode)

	b.advance()
	return node
}

func (b *AstBuilder) parseParamList(parent *AstNode, scope *SymbolTable) {
	b.parseParam(parent, scope)
	for b.peek() == CodeComma {
		b.advance()
		b.parseParam(parent, scope)
	}
}

func (b *AstBuilder) parseParam(parent *AstNode, scope *SymbolTable) {
	nameTok := b.advance()
	b.advance()
	typeTok := b.advance()

	param := newNode("ParamNode", nameTok.Line, nameTok.StartPos)
	param.addAttr("name", quoted(nameTok.Lexeme))
	typeNode := newNode("TypeNode", typeTok.Line, typeTok.StartPos)
	typeNode.addAttr("name", quoted(typeTok.Lexeme))
	param.addChild("type", typeNode)

	sym := &Symbol{
		Name: nameTok.Lexeme,
		Type: typeTok.Lexeme,
		Kind: SymbolParam,
		Line: nameTok.Line,
		Col:  nameTok.StartPos,
	}
	if _, ok := scope.Declare(sym); !ok {
		b.addError(nameTok.Line, nameTok.StartPos, nameTok.Lexeme,
			"semantic.error.duplicateParam", nil)
	}

	parent.addItem(param)
}

func (b *AstBuilder) parseReturnStmt(declaredReturnType string) *AstNode {
	retTok := b.advance()
	stmt := newNode("ReturnStmtNode", retTok.Line, retTok.StartPos)

	expr, exprType := b.parseExpr()
	stmt.addChild("expr", expr)

	if exprType != "" && declaredReturnType != "" && !typeAssignable(exprType, declaredReturnType) {
		line, col := exprPosition(expr, retTok)
		b.addError(line, col, exprType,
			"semantic.error.returnTypeMismatch",
			map[string]string{"expected": declaredReturnType, "got": exprType})
	}
	return stmt
}

func exprPosition(node *AstNode, fallback Token) (int, int) {
	if node != nil && node.Line > 0 {
		return node.Line, node.Col
	}
	return fallback.Line, fallback.StartPos
}

func (b *AstBuilder) parseExpr() (*AstNode, string) {
	left, leftType := b.parseTerm()
	for b.peek() == CodePlus || b.peek() == CodeMinus {
		opTok := b.advance()
		right, rightType := b.parseTerm()
		op := opTok.Lexeme
		bin := newNode("BinaryOpNode", opTok.Line, opTok.StartPos)
		bin.addAttr("op", quoted(op))
		bin.addChild("left", left)
		bin.addChild("right", right)
		leftType = b.combineTypes(leftType, rightType, op, opTok)
		left = bin
	}
	return left, leftType
}

func (b *AstBuilder) parseTerm() (*AstNode, string) {
	left, leftType := b.parseFactor()
	for b.peek() == CodeMultiply || b.peek() == CodeDivide {
		opTok := b.advance()
		right, rightType := b.parseFactor()
		op := opTok.Lexeme
		bin := newNode("BinaryOpNode", opTok.Line, opTok.StartPos)
		bin.addAttr("op", quoted(op))
		bin.addChild("left", left)
		bin.addChild("right", right)
		leftType = b.combineTypes(leftType, rightType, op, opTok)
		left = bin
	}
	return left, leftType
}

func (b *AstBuilder) parseFactor() (*AstNode, string) {
	tok := b.tokens[b.pos]
	switch tok.Code {
	case CodeIdent:
		b.advance()
		node := newNode("IdentifierNode", tok.Line, tok.StartPos)
		node.addAttr("name", quoted(tok.Lexeme))
		sym := b.symbols.Lookup(tok.Lexeme)
		if sym == nil {
			b.addError(tok.Line, tok.StartPos, tok.Lexeme,
				"semantic.error.undeclaredIdent", nil)
			return node, ""
		}
		return node, sym.Type
	case CodeInteger:
		b.advance()
		node := newNode("IntLiteralNode", tok.Line, tok.StartPos)
		node.addAttr("value", tok.Lexeme)
		if _, err := strconv.ParseInt(tok.Lexeme, 10, 32); err != nil {
			b.addError(tok.Line, tok.StartPos, tok.Lexeme,
				"semantic.error.intRange", nil)
		}
		return node, "Int"
	case CodeFloatNum:
		b.advance()
		node := newNode("FloatLiteralNode", tok.Line, tok.StartPos)
		node.addAttr("value", tok.Lexeme)
		v, err := strconv.ParseFloat(tok.Lexeme, 64)
		if err != nil || math.IsInf(v, 0) || float64(float32(v)) != v && math.Abs(v) > math.MaxFloat32 {
			b.addError(tok.Line, tok.StartPos, tok.Lexeme,
				"semantic.error.floatRange", nil)
		}
		return node, "Float"
	case CodeLParen:
		b.advance()
		expr, t := b.parseExpr()
		if b.peek() == CodeRParen {
			b.advance()
		}
		return expr, t
	}
	return nil, ""
}

func (b *AstBuilder) combineTypes(left, right, op string, opTok Token) string {
	if left == "" || right == "" {
		return ""
	}
	if !isNumericType(left) {
		b.addError(opTok.Line, opTok.StartPos, left,
			"semantic.error.nonNumericOperand",
			map[string]string{"type": left, "op": op})
		return ""
	}
	if !isNumericType(right) {
		b.addError(opTok.Line, opTok.StartPos, right,
			"semantic.error.nonNumericOperand",
			map[string]string{"type": right, "op": op})
		return ""
	}
	return promoteNumericTypes(left, right)
}

func isNumericType(t string) bool {
	switch t {
	case "Int", "Float", "Double":
		return true
	}
	return false
}

func numericRank(t string) int {
	switch t {
	case "Int":
		return 1
	case "Float":
		return 2
	case "Double":
		return 3
	}
	return 0
}

func promoteNumericTypes(a, b string) string {
	if numericRank(a) >= numericRank(b) {
		return a
	}
	return b
}

func typeAssignable(src, dst string) bool {
	if src == dst {
		return true
	}
	if isNumericType(src) && isNumericType(dst) {
		return numericRank(src) <= numericRank(dst)
	}
	return false
}
