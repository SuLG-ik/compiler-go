// Code generated from antlr/KotlinFunction.g4 by ANTLR 4.13.1. DO NOT EDIT.

package antlrgen // KotlinFunction
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type KotlinFunctionParser struct {
	*antlr.BaseParser
}

var KotlinFunctionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kotlinfunctionParserInit() {
	staticData := &KotlinFunctionParserStaticData
	staticData.LiteralNames = []string{
		"", "'fun'", "'return'", "'Int'", "'Boolean'", "'Float'", "'Double'",
		"':'", "','", "';'", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'{'",
		"'}'",
	}
	staticData.SymbolicNames = []string{
		"", "FUN", "RETURN", "INT_TYPE", "BOOLEAN_TYPE", "FLOAT_TYPE", "DOUBLE_TYPE",
		"COLON", "COMMA", "SEMI", "PLUS", "MINUS", "STAR", "SLASH", "LPAREN",
		"RPAREN", "LBRACE", "RBRACE", "FLOAT_LITERAL", "INTEGER_LITERAL", "IDENTIFIER",
		"WS",
	}
	staticData.RuleNames = []string{
		"program", "funDecl", "paramList", "param", "typeRef", "body", "returnStmt",
		"expr", "term", "factor",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 82, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 29, 8, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 41, 8, 2, 10, 2, 12, 2,
		44, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 5, 7, 60, 8, 7, 10, 7, 12, 7, 63, 9, 7, 1, 8, 1, 8,
		1, 8, 5, 8, 68, 8, 8, 10, 8, 12, 8, 71, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 3, 9, 80, 8, 9, 1, 9, 0, 0, 10, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 0, 3, 1, 0, 3, 6, 1, 0, 10, 11, 1, 0, 12, 13, 78, 0, 20, 1, 0,
		0, 0, 2, 24, 1, 0, 0, 0, 4, 37, 1, 0, 0, 0, 6, 45, 1, 0, 0, 0, 8, 49, 1,
		0, 0, 0, 10, 51, 1, 0, 0, 0, 12, 53, 1, 0, 0, 0, 14, 56, 1, 0, 0, 0, 16,
		64, 1, 0, 0, 0, 18, 79, 1, 0, 0, 0, 20, 21, 3, 2, 1, 0, 21, 22, 5, 9, 0,
		0, 22, 23, 5, 0, 0, 1, 23, 1, 1, 0, 0, 0, 24, 25, 5, 1, 0, 0, 25, 26, 5,
		20, 0, 0, 26, 28, 5, 14, 0, 0, 27, 29, 3, 4, 2, 0, 28, 27, 1, 0, 0, 0,
		28, 29, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 31, 5, 15, 0, 0, 31, 32, 5,
		7, 0, 0, 32, 33, 3, 8, 4, 0, 33, 34, 5, 16, 0, 0, 34, 35, 3, 10, 5, 0,
		35, 36, 5, 17, 0, 0, 36, 3, 1, 0, 0, 0, 37, 42, 3, 6, 3, 0, 38, 39, 5,
		8, 0, 0, 39, 41, 3, 6, 3, 0, 40, 38, 1, 0, 0, 0, 41, 44, 1, 0, 0, 0, 42,
		40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 5, 1, 0, 0, 0, 44, 42, 1, 0, 0,
		0, 45, 46, 5, 20, 0, 0, 46, 47, 5, 7, 0, 0, 47, 48, 3, 8, 4, 0, 48, 7,
		1, 0, 0, 0, 49, 50, 7, 0, 0, 0, 50, 9, 1, 0, 0, 0, 51, 52, 3, 12, 6, 0,
		52, 11, 1, 0, 0, 0, 53, 54, 5, 2, 0, 0, 54, 55, 3, 14, 7, 0, 55, 13, 1,
		0, 0, 0, 56, 61, 3, 16, 8, 0, 57, 58, 7, 1, 0, 0, 58, 60, 3, 16, 8, 0,
		59, 57, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1,
		0, 0, 0, 62, 15, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 69, 3, 18, 9, 0, 65,
		66, 7, 2, 0, 0, 66, 68, 3, 18, 9, 0, 67, 65, 1, 0, 0, 0, 68, 71, 1, 0,
		0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 17, 1, 0, 0, 0, 71, 69,
		1, 0, 0, 0, 72, 80, 5, 20, 0, 0, 73, 80, 5, 18, 0, 0, 74, 80, 5, 19, 0,
		0, 75, 76, 5, 14, 0, 0, 76, 77, 3, 14, 7, 0, 77, 78, 5, 15, 0, 0, 78, 80,
		1, 0, 0, 0, 79, 72, 1, 0, 0, 0, 79, 73, 1, 0, 0, 0, 79, 74, 1, 0, 0, 0,
		79, 75, 1, 0, 0, 0, 80, 19, 1, 0, 0, 0, 5, 28, 42, 61, 69, 79,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// KotlinFunctionParserInit initializes any static state used to implement KotlinFunctionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewKotlinFunctionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func KotlinFunctionParserInit() {
	staticData := &KotlinFunctionParserStaticData
	staticData.once.Do(kotlinfunctionParserInit)
}

// NewKotlinFunctionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewKotlinFunctionParser(input antlr.TokenStream) *KotlinFunctionParser {
	KotlinFunctionParserInit()
	this := new(KotlinFunctionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &KotlinFunctionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "KotlinFunction.g4"

	return this
}

// KotlinFunctionParser tokens.
const (
	KotlinFunctionParserEOF             = antlr.TokenEOF
	KotlinFunctionParserFUN             = 1
	KotlinFunctionParserRETURN          = 2
	KotlinFunctionParserINT_TYPE        = 3
	KotlinFunctionParserBOOLEAN_TYPE    = 4
	KotlinFunctionParserFLOAT_TYPE      = 5
	KotlinFunctionParserDOUBLE_TYPE     = 6
	KotlinFunctionParserCOLON           = 7
	KotlinFunctionParserCOMMA           = 8
	KotlinFunctionParserSEMI            = 9
	KotlinFunctionParserPLUS            = 10
	KotlinFunctionParserMINUS           = 11
	KotlinFunctionParserSTAR            = 12
	KotlinFunctionParserSLASH           = 13
	KotlinFunctionParserLPAREN          = 14
	KotlinFunctionParserRPAREN          = 15
	KotlinFunctionParserLBRACE          = 16
	KotlinFunctionParserRBRACE          = 17
	KotlinFunctionParserFLOAT_LITERAL   = 18
	KotlinFunctionParserINTEGER_LITERAL = 19
	KotlinFunctionParserIDENTIFIER      = 20
	KotlinFunctionParserWS              = 21
)

// KotlinFunctionParser rules.
const (
	KotlinFunctionParserRULE_program    = 0
	KotlinFunctionParserRULE_funDecl    = 1
	KotlinFunctionParserRULE_paramList  = 2
	KotlinFunctionParserRULE_param      = 3
	KotlinFunctionParserRULE_typeRef    = 4
	KotlinFunctionParserRULE_body       = 5
	KotlinFunctionParserRULE_returnStmt = 6
	KotlinFunctionParserRULE_expr       = 7
	KotlinFunctionParserRULE_term       = 8
	KotlinFunctionParserRULE_factor     = 9
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunDecl() IFunDeclContext
	SEMI() antlr.TerminalNode
	EOF() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KotlinFunctionParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) FunDecl() IFunDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunDeclContext)
}

