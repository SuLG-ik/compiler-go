package main

type AnalyzerError struct {
	Line       int    `json:"line"`
	Col        int    `json:"col"`
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

	analyzerErrors := make([]AnalyzerError, 0, len(result.Errors))
	for _, e := range result.Errors {
		analyzerErrors = append(analyzerErrors, AnalyzerError{
			Line:    e.Line,
			Col:     e.Col,
			Message: e.Message,
		})
	}

	return &AnalyzerResult{
		Tokens: result.Tokens,
		Errors: analyzerErrors,
	}
}
