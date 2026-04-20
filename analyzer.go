package main

import "strconv"

type AnalyzerError struct {
	Line          int               `json:"line"`
	Col           int               `json:"col"`
	Fragment      string            `json:"fragment,omitempty"`
	Message       string            `json:"message"`
	MessageKey    string            `json:"messageKey,omitempty"`
	MessageParams map[string]string `json:"messageParams,omitempty"`
}

type AnalyzerResult struct {
	Output       string            `json:"output"`
	OutputKey    string            `json:"outputKey,omitempty"`
	OutputParams map[string]string `json:"outputParams,omitempty"`
	Errors       []AnalyzerError   `json:"errors"`
	Tokens       []Token           `json:"tokens"`
}

func (a *App) RunAnalyzer(content string) *AnalyzerResult {
	if content == "" {
		return &AnalyzerResult{
			OutputKey: "analyzer.emptyOutput",
			Errors: []AnalyzerError{
				{Line: 0, Col: 0, MessageKey: "analyzer.emptyMessage"},
			},
			Tokens: []Token{},
		}
	}

	result := Tokenize(content)
	analyzerErrors := make([]AnalyzerError, 0)

	for _, e := range result.Errors {
		analyzerErrors = append(analyzerErrors, AnalyzerError{
			Line:       e.Line,
			Col:        e.Col,
			Fragment:   e.Message,
			MessageKey: "lexer.error.unexpectedChar",
		})
	}

	if len(analyzerErrors) > 0 {
		return &AnalyzerResult{
			OutputKey: "analyzer.lexerFailed",
			Tokens:    result.Tokens,
			Errors:    analyzerErrors,
		}
	}

	parser := NewParser(result.Tokens)
	parserErrors := parser.Parse()
	for _, e := range parserErrors {
		analyzerErrors = append(analyzerErrors, AnalyzerError{
			Line:       e.Line,
			Col:        e.Col,
			Fragment:   e.Fragment,
			MessageKey: e.MessageKey,
		})
	}

	outputKey := ""
	if len(analyzerErrors) == 0 {
		outputKey = "parser.success"
	}

	return &AnalyzerResult{
		OutputKey: outputKey,
		Tokens:    result.Tokens,
		Errors:    analyzerErrors,
	}
}

func (a *App) RunSemanticAnalyzer(content string) *AnalyzerResult {
	if content == "" {
		return &AnalyzerResult{
			OutputKey: "analyzer.emptyOutput",
			Errors: []AnalyzerError{
				{Line: 0, Col: 0, MessageKey: "analyzer.emptyMessage"},
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
		return &AnalyzerResult{
			OutputKey: "semantic.lexerFailed",
			Tokens:    lex.Tokens,
			Errors:    errs,
		}
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
		return &AnalyzerResult{
			OutputKey: "semantic.parserFailed",
			Tokens:    lex.Tokens,
			Errors:    errs,
		}
	}

	builder := NewAstBuilder(lex.Tokens)
	root, semErrors := builder.Build()
	astText := PrintAst(root)

	outputKey := "semantic.success"
	if len(semErrors) > 0 {
		outputKey = "semantic.failed"
	}

	return &AnalyzerResult{
		OutputKey:    outputKey,
		OutputParams: map[string]string{"ast": astText, "count": strconv.Itoa(len(semErrors))},
		Errors:       semErrors,
		Tokens:       lex.Tokens,
	}
}
