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
}

func (a *App) RunAnalyzer(content string) *AnalyzerResult {
	if content == "" {
		return &AnalyzerResult{
			OutputKey: "analyzer.emptyOutput",
			Errors: []AnalyzerError{
				{Line: 0, Col: 0, MessageKey: "analyzer.emptyMessage"},
			},
		}
	}
	return &AnalyzerResult{
		OutputKey: "analyzer.notImplemented",
		Errors:    []AnalyzerError{},
	}
}
