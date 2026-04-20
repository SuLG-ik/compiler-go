package main

import (
	"strings"
	"testing"
)

func runSemantic(t *testing.T, src string) *AnalyzerResult {
	t.Helper()
	app := NewApp()
	return app.RunSemanticAnalyzer(src)
}

func hasSemanticError(errors []AnalyzerError, key string) bool {
	for _, e := range errors {
		if e.MessageKey == key {
			return true
		}
	}
	return false
}

func TestSemanticValidFunction(t *testing.T) {
	res := runSemantic(t, "fun calc(a: Int, b: Int, c: Int): Int {\nreturn a + (b * c)\n};")
	if res.OutputKey != "semantic.success" {
		t.Fatalf("expected semantic.success, got %q (errors=%v)", res.OutputKey, res.Errors)
	}
	if len(res.Errors) != 0 {
		t.Fatalf("expected no semantic errors, got %d", len(res.Errors))
	}
	ast := res.OutputParams["ast"]
	if !strings.Contains(ast, "FunctionDeclNode") || !strings.Contains(ast, `name: "calc"`) {
		t.Fatalf("unexpected AST output:\n%s", ast)
	}
	if !strings.Contains(ast, "BinaryOpNode") {
		t.Fatalf("expected BinaryOpNode in AST:\n%s", ast)
	}
}

func TestSemanticDuplicateParam(t *testing.T) {
	res := runSemantic(t, "fun f(a: Int, a: Int): Int {\nreturn a\n};")
	if res.OutputKey != "semantic.failed" {
		t.Fatalf("expected semantic.failed, got %q", res.OutputKey)
	}
	if !hasSemanticError(res.Errors, "semantic.error.duplicateParam") {
		t.Fatalf("expected duplicateParam, got %+v", res.Errors)
	}
}

func TestSemanticUndeclaredIdent(t *testing.T) {
	res := runSemantic(t, "fun f(a: Int): Int {\nreturn a + b\n};")
	if !hasSemanticError(res.Errors, "semantic.error.undeclaredIdent") {
		t.Fatalf("expected undeclaredIdent, got %+v", res.Errors)
	}
}

func TestSemanticReturnTypeMismatch(t *testing.T) {
	res := runSemantic(t, "fun f(a: Float): Int {\nreturn a\n};")
	if !hasSemanticError(res.Errors, "semantic.error.returnTypeMismatch") {
		t.Fatalf("expected returnTypeMismatch, got %+v", res.Errors)
	}
}

func TestSemanticBooleanInArith(t *testing.T) {
	res := runSemantic(t, "fun f(a: Boolean, b: Int): Int {\nreturn a + b\n};")
	if !hasSemanticError(res.Errors, "semantic.error.nonNumericOperand") {
		t.Fatalf("expected nonNumericOperand, got %+v", res.Errors)
	}
}

func TestSemanticIntOverflow(t *testing.T) {
	res := runSemantic(t, "fun f(): Int {\nreturn 9999999999999999\n};")
	if !hasSemanticError(res.Errors, "semantic.error.intRange") {
		t.Fatalf("expected intRange, got %+v", res.Errors)
	}
}

func TestSemanticNumericPromotion(t *testing.T) {
	res := runSemantic(t, "fun mix(a: Int, b: Float): Float {\nreturn a + b\n};")
	if res.OutputKey != "semantic.success" {
		t.Fatalf("expected semantic.success, got %q (errors=%+v)", res.OutputKey, res.Errors)
	}
}

func TestSemanticParserFailureSkipsAst(t *testing.T) {
	res := runSemantic(t, "fun f(): Int {\nreturn\n};")
	if res.OutputKey != "semantic.parserFailed" {
		t.Fatalf("expected semantic.parserFailed, got %q", res.OutputKey)
	}
	if res.OutputParams["ast"] != "" {
		t.Fatalf("expected empty AST, got %q", res.OutputParams["ast"])
	}
}
