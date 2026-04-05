package main

import (
	"fmt"
	"strings"

	antlr "github.com/antlr4-go/antlr/v4"

	antlrgen "compiler/antlrgen/antlr"
)

type antlrErrorCollector struct {
	*antlr.DefaultErrorListener
	errors []AnalyzerError
}

func (c *antlrErrorCollector) SyntaxError(_ antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	fragment := antlrFragmentFromSymbol(offendingSymbol)
	if fragment == "" {
		fragment = antlrFragmentFromMessage(msg)
	}

	c.errors = append(c.errors, AnalyzerError{
		Line:     line,
		Col:      column + 1,
		Fragment: fragment,
		Message:  msg,
	})
}

func antlrFragmentFromSymbol(offendingSymbol interface{}) string {
	token, ok := offendingSymbol.(antlr.Token)
	if !ok || token == nil {
		return ""
	}
	text := token.GetText()
	if text == "<EOF>" {
		return ""
	}
	return text
}

func antlrFragmentFromMessage(msg string) string {
	start := strings.Index(msg, "'")
	if start == -1 {
		return ""
	}
	end := strings.Index(msg[start+1:], "'")
	if end == -1 {
		return ""
	}
	return msg[start+1 : start+1+end]
}

func newAntlrAnalyzerResult(outputKey string, errors []AnalyzerError) *AnalyzerResult {
	return &AnalyzerResult{
		OutputKey: outputKey,
		Errors:    errors,
		Tokens:    []Token{},
	}
}

func (a *App) RunAntlrAnalyzer(content string) *AnalyzerResult {
	if content == "" {
		return &AnalyzerResult{
			OutputKey: "analyzer.emptyOutput",
			Errors: []AnalyzerError{
				{Line: 0, Col: 0, MessageKey: "analyzer.emptyMessage"},
			},
			Tokens: []Token{},
		}
	}

	input := antlr.NewInputStream(content)
	lexer := antlrgen.NewKotlinFunctionLexer(input)
	collector := &antlrErrorCollector{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(collector)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	stream.Fill()

	if len(collector.errors) > 0 {
		return newAntlrAnalyzerResult("antlr.lexerFailed", collector.errors)
	}

	parser := antlrgen.NewKotlinFunctionParser(stream)
	parser.BuildParseTrees = false
	parser.RemoveErrorListeners()
	parser.AddErrorListener(collector)
	parser.Program()

	if len(collector.errors) > 0 {
		return newAntlrAnalyzerResult("antlr.syntaxFailed", collector.errors)
	}

	return newAntlrAnalyzerResult("antlr.success", []AnalyzerError{})
}

func formatAntlrErrorList(errors []AnalyzerError) string {
	parts := make([]string, 0, len(errors))
	for _, err := range errors {
		parts = append(parts, fmt.Sprintf("%d:%d %s", err.Line, err.Col, err.Message))
	}
	return strings.Join(parts, "\n")
}
