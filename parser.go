package main

const codeEOF = -1

type ParserError struct {
	Line       int    `json:"line"`
	Col        int    `json:"col"`
	Fragment   string `json:"fragment"`
	MessageKey string `json:"messageKey"`
}

type Parser struct {
	tokens []Token
	pos    int
	errors []ParserError
}

func NewParser(allTokens []Token) *Parser {
	filtered := make([]Token, 0, len(allTokens))
	for _, t := range allTokens {
		if t.Code != CodeSpace {
			filtered = append(filtered, t)
		}
	}
	return &Parser{tokens: filtered}
}

func (p *Parser) peek() int {
	if p.pos < len(p.tokens) {
		return p.tokens[p.pos].Code
	}
	return codeEOF
}

func (p *Parser) peekOffset(offset int) int {
	index := p.pos + offset
	if index >= 0 && index < len(p.tokens) {
		return p.tokens[index].Code
	}
	return codeEOF
}

func (p *Parser) current() Token {
	if p.pos < len(p.tokens) {
		return p.tokens[p.pos]
	}
	if len(p.tokens) > 0 {
		last := p.tokens[len(p.tokens)-1]
		return Token{Code: codeEOF, Lexeme: "", Line: last.Line, StartPos: last.EndPos + 1, EndPos: last.EndPos + 1}
	}
	return Token{Code: codeEOF, Lexeme: "", Line: 1, StartPos: 1, EndPos: 1}
}

func (p *Parser) advance() Token {
	tok := p.current()
	if p.pos < len(p.tokens) {
		p.pos++
	}
	return tok
}

func (p *Parser) atEnd() bool {
	return p.pos >= len(p.tokens)
}

func (p *Parser) addError(msgKey string) {
	tok := p.current()
	p.errors = append(p.errors, ParserError{
		Line:       tok.Line,
		Col:        tok.StartPos,
		Fragment:   tok.Lexeme,
		MessageKey: msgKey,
	})
}

func containsCode(code int, codes []int) bool {
	for _, candidate := range codes {
		if code == candidate {
			return true
		}
	}
	return false
}

func (p *Parser) skipTo(codes ...int) {
	for !p.atEnd() {
		cur := p.peek()
		for _, c := range codes {
			if cur == c {
				return
			}
		}
		p.advance()
	}
}

func (p *Parser) expect(code int, msgKey string, follow ...int) bool {
	if p.peek() == code {
		p.advance()
		return true
	}

	if p.atEnd() {
		p.addError(msgKey)
		return false
	}

	startPos := p.pos
	p.addError(msgKey)

	syncSet := make([]int, 0, 1+len(follow))
	syncSet = append(syncSet, code)
	syncSet = append(syncSet, follow...)
	p.skipTo(syncSet...)

	if p.peek() == code {
		p.advance()
		return true
	}

	if containsCode(p.peek(), follow) || p.atEnd() {
		return false
	}

	if p.pos == startPos && !p.atEnd() {
		p.advance()
	}

	return false
}

func (p *Parser) Parse() []ParserError {
	if len(p.tokens) == 0 {
		return p.errors
	}
	p.parseProgram()
	return p.errors
}

func (p *Parser) parseProgram() {
	p.parseFunDecl()

	if p.atEnd() {
		p.addError("parser.error.expectedSemi")
	} else if !p.expect(CodeSemi, "parser.error.expectedSemi") {
	}

	if !p.atEnd() {
		p.addError("parser.error.unexpectedAfter")
	}
}

