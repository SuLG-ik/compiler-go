package main

type Token struct {
	Code     int    `json:"code"`
	TypeKey  string `json:"typeKey"`
	Lexeme   string `json:"lexeme"`
	Line     int    `json:"line"`
	StartPos int    `json:"startPos"`
	EndPos   int    `json:"endPos"`
}

type LexerError struct {
	Line    int    `json:"line"`
	Col     int    `json:"col"`
	Message string `json:"message"`
}

type LexerResult struct {
	Tokens []Token      `json:"tokens"`
	Errors []LexerError `json:"errors"`
}

const (
	CodeInt      = 1
	CodeBoolean  = 2
	CodeFloat    = 3
	CodeDouble   = 4
	CodeFun      = 5
	CodeReturn   = 6
	CodeIdent    = 7
	CodeSpace    = 8
	CodeColon    = 9
	CodeComma    = 10
	CodeSemi     = 11
	CodeMinus    = 12
	CodePlus     = 13
	CodeDivide   = 14
	CodeMultiply = 15
	CodeLParen   = 16
	CodeRParen   = 17
	CodeInteger  = 18
	CodeFloatNum = 19
	CodeLBrace   = 20
	CodeRBrace   = 21
)

var keywords = map[string]int{
	"Int":     CodeInt,
	"Boolean": CodeBoolean,
	"Float":   CodeFloat,
	"Double":  CodeDouble,
	"fun":     CodeFun,
	"return":  CodeReturn,
}
var typeKeys = map[int]string{
	CodeInt:      "lexer.type.keyword",
	CodeBoolean:  "lexer.type.keyword",
	CodeFloat:    "lexer.type.keyword",
	CodeDouble:   "lexer.type.keyword",
	CodeFun:      "lexer.type.keyword",
	CodeReturn:   "lexer.type.keyword",
	CodeIdent:    "lexer.type.identifier",
	CodeSpace:    "lexer.type.whitespace",
	CodeColon:    "lexer.type.colon",
	CodeComma:    "lexer.type.comma",
	CodeSemi:     "lexer.type.semicolon",
	CodeMinus:    "lexer.type.minus",
	CodePlus:     "lexer.type.plus",
	CodeDivide:   "lexer.type.divide",
	CodeMultiply: "lexer.type.multiply",
	CodeLParen:   "lexer.type.lparen",
	CodeRParen:   "lexer.type.rparen",
	CodeInteger:  "lexer.type.integer",
	CodeFloatNum: "lexer.type.float",
	CodeLBrace:   "lexer.type.lbrace",
	CodeRBrace:   "lexer.type.rbrace",
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

const (
	stateStart      = 0
	stateIdent      = 1
	stateWhitespace = 2
	stateInteger    = 12
	stateDot        = 13
	stateFraction   = 14
)

func Tokenize(input string) *LexerResult {
	tokens := make([]Token, 0)
	errors := make([]LexerError, 0)

	line := 1
	col := 1
	pos := 0
	n := len(input)

	state := stateStart
	lexemeStart := 0
	startLine := 1
	startCol := 1

	for pos <= n {
		var ch byte
		atEnd := pos >= n
		if !atEnd {
			ch = input[pos]
		}

		switch state {
		case stateStart:
			if atEnd {
				pos++
				continue
			}

			startLine = line
			startCol = col
			lexemeStart = pos

			if isLetter(ch) {
				state = stateIdent
				pos++
				col++
			} else if isDigit(ch) {
				state = stateInteger
				pos++
				col++
			} else if isWhitespace(ch) {
				state = stateWhitespace
				if ch == '\n' {
					pos++
					line++
					col = 1
				} else {
					pos++
					col++
				}
			} else {
				code := 0
				switch ch {
				case ':':
					code = CodeColon
				case ',':
					code = CodeComma
				case ';':
					code = CodeSemi
				case '-':
					code = CodeMinus
				case '+':
					code = CodePlus
				case '/':
					code = CodeDivide
				case '*':
					code = CodeMultiply
				case '(':
					code = CodeLParen
				case ')':
					code = CodeRParen
				case '{':
					code = CodeLBrace
				case '}':
					code = CodeRBrace
				}

				if code != 0 {
					tokens = append(tokens, Token{
						Code:     code,
						TypeKey:  typeKeys[code],
						Lexeme:   string(ch),
						Line:     line,
						StartPos: col,
						EndPos:   col,
					})
					pos++
					col++
				} else {
					errors = append(errors, LexerError{
						Line:    line,
						Col:     col,
						Message: string(ch),
					})
					pos++
					col++
				}
			}

		case stateIdent:
			if !atEnd && isLetter(ch) {
				pos++
				col++
			} else {
				lexeme := input[lexemeStart:pos]
				code := CodeIdent
				if kwCode, ok := keywords[lexeme]; ok {
					code = kwCode
				}
				tokens = append(tokens, Token{
					Code:     code,
					TypeKey:  typeKeys[code],
					Lexeme:   lexeme,
					Line:     startLine,
					StartPos: startCol,
					EndPos:   startCol + len(lexeme) - 1,
				})
				state = stateStart
			}

		case stateWhitespace:
			if !atEnd && isWhitespace(ch) {
				if ch == '\n' {
					pos++
					line++
					col = 1
				} else {
					pos++
					col++
				}
			} else {
				lexeme := input[lexemeStart:pos]
				tokens = append(tokens, Token{
					Code:     CodeSpace,
					TypeKey:  typeKeys[CodeSpace],
					Lexeme:   lexeme,
					Line:     startLine,
					StartPos: startCol,
					EndPos:   startCol + visualLen(input[lexemeStart:pos], startCol) - 1,
				})
				state = stateStart
			}

		case stateInteger:
			if !atEnd && isDigit(ch) {
				pos++
				col++
			} else if !atEnd && ch == '.' {
				state = stateDot
				pos++
				col++
			} else {
				lexeme := input[lexemeStart:pos]
				tokens = append(tokens, Token{
					Code:     CodeInteger,
					TypeKey:  typeKeys[CodeInteger],
					Lexeme:   lexeme,
					Line:     startLine,
					StartPos: startCol,
					EndPos:   startCol + len(lexeme) - 1,
				})
				state = stateStart
			}

		case stateDot:
			if !atEnd && isDigit(ch) {
				state = stateFraction
				pos++
				col++
			} else {
				errors = append(errors, LexerError{
					Line:    line,
					Col:     col - 1,
					Message: input[lexemeStart:pos],
				})
				state = stateStart
			}

		case stateFraction:
			if !atEnd && isDigit(ch) {
				pos++
				col++
			} else {
				lexeme := input[lexemeStart:pos]
				tokens = append(tokens, Token{
					Code:     CodeFloatNum,
					TypeKey:  typeKeys[CodeFloatNum],
					Lexeme:   lexeme,
					Line:     startLine,
					StartPos: startCol,
					EndPos:   startCol + len(lexeme) - 1,
				})
				state = stateStart
			}
		}
	}

	return &LexerResult{
		Tokens: tokens,
		Errors: errors,
	}
}

func visualLen(s string, _ int) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count++
	}
	return count
}
