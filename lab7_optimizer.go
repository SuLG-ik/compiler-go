package main

import (
	"fmt"
	"strconv"
	"strings"
)

type lab7ExprKind int

const (
	lab7ExprIdentifier lab7ExprKind = iota
	lab7ExprLiteral
	lab7ExprBinary
)

type lab7Expr struct {
	kind        lab7ExprKind
	value       string
	op          string
	left, right *lab7Expr
}

type lab7Program struct {
	source string
	tokens []Token
	ast    *AstNode
	expr   *lab7Expr
}

type lab7TacInstr struct {
	op     string
	arg1   string
	arg2   string
	result string
}

func (a *App) RunLab7Optimizer(content string) *AnalyzerResult {
	program, failed := buildLab7Program(content)
	if failed != nil {
		return failed
	}

	folded, foldedChanged := lab7FoldConstants(program.expr)
	simplified, simplifiedChanged := lab7SimplifyNeutral(program.expr)
	optimized, optimizedChanged := lab7SimplifyNeutral(folded)
	optimizedChanged = optimizedChanged || foldedChanged

	return &AnalyzerResult{
		OutputKey: "lab7.success",
		OutputParams: map[string]string{
			"source":            program.source,
			"ast":               PrintAst(program.ast),
			"inputIR":           lab7FormatTAC(lab7GenerateTAC(program.expr)),
			"foldOutputIR":      lab7FormatTAC(lab7GenerateTAC(folded)),
			"foldChanged":       lab7BoolString(foldedChanged),
			"foldStatus":        lab7ChangeStatus(foldedChanged),
			"foldNote":          lab7ChangeNote(foldedChanged, "В текущем выражении нет операций, где оба операнда являются константами."),
			"neutralOutputIR":   lab7FormatTAC(lab7GenerateTAC(simplified)),
			"neutralChanged":    lab7BoolString(simplifiedChanged),
			"neutralStatus":     lab7ChangeStatus(simplifiedChanged),
			"neutralNote":       lab7ChangeNote(simplifiedChanged, "В текущем выражении нет нейтральных операций вида x + 0, x * 1 или x / 1."),
			"optimizedOutputIR": lab7FormatTAC(lab7GenerateTAC(optimized)),
			"optimizedChanged":  lab7BoolString(optimizedChanged),
			"optimizedStatus":   lab7ChangeStatus(optimizedChanged),
		},
		Errors: []AnalyzerError{},
		Tokens: program.tokens,
	}
}

func buildLab7Program(content string) (*lab7Program, *AnalyzerResult) {
	if strings.TrimSpace(content) == "" {
		return nil, &AnalyzerResult{
			OutputKey: "lab7.emptyOutput",
			Errors: []AnalyzerError{
				{Line: 0, Col: 0, MessageKey: "lab7.emptyMessage"},
			},
			Tokens: []Token{},
		}
	}

	lex := Tokenize(content)
	if len(lex.Errors) > 0 {
		errs := make([]AnalyzerError, 0, len(lex.Errors))
		for _, e := range lex.Errors {
			errs = append(errs, AnalyzerError{
				Line:       e.Line,
				Col:        e.Col,
				Fragment:   e.Message,
				MessageKey: "lexer.error.unexpectedChar",
			})
		}
		return nil, &AnalyzerResult{OutputKey: "lab7.lexerFailed", Errors: errs, Tokens: lex.Tokens}
	}

	parser := NewParser(lex.Tokens)
	parserErrors := parser.Parse()
	if len(parserErrors) > 0 {
		errs := make([]AnalyzerError, 0, len(parserErrors))
		for _, e := range parserErrors {
			errs = append(errs, AnalyzerError{
				Line:       e.Line,
				Col:        e.Col,
				Fragment:   e.Fragment,
				MessageKey: e.MessageKey,
			})
		}
		return nil, &AnalyzerResult{OutputKey: "lab7.parserFailed", Errors: errs, Tokens: lex.Tokens}
	}

	builder := NewAstBuilder(lex.Tokens)
	root, semanticErrors := builder.Build()
	if len(semanticErrors) > 0 {
		return nil, &AnalyzerResult{
			OutputKey:    "lab7.semanticFailed",
			OutputParams: map[string]string{"count": strconv.Itoa(len(semanticErrors))},
			Errors:       semanticErrors,
			Tokens:       lex.Tokens,
		}
	}

	expr := lab7ReturnExpr(root)
	if expr == nil {
		return nil, &AnalyzerResult{
			OutputKey: "lab7.noExpression",
			Errors: []AnalyzerError{
				{Line: 0, Col: 0, MessageKey: "lab7.noExpression"},
			},
			Tokens: lex.Tokens,
		}
	}

	return &lab7Program{
		source: strings.TrimSpace(content),
		tokens: lex.Tokens,
		ast:    root,
		expr:   expr,
	}, nil
}

func lab7ReturnExpr(root *AstNode) *lab7Expr {
	body := lab7Child(root, "body")
	retExpr := lab7Child(body, "expr")
	return lab7ExprFromAst(retExpr)
}

func lab7ExprFromAst(node *AstNode) *lab7Expr {
	if node == nil {
		return nil
	}

	switch node.Kind {
	case "IdentifierNode":
		return &lab7Expr{kind: lab7ExprIdentifier, value: lab7Attr(node, "name")}
	case "IntLiteralNode", "FloatLiteralNode":
		return &lab7Expr{kind: lab7ExprLiteral, value: lab7Attr(node, "value")}
	case "BinaryOpNode":
		return &lab7Expr{
			kind:  lab7ExprBinary,
			op:    lab7Attr(node, "op"),
			left:  lab7ExprFromAst(lab7Child(node, "left")),
			right: lab7ExprFromAst(lab7Child(node, "right")),
		}
	}

	return nil
}