func (s *ProgramContext) SEMI() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserSEMI, 0)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *KotlinFunctionParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, KotlinFunctionParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.FunDecl()
	}
	{
		p.SetState(21)
		p.Match(KotlinFunctionParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(22)
		p.Match(KotlinFunctionParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunDeclContext is an interface to support dynamic dispatch.
type IFunDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	COLON() antlr.TerminalNode
	TypeRef() ITypeRefContext
	LBRACE() antlr.TerminalNode
	Body() IBodyContext
	RBRACE() antlr.TerminalNode
	ParamList() IParamListContext

	// IsFunDeclContext differentiates from other interfaces.
	IsFunDeclContext()
}

type FunDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunDeclContext() *FunDeclContext {
	var p = new(FunDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_funDecl
	return p
}

func InitEmptyFunDeclContext(p *FunDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_funDecl
}

func (*FunDeclContext) IsFunDeclContext() {}

func NewFunDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunDeclContext {
	var p = new(FunDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KotlinFunctionParserRULE_funDecl

	return p
}

func (s *FunDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FunDeclContext) FUN() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserFUN, 0)
}

func (s *FunDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserIDENTIFIER, 0)
}

func (s *FunDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserLPAREN, 0)
}

func (s *FunDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserRPAREN, 0)
}

func (s *FunDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserCOLON, 0)
}

func (s *FunDeclContext) TypeRef() ITypeRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *FunDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserLBRACE, 0)
}

func (s *FunDeclContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *FunDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserRBRACE, 0)
}

