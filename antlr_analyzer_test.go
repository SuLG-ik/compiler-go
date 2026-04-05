package main

import "testing"

func TestRunAntlrAnalyzerValidFunction(t *testing.T) {
	app := NewApp()
	result := app.RunAntlrAnalyzer("fun calc(a: Int, b: Int, c: Int): Int {\nreturn a + (b * c)\n};")

	if result.OutputKey != "antlr.success" {
		t.Fatalf("expected antlr.success, got %q", result.OutputKey)
	}

	if len(result.Errors) != 0 {
		t.Fatalf("expected 0 errors, got %d", len(result.Errors))
	}
}

func TestRunAntlrAnalyzerLexerError(t *testing.T) {
	app := NewApp()
	result := app.RunAntlrAnalyzer("fun calc(a: Int): Int {\nreturn a @ b\n};")

	if result.OutputKey != "antlr.lexerFailed" {
		t.Fatalf("expected antlr.lexerFailed, got %q", result.OutputKey)
	}

	if len(result.Errors) == 0 {
		t.Fatal("expected lexical errors")
	}
}

func TestRunAntlrAnalyzerSyntaxError(t *testing.T) {
	app := NewApp()
	result := app.RunAntlrAnalyzer("func calc(a: Int) Int {\nreturn a\n};")

	if result.OutputKey != "antlr.syntaxFailed" {
		t.Fatalf("expected antlr.syntaxFailed, got %q", result.OutputKey)
	}

	if len(result.Errors) == 0 {
		t.Fatal("expected syntax errors")
	}
}
