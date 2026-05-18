package main

import (
	"fmt"
	"strconv"
	"strings"
)

type expressionTokenKind int

const (
	expressionTokenEOF expressionTokenKind = iota
	expressionTokenNumber
	expressionTokenIdentifier
	expressionTokenPlus
	expressionTokenMinus
	expressionTokenMultiply
	expressionTokenDivide
	expressionTokenModulo
	expressionTokenLParen
	expressionTokenRParen
)

type expressionToken struct {
	kind   expressionTokenKind
	lexeme string
	line   int
	col    int
}

type expressionNodeKind int

const (
	expressionNodeNumber expressionNodeKind = iota
	expressionNodeIdentifier
	expressionNodeBinary
)

type expressionNode struct {
	kind  expressionNodeKind
	token expressionToken
	left  *expressionNode
	right *expressionNode
}

type expressionQuad struct {
	Op     string
	Arg1   string
	Arg2   string
	Result string
}

type expressionParser struct {
	tokens []expressionToken
	pos    int
	errors []AnalyzerError
}

func (a *App) RunPoliz(content string) *AnalyzerResult {
	if strings.TrimSpace(content) == "" {
		return &AnalyzerResult{
			OutputKey: "poliz.emptyOutput",
			Errors: []AnalyzerError{
				{Line: 0, Col: 0, MessageKey: "poliz.emptyMessage"},
			},
			Tokens: []Token{},
		}
	}

	tokens, lexerErrors := tokenizeExpression(content)
	if len(lexerErrors) > 0 {
		return &AnalyzerResult{
			OutputKey: "poliz.lexerFailed",
			Errors:    lexerErrors,
			Tokens:    []Token{},
		}
	}

	parser := &expressionParser{tokens: tokens}
	root := parser.parse()
	if len(parser.errors) > 0 || root == nil {
		return &AnalyzerResult{
			OutputKey: "poliz.parserFailed",
			Errors:    parser.errors,
			Tokens:    []Token{},
		}
	}

	quads := buildExpressionQuads(root)
	poliz := buildExpressionPoliz(root)
	value, integerOnly, divisionByZero := evaluateExpression(root)

	outputKey := "poliz.successValue"
	outputParams := map[string]string{
		"expr":  strings.TrimSpace(content),
		"quads": formatExpressionQuads(quads),
		"poliz": strings.Join(poliz, " "),
		"value": strconv.FormatInt(value, 10),
	}

	if !integerOnly {
		outputKey = "poliz.successNoEval"
		delete(outputParams, "value")
	} else if divisionByZero {
		outputKey = "poliz.successDivZero"
		delete(outputParams, "value")
	}

	return &AnalyzerResult{
		OutputKey:    outputKey,
		OutputParams: outputParams,
		Errors:       []AnalyzerError{},
		Tokens:       []Token{},
	}
}

func tokenizeExpression(content string) ([]expressionToken, []AnalyzerError) {
	runes := []rune(content)
	tokens := make([]expressionToken, 0, len(runes)+1)
	errors := make([]AnalyzerError, 0)
	line := 1
	col := 1
	pos := 0

	for pos < len(runes) {
		ch := runes[pos]

		switch ch {
		case ' ', '\t', '\r':
			pos++
			col++
			continue
		case '\n':
			pos++
			line++
			col = 1
			continue
		}

		startLine, startCol := line, col

		if isExpressionDigit(ch) {
			start := pos
			for pos < len(runes) && isExpressionDigit(runes[pos]) {
				pos++
				col++
			}
			tokens = append(tokens, expressionToken{
				kind:   expressionTokenNumber,
				lexeme: string(runes[start:pos]),
				line:   startLine,
				col:    startCol,
			})
			continue
		}

		if isExpressionLetter(ch) {
			start := pos
			for pos < len(runes) && (isExpressionLetter(runes[pos]) || isExpressionDigit(runes[pos]) || runes[pos] == '_') {
				pos++
				col++
			}
			tokens = append(tokens, expressionToken{
				kind:   expressionTokenIdentifier,
				lexeme: string(runes[start:pos]),
				line:   startLine,
				col:    startCol,
			})
			continue
		}

		kind, ok := singleCharExpressionToken(ch)
		if ok {
			tokens = append(tokens, expressionToken{
				kind:   kind,
				lexeme: string(ch),
				line:   startLine,
				col:    startCol,
			})
			pos++
			col++
			continue
		}

		errors = append(errors, AnalyzerError{
			Line:          startLine,
			Col:           startCol,
			Fragment:      string(ch),
			MessageKey:    "poliz.error.unexpectedChar",
			MessageParams: map[string]string{"ch": string(ch)},
		})
		pos++
		col++
	}

	tokens = append(tokens, expressionToken{kind: expressionTokenEOF, line: line, col: col})
	return tokens, errors
}