func (s *FunDeclContext) ParamList() IParamListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *FunDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *KotlinFunctionParser) FunDecl() (localctx IFunDeclContext) {
	localctx = NewFunDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, KotlinFunctionParserRULE_funDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(KotlinFunctionParserFUN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(25)
		p.Match(KotlinFunctionParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(26)
		p.Match(KotlinFunctionParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KotlinFunctionParserIDENTIFIER {
		{
			p.SetState(27)
			p.ParamList()
		}

	}
	{
		p.SetState(30)
		p.Match(KotlinFunctionParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(31)
		p.Match(KotlinFunctionParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.TypeRef()
	}
	{
		p.SetState(33)
		p.Match(KotlinFunctionParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.Body()
	}
	{
		p.SetState(35)
		p.Match(KotlinFunctionParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_paramList
	return p
}

func InitEmptyParamListContext(p *ParamListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_paramList
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KotlinFunctionParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamListContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KotlinFunctionParserCOMMA)
}

func (s *ParamListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserCOMMA, i)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *KotlinFunctionParser) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, KotlinFunctionParserRULE_paramList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Param()
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == KotlinFunctionParserCOMMA {
		{
			p.SetState(38)
			p.Match(KotlinFunctionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Param()
		}

		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	TypeRef() ITypeRefContext

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KotlinFunctionParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserIDENTIFIER, 0)
}

func (s *ParamContext) COLON() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserCOLON, 0)
}

func (s *ParamContext) TypeRef() ITypeRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *KotlinFunctionParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, KotlinFunctionParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(KotlinFunctionParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(KotlinFunctionParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.TypeRef()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeRefContext is an interface to support dynamic dispatch.
type ITypeRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_TYPE() antlr.TerminalNode
	BOOLEAN_TYPE() antlr.TerminalNode
	FLOAT_TYPE() antlr.TerminalNode
	DOUBLE_TYPE() antlr.TerminalNode

	// IsTypeRefContext differentiates from other interfaces.
	IsTypeRefContext()
}

type TypeRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRefContext() *TypeRefContext {
	var p = new(TypeRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_typeRef
	return p
}

func InitEmptyTypeRefContext(p *TypeRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_typeRef
}

func (*TypeRefContext) IsTypeRefContext() {}

func NewTypeRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRefContext {
	var p = new(TypeRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KotlinFunctionParserRULE_typeRef

	return p
}

func (s *TypeRefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeRefContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserINT_TYPE, 0)
}

func (s *TypeRefContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserBOOLEAN_TYPE, 0)
}

func (s *TypeRefContext) FLOAT_TYPE() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserFLOAT_TYPE, 0)
}

func (s *TypeRefContext) DOUBLE_TYPE() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserDOUBLE_TYPE, 0)
}

func (s *TypeRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *KotlinFunctionParser) TypeRef() (localctx ITypeRefContext) {
	localctx = NewTypeRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, KotlinFunctionParserRULE_typeRef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&120) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ReturnStmt() IReturnStmtContext

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_body
	return p
}

func InitEmptyBodyContext(p *BodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_body
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KotlinFunctionParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *KotlinFunctionParser) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, KotlinFunctionParserRULE_body)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.ReturnStmt()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KotlinFunctionParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserRETURN, 0)
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *KotlinFunctionParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, KotlinFunctionParserRULE_returnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(KotlinFunctionParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Expr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KotlinFunctionParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExprContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(KotlinFunctionParserPLUS)
}

func (s *ExprContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserPLUS, i)
}

func (s *ExprContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(KotlinFunctionParserMINUS)
}

func (s *ExprContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserMINUS, i)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *KotlinFunctionParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, KotlinFunctionParserRULE_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Term()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == KotlinFunctionParserPLUS || _la == KotlinFunctionParserMINUS {
		{
			p.SetState(57)
			_la = p.GetTokenStream().LA(1)

			if !(_la == KotlinFunctionParserPLUS || _la == KotlinFunctionParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(58)
			p.Term()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFactor() []IFactorContext
	Factor(i int) IFactorContext
	AllSTAR() []antlr.TerminalNode
	STAR(i int) antlr.TerminalNode
	AllSLASH() []antlr.TerminalNode
	SLASH(i int) antlr.TerminalNode

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KotlinFunctionParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFactorContext); ok {
			len++
		}
	}

	tst := make([]IFactorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFactorContext); ok {
			tst[i] = t.(IFactorContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) AllSTAR() []antlr.TerminalNode {
	return s.GetTokens(KotlinFunctionParserSTAR)
}

func (s *TermContext) STAR(i int) antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserSTAR, i)
}

func (s *TermContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(KotlinFunctionParserSLASH)
}

func (s *TermContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserSLASH, i)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *KotlinFunctionParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, KotlinFunctionParserRULE_term)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Factor()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == KotlinFunctionParserSTAR || _la == KotlinFunctionParserSLASH {
		{
			p.SetState(65)
			_la = p.GetTokenStream().LA(1)

			if !(_la == KotlinFunctionParserSTAR || _la == KotlinFunctionParserSLASH) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(66)
			p.Factor()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	FLOAT_LITERAL() antlr.TerminalNode
	INTEGER_LITERAL() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KotlinFunctionParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KotlinFunctionParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserIDENTIFIER, 0)
}

func (s *FactorContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserFLOAT_LITERAL, 0)
}

func (s *FactorContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserINTEGER_LITERAL, 0)
}

func (s *FactorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserLPAREN, 0)
}

func (s *FactorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FactorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(KotlinFunctionParserRPAREN, 0)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *KotlinFunctionParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, KotlinFunctionParserRULE_factor)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case KotlinFunctionParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Match(KotlinFunctionParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case KotlinFunctionParserFLOAT_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Match(KotlinFunctionParserFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case KotlinFunctionParserINTEGER_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(74)
			p.Match(KotlinFunctionParserINTEGER_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case KotlinFunctionParserLPAREN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(75)
			p.Match(KotlinFunctionParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.Expr()
		}
		{
			p.SetState(77)
			p.Match(KotlinFunctionParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
