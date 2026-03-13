package main

import "testing"

func TestTokenizeKotlinFunction(t *testing.T) {
	input := "fun calc(a: Int, b: Int, c: Int): Int {\n    return a + (b * c)\n};"
	result := Tokenize(input)

	if len(result.Errors) != 0 {
		t.Errorf("expected no errors, got %d:", len(result.Errors))
		for _, e := range result.Errors {
			t.Errorf("  line %d col %d: %s", e.Line, e.Col, e.Message)
		}
	}

	expected := []struct {
		code   int
		lexeme string
	}{
		{CodeFun, "fun"},
		{CodeSpace, " "},
		{CodeIdent, "calc"},
		{CodeLParen, "("},
		{CodeIdent, "a"},
		{CodeColon, ":"},
		{CodeSpace, " "},
		{CodeInt, "Int"},
		{CodeComma, ","},
		{CodeSpace, " "},
		{CodeIdent, "b"},
		{CodeColon, ":"},
		{CodeSpace, " "},
		{CodeInt, "Int"},
		{CodeComma, ","},
		{CodeSpace, " "},
		{CodeIdent, "c"},
		{CodeColon, ":"},
		{CodeSpace, " "},
		{CodeInt, "Int"},
		{CodeRParen, ")"},
		{CodeColon, ":"},
		{CodeSpace, " "},
		{CodeInt, "Int"},
		{CodeSpace, " "},
		{CodeLBrace, "{"},
		{CodeSpace, "\n    "},
		{CodeReturn, "return"},
		{CodeSpace, " "},
		{CodeIdent, "a"},
		{CodeSpace, " "},
		{CodePlus, "+"},
		{CodeSpace, " "},
		{CodeLParen, "("},
		{CodeIdent, "b"},
		{CodeSpace, " "},
		{CodeMultiply, "*"},
		{CodeSpace, " "},
		{CodeIdent, "c"},
		{CodeRParen, ")"},
		{CodeSpace, "\n"},
		{CodeRBrace, "}"},
		{CodeSemi, ";"},
	}

	if len(result.Tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(result.Tokens))
		for i, tok := range result.Tokens {
			t.Logf("  [%d] code=%d lexeme=%q", i, tok.Code, tok.Lexeme)
		}
	}

	for i, exp := range expected {
		tok := result.Tokens[i]
		if tok.Code != exp.code || tok.Lexeme != exp.lexeme {
			t.Errorf("token %d: expected code=%d lexeme=%q, got code=%d lexeme=%q",
				i, exp.code, exp.lexeme, tok.Code, tok.Lexeme)
		}
	}
}

func TestTokenizeMultiline(t *testing.T) {
	input := "a\nb"
	result := Tokenize(input)
	if len(result.Errors) != 0 {
		t.Errorf("expected no errors, got %d", len(result.Errors))
	}
	if len(result.Tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(result.Tokens))
	}
	if result.Tokens[0].Line != 1 {
		t.Errorf("token 0: expected line 1, got %d", result.Tokens[0].Line)
	}
	if result.Tokens[2].Line != 2 {
		t.Errorf("token 2: expected line 2, got %d", result.Tokens[2].Line)
	}
}

func TestTokenizeNumbers(t *testing.T) {
	input := "123 45.67"
	result := Tokenize(input)
	if len(result.Errors) != 0 {
		t.Errorf("expected no errors, got %d", len(result.Errors))
	}
	if len(result.Tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(result.Tokens))
	}
	if result.Tokens[0].Code != CodeInteger || result.Tokens[0].Lexeme != "123" {
		t.Errorf("token 0: expected integer 123, got code=%d lexeme=%q", result.Tokens[0].Code, result.Tokens[0].Lexeme)
	}
	if result.Tokens[2].Code != CodeFloatNum || result.Tokens[2].Lexeme != "45.67" {
		t.Errorf("token 2: expected float 45.67, got code=%d lexeme=%q", result.Tokens[2].Code, result.Tokens[2].Lexeme)
	}
}

func TestTokenizeErrors(t *testing.T) {
	input := "a @ b"
	result := Tokenize(input)
	if len(result.Errors) != 1 {
		t.Fatalf("expected 1 error, got %d", len(result.Errors))
	}
	if result.Errors[0].Line != 1 || result.Errors[0].Col != 3 {
		t.Errorf("error at wrong position: line=%d col=%d", result.Errors[0].Line, result.Errors[0].Col)
	}
}

func TestTokenizeKeywords(t *testing.T) {
	input := "Int Boolean Float Double fun return"
	result := Tokenize(input)
	if len(result.Errors) != 0 {
		t.Errorf("expected no errors, got %d", len(result.Errors))
	}
	expectedCodes := []int{CodeInt, CodeSpace, CodeBoolean, CodeSpace, CodeFloat, CodeSpace, CodeDouble, CodeSpace, CodeFun, CodeSpace, CodeReturn}
	if len(result.Tokens) != len(expectedCodes) {
		t.Fatalf("expected %d tokens, got %d", len(expectedCodes), len(result.Tokens))
	}
	for i, code := range expectedCodes {
		if result.Tokens[i].Code != code {
			t.Errorf("token %d: expected code %d, got %d (lexeme=%q)", i, code, result.Tokens[i].Code, result.Tokens[i].Lexeme)
		}
	}
}

func TestTokenizeEmpty(t *testing.T) {
	result := Tokenize("")
	if len(result.Tokens) != 0 {
		t.Errorf("expected 0 tokens, got %d", len(result.Tokens))
	}
	if len(result.Errors) != 0 {
		t.Errorf("expected 0 errors, got %d", len(result.Errors))
	}
}