func singleCharExpressionToken(ch rune) (expressionTokenKind, bool) {
	switch ch {
	case '+':
		return expressionTokenPlus, true
	case '-':
		return expressionTokenMinus, true
	case '*':
		return expressionTokenMultiply, true
	case '/':
		return expressionTokenDivide, true
	case '%':
		return expressionTokenModulo, true
	case '(':
		return expressionTokenLParen, true
	case ')':
		return expressionTokenRParen, true
	default:
		return expressionTokenEOF, false
	}
}

func isExpressionLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isExpressionDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func (p *expressionParser) parse() *expressionNode {
	root := p.parseExpression()

	for p.current().kind != expressionTokenEOF {
		tok := p.current()
		if tok.kind == expressionTokenRParen {
			p.errors = append(p.errors, AnalyzerError{
				Line:       tok.line,
				Col:        tok.col,
				Fragment:   tok.lexeme,
				MessageKey: "poliz.error.unexpectedClosingParen",
			})
		} else {
			p.errors = append(p.errors, AnalyzerError{
				Line:       tok.line,
				Col:        tok.col,
				Fragment:   tok.lexeme,
				MessageKey: "poliz.error.unexpectedToken",
			})
		}
		p.pos++
	}

	if len(p.errors) > 0 {
		return nil
	}
	return root
}

func (p *expressionParser) parseExpression() *expressionNode {
	left := p.parseTerm()
	for {
		kind := p.current().kind
		if kind != expressionTokenPlus && kind != expressionTokenMinus {
			break
		}

		op := p.advance()
		right := p.parseTerm()
		if left == nil || right == nil {
			return left
		}
		left = &expressionNode{kind: expressionNodeBinary, token: op, left: left, right: right}
	}
	return left
}

func (p *expressionParser) parseTerm() *expressionNode {
	left := p.parseFactor()
	for {
		kind := p.current().kind
		if kind != expressionTokenMultiply && kind != expressionTokenDivide && kind != expressionTokenModulo {
			break
		}

		op := p.advance()
		right := p.parseFactor()
		if left == nil || right == nil {
			return left
		}
		left = &expressionNode{kind: expressionNodeBinary, token: op, left: left, right: right}
	}
	return left
}

func (p *expressionParser) parseFactor() *expressionNode {
	tok := p.current()

	switch tok.kind {
	case expressionTokenNumber:
		p.advance()
		return &expressionNode{kind: expressionNodeNumber, token: tok}
	case expressionTokenIdentifier:
		p.advance()
		return &expressionNode{kind: expressionNodeIdentifier, token: tok}
	case expressionTokenLParen:
		p.advance()
		node := p.parseExpression()
		if p.current().kind == expressionTokenRParen {
			p.advance()
		} else {
			p.errors = append(p.errors, AnalyzerError{
				Line:       tok.line,
				Col:        tok.col,
				Fragment:   tok.lexeme,
				MessageKey: "poliz.error.expectedClosingParen",
			})
		}
		return node
	case expressionTokenEOF:
		p.errors = append(p.errors, AnalyzerError{
			Line:       tok.line,
			Col:        tok.col,
			MessageKey: "poliz.error.expectedOperand",
		})
		return nil
	case expressionTokenRParen:
		p.errors = append(p.errors, AnalyzerError{
			Line:       tok.line,
			Col:        tok.col,
			Fragment:   tok.lexeme,
			MessageKey: "poliz.error.expectedOperand",
		})
		return nil
	default:
		p.errors = append(p.errors, AnalyzerError{
			Line:       tok.line,
			Col:        tok.col,
			Fragment:   tok.lexeme,
			MessageKey: "poliz.error.expectedOperand",
		})
		p.advance()
		return p.parseFactor()
	}
}

