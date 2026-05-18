package main

import (
	"strings"
	"testing"
)

func TestRunPolizBuildsQuadsAndValue(t *testing.T) {
	app := NewApp()
	result := app.RunPoliz("2 + 3 * 4")

	if result.OutputKey != "poliz.successValue" {
		t.Fatalf("expected poliz.successValue, got %q (errors=%+v)", result.OutputKey, result.Errors)
	}

	if len(result.Errors) != 0 {
		t.Fatalf("expected no errors, got %+v", result.Errors)
	}

	if got := result.OutputParams["poliz"]; got != "2 3 4 * +" {
		t.Fatalf("unexpected poliz: %q", got)
	}

	quads := result.OutputParams["quads"]
	if !strings.Contains(quads, "(*, 3, 4, t1)") {
		t.Fatalf("expected multiply quad, got %q", quads)
	}
	if !strings.Contains(quads, "(+, 2, t1, t2)") {
		t.Fatalf("expected add quad, got %q", quads)
	}

	if got := result.OutputParams["value"]; got != "14" {
		t.Fatalf("expected value 14, got %q", got)
	}
}

func TestRunPolizSkipsEvaluationForIdentifiers(t *testing.T) {
	app := NewApp()
	result := app.RunPoliz("a + b * 3")

	if result.OutputKey != "poliz.successNoEval" {
		t.Fatalf("expected poliz.successNoEval, got %q", result.OutputKey)
	}

	if result.OutputParams["poliz"] != "a b 3 * +" {
		t.Fatalf("unexpected poliz: %q", result.OutputParams["poliz"])
	}

	if _, ok := result.OutputParams["value"]; ok {
		t.Fatal("did not expect computed value for identifier expression")
	}
}

func TestRunPolizReportsLexerErrors(t *testing.T) {
	app := NewApp()
	result := app.RunPoliz("2 @ 3")

	if result.OutputKey != "poliz.lexerFailed" {
		t.Fatalf("expected poliz.lexerFailed, got %q", result.OutputKey)
	}

	if len(result.Errors) != 1 {
		t.Fatalf("expected 1 lexer error, got %d", len(result.Errors))
	}

	if result.Errors[0].MessageKey != "poliz.error.unexpectedChar" {
		t.Fatalf("expected poliz.error.unexpectedChar, got %q", result.Errors[0].MessageKey)
	}
	if result.Errors[0].Fragment != "@" {
		t.Fatalf("expected fragment @, got %q", result.Errors[0].Fragment)
	}
}

func TestRunPolizReportsSyntaxErrors(t *testing.T) {
	app := NewApp()
	result := app.RunPoliz("2 + )")

	if result.OutputKey != "poliz.parserFailed" {
		t.Fatalf("expected poliz.parserFailed, got %q", result.OutputKey)
	}

	if len(result.Errors) == 0 {
		t.Fatal("expected syntax errors")
	}

	foundOperandError := false
	for _, err := range result.Errors {
		if err.MessageKey == "poliz.error.expectedOperand" {
			foundOperandError = true
		}
	}

	if !foundOperandError {
		t.Fatalf("expected poliz.error.expectedOperand, got %+v", result.Errors)
	}
}

func TestRunPolizReportsDivisionByZero(t *testing.T) {
	app := NewApp()
	result := app.RunPoliz("8 / (3 - 3)")

	if result.OutputKey != "poliz.successDivZero" {
		t.Fatalf("expected poliz.successDivZero, got %q", result.OutputKey)
	}

	if result.OutputParams["poliz"] != "8 3 3 - /" {
		t.Fatalf("unexpected poliz: %q", result.OutputParams["poliz"])
	}
	if _, ok := result.OutputParams["value"]; ok {
		t.Fatal("did not expect numeric value on division by zero")
	}
}