func lab7Attr(node *AstNode, label string) string {
	if node == nil {
		return ""
	}
	for _, field := range node.Fields {
		if field.Node == nil && field.Label == label {
			if unquoted, err := strconv.Unquote(field.Value); err == nil {
				return unquoted
			}
			return field.Value
		}
	}
	return ""
}

func lab7Child(node *AstNode, label string) *AstNode {
	if node == nil {
		return nil
	}
	for _, field := range node.Fields {
		if field.Node != nil && field.Label == label {
			return field.Node
		}
	}
	return nil
}

func lab7GenerateTAC(expr *lab7Expr) []lab7TacInstr {
	instructions := make([]lab7TacInstr, 0)
	counter := 0
	value := lab7EmitTAC(expr, &instructions, &counter)
	instructions = append(instructions, lab7TacInstr{op: "return", arg1: value})
	return instructions
}

func lab7EmitTAC(expr *lab7Expr, instructions *[]lab7TacInstr, counter *int) string {
	if expr == nil {
		return ""
	}

	if expr.kind == lab7ExprIdentifier || expr.kind == lab7ExprLiteral {
		return expr.value
	}

	left := lab7EmitTAC(expr.left, instructions, counter)
	right := lab7EmitTAC(expr.right, instructions, counter)
	*counter++
	result := fmt.Sprintf("t%d", *counter)
	*instructions = append(*instructions, lab7TacInstr{
		op:     expr.op,
		arg1:   left,
		arg2:   right,
		result: result,
	})
	return result
}

func lab7FormatTAC(instructions []lab7TacInstr) string {
	if len(instructions) == 0 {
		return "—"
	}

	lines := make([]string, 0, len(instructions))
	for i, instruction := range instructions {
		if instruction.op == "return" {
			lines = append(lines, fmt.Sprintf("%d. return %s", i+1, instruction.arg1))
			continue
		}
		lines = append(lines, fmt.Sprintf("%d. %s = %s %s %s", i+1, instruction.result, instruction.arg1, instruction.op, instruction.arg2))
	}
	return strings.Join(lines, "\n")
}

func lab7BoolString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}

func lab7ChangeStatus(changed bool) string {
	if changed {
		return "IR изменён"
	}
	return "Без изменений"
}

func lab7ChangeNote(changed bool, unchanged string) string {
	if changed {
		return "Оптимизация упростила выражение без изменения результата вычисления."
	}
	return unchanged
}

func lab7FoldConstants(expr *lab7Expr) (*lab7Expr, bool) {
	if expr == nil {
		return nil, false
	}
	if expr.kind != lab7ExprBinary {
		return lab7Clone(expr), false
	}

	left, leftChanged := lab7FoldConstants(expr.left)
	right, rightChanged := lab7FoldConstants(expr.right)
	if value, ok := lab7FoldBinary(expr.op, left, right); ok {
		return &lab7Expr{kind: lab7ExprLiteral, value: value}, true
	}

	return &lab7Expr{kind: lab7ExprBinary, op: expr.op, left: left, right: right}, leftChanged || rightChanged
}

func lab7FoldBinary(op string, left, right *lab7Expr) (string, bool) {
	leftValue, leftOK := lab7IntLiteral(left)
	rightValue, rightOK := lab7IntLiteral(right)
	if !leftOK || !rightOK {
		return "", false
	}

	switch op {
	case "+":
		return strconv.FormatInt(leftValue+rightValue, 10), true
	case "-":
		return strconv.FormatInt(leftValue-rightValue, 10), true
	case "*":
		return strconv.FormatInt(leftValue*rightValue, 10), true
	case "/":
		if rightValue == 0 {
			return "", false
		}
		return strconv.FormatInt(leftValue/rightValue, 10), true
	}

	return "", false
}

func lab7SimplifyNeutral(expr *lab7Expr) (*lab7Expr, bool) {
	if expr == nil {
		return nil, false
	}
	if expr.kind != lab7ExprBinary {
		return lab7Clone(expr), false
	}

	left, leftChanged := lab7SimplifyNeutral(expr.left)
	right, rightChanged := lab7SimplifyNeutral(expr.right)
	changed := leftChanged || rightChanged

	switch expr.op {
	case "+":
		if lab7IsZero(left) {
			return right, true
		}
		if lab7IsZero(right) {
			return left, true
		}
	case "-":
		if lab7IsZero(right) {
			return left, true
		}
	case "*":
		if lab7IsZero(left) || lab7IsZero(right) {
			return &lab7Expr{kind: lab7ExprLiteral, value: "0"}, true
		}
		if lab7IsOne(left) {
			return right, true
		}
		if lab7IsOne(right) {
			return left, true
		}
	case "/":
		if lab7IsOne(right) {
			return left, true
		}
	}

	return &lab7Expr{kind: lab7ExprBinary, op: expr.op, left: left, right: right}, changed
}

func lab7Clone(expr *lab7Expr) *lab7Expr {
	if expr == nil {
		return nil
	}
	return &lab7Expr{
		kind:  expr.kind,
		value: expr.value,
		op:    expr.op,
		left:  lab7Clone(expr.left),
		right: lab7Clone(expr.right),
	}
}

func lab7IntLiteral(expr *lab7Expr) (int64, bool) {
	if expr == nil || expr.kind != lab7ExprLiteral {
		return 0, false
	}
	value, err := strconv.ParseInt(expr.value, 10, 64)
	return value, err == nil
}

func lab7IsZero(expr *lab7Expr) bool {
	value, ok := lab7IntLiteral(expr)
	return ok && value == 0
}

func lab7IsOne(expr *lab7Expr) bool {
	value, ok := lab7IntLiteral(expr)
	return ok && value == 1
}
