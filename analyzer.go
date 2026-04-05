package main

type AnalyzerError struct {
	Line       int    `json:"line"`
	Col        int    `json:"col"`
	Fragment   string `json:"fragment,omitempty"`
	Message    string `json:"message"`
	MessageKey string `json:"messageKey,omitempty"`
}

type AnalyzerResult struct {
	Output    string          `json:"output"`
	OutputKey string          `json:"outputKey,omitempty"`
	Errors    []AnalyzerError `json:"errors"`
	Tokens    []Token         `json:"tokens"`
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