func (p *expressionParser) current() expressionToken {
	if p.pos >= len(p.tokens) {
		return p.tokens[len(p.tokens)-1]
	}
	return p.tokens[p.pos]
}

func (p *expressionParser) advance() expressionToken {
	tok := p.current()
	if p.pos < len(p.tokens)-1 {
		p.pos++
	}
	return tok
}

func buildExpressionQuads(root *expressionNode) []expressionQuad {
	quads := make([]expressionQuad, 0)
	counter := 0
	buildExpressionQuadValue(root, &quads, &counter)
	return quads
}

func buildExpressionQuadValue(node *expressionNode, quads *[]expressionQuad, counter *int) string {
	if node == nil {
		return ""
	}

	if node.kind == expressionNodeNumber || node.kind == expressionNodeIdentifier {
		return node.token.lexeme
	}

	left := buildExpressionQuadValue(node.left, quads, counter)
	right := buildExpressionQuadValue(node.right, quads, counter)
	*counter++
	result := fmt.Sprintf("t%d", *counter)
	*quads = append(*quads, expressionQuad{
		Op:     node.token.lexeme,
		Arg1:   left,
		Arg2:   right,
		Result: result,
	})
	return result
}

func buildExpressionPoliz(root *expressionNode) []string {
	if root == nil {
		return nil
	}
	output := make([]string, 0)
	appendExpressionPoliz(root, &output)
	return output
}

func appendExpressionPoliz(node *expressionNode, output *[]string) {
	if node == nil {
		return
	}

	appendExpressionPoliz(node.left, output)
	appendExpressionPoliz(node.right, output)
	*output = append(*output, node.token.lexeme)
}

func formatExpressionQuads(quads []expressionQuad) string {
	if len(quads) == 0 {
		return "—"
	}

	lines := make([]string, 0, len(quads))
	for i, quad := range quads {
		lines = append(lines, fmt.Sprintf("%d. (%s, %s, %s, %s)", i+1, quad.Op, quad.Arg1, quad.Arg2, quad.Result))
	}
	return strings.Join(lines, "\n")
}

func evaluateExpression(root *expressionNode) (int64, bool, bool) {
	if root == nil {
		return 0, false, false
	}

	switch root.kind {
	case expressionNodeIdentifier:
		return 0, false, false
	case expressionNodeNumber:
		value, err := strconv.ParseInt(root.token.lexeme, 10, 64)
		if err != nil {
			return 0, false, false
		}
		return value, true, false
	default:
		left, leftOK, leftDivZero := evaluateExpression(root.left)
		right, rightOK, rightDivZero := evaluateExpression(root.right)
		if leftDivZero || rightDivZero {
			return 0, true, true
		}
		if !leftOK || !rightOK {
			return 0, false, false
		}

		switch root.token.kind {
		case expressionTokenPlus:
			return left + right, true, false
		case expressionTokenMinus:
			return left - right, true, false
		case expressionTokenMultiply:
			return left * right, true, false
		case expressionTokenDivide:
			if right == 0 {
				return 0, true, true
			}
			return left / right, true, false
		case expressionTokenModulo:
			if right == 0 {
				return 0, true, true
			}
			return left % right, true, false
		default:
			return 0, false, false
		}
	}
}
