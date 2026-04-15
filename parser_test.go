package main

import (
	"testing"
)

func hasMessage(errors []ParserError, key string) bool {
	for _, err := range errors {
		if err.MessageKey == key {
			return true
		}
	}
	return false
}

func countMessage(errors []ParserError, key string) int {
	count := 0
	for _, err := range errors {
		if err.MessageKey == key {
			count++
		}
	}
	return count
}

func TestParserValidFunction(t *testing.T) {
	input := "fun calc(a: Int): Int {\nreturn a + b * c\n};"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if len(errors) != 0 {
		for _, e := range errors {
			t.Errorf("unexpected error at line %d, col %d: %s (fragment: %q)",
				e.Line, e.Col, e.MessageKey, e.Fragment)
		}
	}
}

func TestParserValidMultipleParams(t *testing.T) {
	input := "fun add(a: Int, b: Float): Double {\nreturn a + b\n};"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if len(errors) != 0 {
		for _, e := range errors {
			t.Errorf("unexpected error at line %d, col %d: %s (fragment: %q)",
				e.Line, e.Col, e.MessageKey, e.Fragment)
		}
	}
}

func TestParserValidNoParams(t *testing.T) {
	input := "fun f(): Boolean {\nreturn x\n};"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if len(errors) != 0 {
		for _, e := range errors {
			t.Errorf("unexpected error at line %d, col %d: %s (fragment: %q)",
				e.Line, e.Col, e.MessageKey, e.Fragment)
		}
	}
}

func TestParserValidParenExpr(t *testing.T) {
	input := "fun f(x: Int): Int {\nreturn (a + b) * c\n};"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if len(errors) != 0 {
		for _, e := range errors {
			t.Errorf("unexpected error at line %d, col %d: %s (fragment: %q)",
				e.Line, e.Col, e.MessageKey, e.Fragment)
		}
	}
}

func TestParserMissingSemicolon(t *testing.T) {
	input := "fun f(): Int {\nreturn a\n}"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if len(errors) != 1 {
		t.Fatalf("expected 1 error, got %d", len(errors))
	}
	if errors[0].MessageKey != "parser.error.expectedSemi" {
		t.Errorf("expected parser.error.expectedSemi, got %s", errors[0].MessageKey)
	}
}

func TestParserMissingReturnExpr(t *testing.T) {
	input := "fun f(): Int {\nreturn\n};"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if len(errors) < 1 {
		t.Fatal("expected at least 1 error for missing expression after return")
	}
	found := false
	for _, e := range errors {
		if e.MessageKey == "parser.error.expectedExpr" {
			found = true
		}
	}
	if !found {
		t.Error("expected parser.error.expectedExpr error")
	}
}

func TestParserMissingOperand(t *testing.T) {
	input := "fun f(): Int {\nreturn a +\n};"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if len(errors) < 1 {
		t.Fatal("expected at least 1 error for missing operand after +")
	}
	found := false
	for _, e := range errors {
		if e.MessageKey == "parser.error.expectedExpr" {
			found = true
		}
	}
	if !found {
		t.Error("expected parser.error.expectedExpr error")
	}
}

func TestParserMissingFunKeyword(t *testing.T) {
	input := "calc(a: Int): Int {\nreturn a\n};"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if len(errors) < 1 {
		t.Fatal("expected at least 1 error for missing 'fun'")
	}
	if errors[0].MessageKey != "parser.error.expectedFun" {
		t.Errorf("expected parser.error.expectedFun, got %s", errors[0].MessageKey)
	}
}

func TestParserEmptyBody(t *testing.T) {
	input := "fun f(): Int {\n};"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if len(errors) < 1 {
		t.Fatal("expected at least 1 error for empty body")
	}
	found := false
	for _, e := range errors {
		if e.MessageKey == "parser.error.expectedReturn" {
			found = true
		}
	}
	if !found {
		t.Error("expected parser.error.expectedReturn error")
	}
}

func TestParserEmptyInput(t *testing.T) {
	result := Tokenize("")
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if len(errors) != 0 {
		t.Errorf("expected 0 errors for empty input, got %d", len(errors))
	}
}

func TestParserComplexExpression(t *testing.T) {
	input := "fun f(a: Int, b: Int, c: Float): Double {\nreturn a * b + c / (a - b)\n};"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if len(errors) != 0 {
		for _, e := range errors {
			t.Errorf("unexpected error at line %d, col %d: %s (fragment: %q)",
				e.Line, e.Col, e.MessageKey, e.Fragment)
		}
	}
}

func TestParserRecoversBeforeReturnInBody(t *testing.T) {
	input := "fun f(): Int {\nfoo bar return a\n};"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if len(errors) != 1 {
		t.Fatalf("expected 1 error, got %d", len(errors))
	}
	if !hasMessage(errors, "parser.error.expectedReturn") {
		t.Fatal("expected parser.error.expectedReturn")
	}
	if hasMessage(errors, "parser.error.expectedSemi") {
		t.Fatal("parser did not recover to program terminator")
	}
}

func TestParserRecoversAfterTrailingCommaInParams(t *testing.T) {
	input := "fun f(a: Int,): Int {\nreturn a\n};"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if len(errors) != 1 {
		t.Fatalf("expected 1 error, got %d", len(errors))
	}
	if !hasMessage(errors, "parser.error.expectedParam") {
		t.Fatal("expected parser.error.expectedParam")
	}
	if hasMessage(errors, "parser.error.expectedSemi") {
		t.Fatal("parser did not recover after parameter list error")
	}
}

func TestParserRecoversAfterUnexpectedTokensInBody(t *testing.T) {
	input := "fun f(): Int {\nreturn a b c\n};"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if countMessage(errors, "parser.error.unexpectedBody") != 2 {
		t.Fatalf("expected 2 parser.error.unexpectedBody errors, got %d", countMessage(errors, "parser.error.unexpectedBody"))
	}
	if hasMessage(errors, "parser.error.expectedSemi") {
		t.Fatal("parser did not recover after unexpected body tokens")
	}
}

func TestParserRecoversAfterMissingOperand(t *testing.T) {
	input := "fun f(): Int {\nreturn a +\n};"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	if !hasMessage(errors, "parser.error.expectedExpr") {
		t.Fatal("expected parser.error.expectedExpr")
	}
	if hasMessage(errors, "parser.error.expectedSemi") {
		t.Fatal("parser did not recover after missing operand")
	}
}

func TestParserReportsCascadeForFuncAndMissingColon(t *testing.T) {
	input := "func calc(a: Int) Int {\nreturn a\n};"
	result := Tokenize(input)
	parser := NewParser(result.Tokens)
	errors := parser.Parse()

	for index, err := range errors {
		t.Logf("error[%d]: key=%s fragment=%q line=%d col=%d", index, err.MessageKey, err.Fragment, err.Line, err.Col)
	}

	expectedKeys := []string{
		"parser.error.expectedFun",
		"parser.error.expectedColon",
	}

	if len(errors) != len(expectedKeys) {
		t.Fatalf("expected %d errors, got %d", len(expectedKeys), len(errors))
	}

	for index, key := range expectedKeys {
		if errors[index].MessageKey != key {
			t.Fatalf("expected error[%d] to be %s, got %s", index, key, errors[index].MessageKey)
		}
	}
}