func (p *Parser) parseFunDecl() {
	if p.peek() == CodeFun {
		p.advance()
	} else {
		p.addError("parser.error.expectedFun")
		if p.peek() == CodeIdent && p.peekOffset(1) == CodeIdent {
			p.advance()
		}
	}

	p.expect(CodeIdent, "parser.error.expectedName",
		CodeLParen, CodeRParen, CodeColon,
		CodeInt, CodeBoolean, CodeFloat, CodeDouble,
		CodeLBrace, CodeRBrace, CodeSemi)

	p.expect(CodeLParen, "parser.error.expectedLParen",
		CodeIdent, CodeRParen, CodeColon,
		CodeInt, CodeBoolean, CodeFloat, CodeDouble,
		CodeLBrace, CodeRBrace, CodeSemi)

	if p.peek() != CodeRParen && p.peek() != codeEOF {
		p.parseParamList()
	}

	p.expect(CodeRParen, "parser.error.expectedRParen",
		CodeColon, CodeInt, CodeBoolean, CodeFloat, CodeDouble,
		CodeLBrace, CodeRBrace, CodeSemi)

	p.expect(CodeColon, "parser.error.expectedColon",
		CodeInt, CodeBoolean, CodeFloat, CodeDouble,
		CodeLBrace, CodeRBrace, CodeSemi)

	p.parseType(CodeLBrace, CodeRBrace, CodeSemi)

	p.expect(CodeLBrace, "parser.error.expectedLBrace",
		CodeReturn, CodeIdent, CodeInteger, CodeFloatNum,
		CodeRBrace, CodeSemi)

	p.parseBody()

	p.expect(CodeRBrace, "parser.error.expectedRBrace",
		CodeSemi)
}

func (p *Parser) parseParamList() {
	p.parseParam()
	for p.peek() == CodeComma {
		p.advance()
		if p.peek() == CodeRParen || p.atEnd() {
			p.addError("parser.error.expectedParam")
			break
		}
		p.parseParam()
	}
}

func (p *Parser) parseParam() {
	p.expect(CodeIdent, "parser.error.expectedParam",
		CodeColon, CodeInt, CodeBoolean, CodeFloat, CodeDouble,
		CodeComma, CodeRParen)

	p.expect(CodeColon, "parser.error.expectedColon",
		CodeInt, CodeBoolean, CodeFloat, CodeDouble,
		CodeComma, CodeRParen)

	p.parseType(CodeComma, CodeRParen)
}

func (p *Parser) parseType(follow ...int) {
	code := p.peek()
	if code == CodeInt || code == CodeBoolean || code == CodeFloat || code == CodeDouble {
		p.advance()
		return
	}

	if p.atEnd() {
		p.addError("parser.error.expectedType")
		return
	}

	startPos := p.pos
	p.addError("parser.error.expectedType")

	typeAndFollow := []int{CodeInt, CodeBoolean, CodeFloat, CodeDouble}
	typeAndFollow = append(typeAndFollow, follow...)
	p.skipTo(typeAndFollow...)

	code = p.peek()
	if code == CodeInt || code == CodeBoolean || code == CodeFloat || code == CodeDouble {
		p.advance()
		return
	}

	if containsCode(p.peek(), follow) || p.atEnd() {
		return
	}

	if p.pos == startPos && !p.atEnd() {
		p.advance()
	}
}

func (p *Parser) parseBody() {
	if p.peek() == CodeReturn {
		p.parseReturnStmt()
	} else if p.peek() == CodeRBrace {
		p.addError("parser.error.expectedReturn")
	} else if !p.atEnd() {
		p.addError("parser.error.expectedReturn")
		p.skipTo(CodeReturn, CodeRBrace, CodeSemi)
		if p.peek() == CodeReturn {
			p.parseReturnStmt()
		}
	}

	for p.peek() != CodeRBrace && !p.atEnd() {
		p.addError("parser.error.unexpectedBody")
		p.advance()
	}
}

func (p *Parser) parseReturnStmt() {
	p.expect(CodeReturn, "parser.error.expectedReturn",
		CodeIdent, CodeInteger, CodeFloatNum, CodeLParen, CodeRBrace)
	p.parseExpr()
}

func (p *Parser) parseExpr() {
	p.parseTerm()
	for p.peek() == CodePlus || p.peek() == CodeMinus {
		p.advance()
		p.parseTerm()
	}
}

func (p *Parser) parseTerm() {
	p.parseFactor()
	for p.peek() == CodeMultiply || p.peek() == CodeDivide {
		p.advance()
		p.parseFactor()
	}
}

func (p *Parser) parseFactor() {
	switch p.peek() {
	case CodeIdent, CodeInteger, CodeFloatNum:
		p.advance()
	case CodeLParen:
		p.advance()
		p.parseExpr()
		p.expect(CodeRParen, "parser.error.expectedRParen",
			CodePlus, CodeMinus, CodeMultiply, CodeDivide,
			CodeRBrace, CodeSemi)
	default:
		if !p.atEnd() {
			p.addError("parser.error.expectedExpr")
		}
	}
}
