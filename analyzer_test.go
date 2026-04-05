package main

import "testing"

func TestRunAnalyzerSkipsParserOnLexerError(t *testing.T) {
	app := NewApp()
	result := app.RunAnalyzer("fun f(): Int {\nreturn a @ b\n};")

	if result.OutputKey != "analyzer.lexerFailed" {
		t.Fatalf("expected analyzer.lexerFailed, got %q", result.OutputKey)
	}

	if len(result.Errors) != 1 {
		t.Fatalf("expected 1 lexer error, got %d", len(result.Errors))
	}

	if result.Errors[0].MessageKey != "lexer.error.unexpectedChar" {
		t.Fatalf("expected lexer.error.unexpectedChar, got %q", result.Errors[0].MessageKey)
	}

	for _, err := range result.Errors {
		if err.MessageKey != "lexer.error.unexpectedChar" {
			t.Fatalf("expected only lexer errors, got %q", err.MessageKey)
		}
	}
}

func TestRunAnalyzerRunsParserWithoutLexerErrors(t *testing.T) {
	app := NewApp()
	result := app.RunAnalyzer("fun f(): Int {\nreturn\n};")

	if result.OutputKey != "" {
		t.Fatalf("expected empty output key for parser errors, got %q", result.OutputKey)
	}

	if len(result.Errors) == 0 {
		t.Fatal("expected parser errors")
	}

	foundParserError := false
	for _, err := range result.Errors {
		if err.MessageKey == "parser.error.expectedExpr" {
			foundParserError = true
		}
		if err.MessageKey == "lexer.error.unexpectedChar" {
			t.Fatal("did not expect lexer error")
		}
	}

	if !foundParserError {
		t.Fatal("expected parser.error.expectedExpr")
	}
}
