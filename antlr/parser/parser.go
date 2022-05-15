// Code generated from ./Parser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Parser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type Parser struct {
	*antlr.BaseParser
}

var parserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func parserParserInit() {
	staticData := &parserParserStaticData
	staticData.literalNames = []string{
		"", "", "'component'", "", "'$'", "'{'", "'}'", "'('", "')'", "','",
		"':'", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "STRING", "COMPONENT", "COMMENT", "DOLLAR", "LBRACE", "RBRACE",
		"LPAREN", "RPAREN", "COMMA", "COLON", "DOT", "IDEN", "ID", "INT", "FLOAT",
		"BOOLEAN", "WS",
	}
	staticData.ruleNames = []string{
		"document", "componentDef", "variableDefList", "variableDef", "variable",
		"selectionSet", "tagDef", "argumentList", "argument",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 83, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 4, 0, 20, 8, 0,
		11, 0, 12, 0, 21, 1, 0, 3, 0, 25, 8, 0, 1, 1, 1, 1, 1, 1, 3, 1, 30, 8,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 38, 8, 2, 10, 2, 12, 2, 41,
		9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 4, 5, 55, 8, 5, 11, 5, 12, 5, 56, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 63,
		8, 6, 1, 6, 3, 6, 66, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 72, 8, 7, 10,
		7, 12, 7, 75, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 0, 0, 9,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 0, 82, 0, 24, 1, 0, 0, 0, 2, 26, 1, 0,
		0, 0, 4, 33, 1, 0, 0, 0, 6, 44, 1, 0, 0, 0, 8, 48, 1, 0, 0, 0, 10, 51,
		1, 0, 0, 0, 12, 60, 1, 0, 0, 0, 14, 67, 1, 0, 0, 0, 16, 78, 1, 0, 0, 0,
		18, 20, 3, 2, 1, 0, 19, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 19, 1,
		0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 25, 1, 0, 0, 0, 23, 25, 5, 0, 0, 1, 24,
		19, 1, 0, 0, 0, 24, 23, 1, 0, 0, 0, 25, 1, 1, 0, 0, 0, 26, 27, 5, 2, 0,
		0, 27, 29, 5, 12, 0, 0, 28, 30, 3, 4, 2, 0, 29, 28, 1, 0, 0, 0, 29, 30,
		1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 32, 3, 10, 5, 0, 32, 3, 1, 0, 0, 0,
		33, 34, 5, 7, 0, 0, 34, 39, 3, 6, 3, 0, 35, 36, 5, 9, 0, 0, 36, 38, 3,
		6, 3, 0, 37, 35, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39,
		40, 1, 0, 0, 0, 40, 42, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 43, 5, 8, 0,
		0, 43, 5, 1, 0, 0, 0, 44, 45, 3, 8, 4, 0, 45, 46, 5, 10, 0, 0, 46, 47,
		5, 12, 0, 0, 47, 7, 1, 0, 0, 0, 48, 49, 5, 4, 0, 0, 49, 50, 5, 12, 0, 0,
		50, 9, 1, 0, 0, 0, 51, 54, 5, 5, 0, 0, 52, 55, 3, 12, 6, 0, 53, 55, 5,
		1, 0, 0, 54, 52, 1, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56,
		54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 5, 6, 0,
		0, 59, 11, 1, 0, 0, 0, 60, 62, 5, 12, 0, 0, 61, 63, 3, 14, 7, 0, 62, 61,
		1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 65, 1, 0, 0, 0, 64, 66, 3, 10, 5, 0,
		65, 64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 13, 1, 0, 0, 0, 67, 68, 5,
		7, 0, 0, 68, 73, 3, 16, 8, 0, 69, 70, 5, 9, 0, 0, 70, 72, 3, 16, 8, 0,
		71, 69, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1,
		0, 0, 0, 74, 76, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 77, 5, 8, 0, 0, 77,
		15, 1, 0, 0, 0, 78, 79, 5, 12, 0, 0, 79, 80, 5, 10, 0, 0, 80, 81, 5, 1,
		0, 0, 81, 17, 1, 0, 0, 0, 9, 21, 24, 29, 39, 54, 56, 62, 65, 73,
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

// ParserInit initializes any static state used to implement Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ParserInit() {
	staticData := &parserParserStaticData
	staticData.once.Do(parserParserInit)
}

// NewParser produces a new parser instance for the optional input antlr.TokenStream.
func NewParser(input antlr.TokenStream) *Parser {
	ParserInit()
	this := new(Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &parserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Parser.g4"

	return this
}

// Parser tokens.
const (
	ParserEOF       = antlr.TokenEOF
	ParserSTRING    = 1
	ParserCOMPONENT = 2
	ParserCOMMENT   = 3
	ParserDOLLAR    = 4
	ParserLBRACE    = 5
	ParserRBRACE    = 6
	ParserLPAREN    = 7
	ParserRPAREN    = 8
	ParserCOMMA     = 9
	ParserCOLON     = 10
	ParserDOT       = 11
	ParserIDEN      = 12
	ParserID        = 13
	ParserINT       = 14
	ParserFLOAT     = 15
	ParserBOOLEAN   = 16
	ParserWS        = 17
)

// Parser rules.
const (
	ParserRULE_document        = 0
	ParserRULE_componentDef    = 1
	ParserRULE_variableDefList = 2
	ParserRULE_variableDef     = 3
	ParserRULE_variable        = 4
	ParserRULE_selectionSet    = 5
	ParserRULE_tagDef          = 6
	ParserRULE_argumentList    = 7
	ParserRULE_argument        = 8
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) AllComponentDef() []IComponentDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComponentDefContext); ok {
			len++
		}
	}

	tst := make([]IComponentDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComponentDefContext); ok {
			tst[i] = t.(IComponentDefContext)
			i++
		}
	}

	return tst
}

func (s *DocumentContext) ComponentDef(i int) IComponentDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComponentDefContext); ok {
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

	return t.(IComponentDefContext)
}

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(ParserEOF, 0)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *Parser) Document() (localctx IDocumentContext) {
	this := p
	_ = this

	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParserRULE_document)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(24)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ParserCOMPONENT:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ParserCOMPONENT {
			{
				p.SetState(18)
				p.ComponentDef()
			}

			p.SetState(21)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case ParserEOF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Match(ParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IComponentDefContext is an interface to support dynamic dispatch.
type IComponentDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsComponentDefContext differentiates from other interfaces.
	IsComponentDefContext()
}

type ComponentDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyComponentDefContext() *ComponentDefContext {
	var p = new(ComponentDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParserRULE_componentDef
	return p
}

func (*ComponentDefContext) IsComponentDefContext() {}

func NewComponentDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentDefContext {
	var p = new(ComponentDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserRULE_componentDef

	return p
}

func (s *ComponentDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentDefContext) GetName() antlr.Token { return s.name }

func (s *ComponentDefContext) SetName(v antlr.Token) { s.name = v }

func (s *ComponentDefContext) COMPONENT() antlr.TerminalNode {
	return s.GetToken(ParserCOMPONENT, 0)
}

func (s *ComponentDefContext) SelectionSet() ISelectionSetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectionSetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *ComponentDefContext) IDEN() antlr.TerminalNode {
	return s.GetToken(ParserIDEN, 0)
}

func (s *ComponentDefContext) VariableDefList() IVariableDefListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDefListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDefListContext)
}

func (s *ComponentDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterComponentDef(s)
	}
}

func (s *ComponentDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitComponentDef(s)
	}
}

func (p *Parser) ComponentDef() (localctx IComponentDefContext) {
	this := p
	_ = this

	localctx = NewComponentDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ParserRULE_componentDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Match(ParserCOMPONENT)
	}
	{
		p.SetState(27)

		var _m = p.Match(ParserIDEN)

		localctx.(*ComponentDefContext).name = _m
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ParserLPAREN {
		{
			p.SetState(28)
			p.VariableDefList()
		}

	}
	{
		p.SetState(31)
		p.SelectionSet()
	}

	return localctx
}

// IVariableDefListContext is an interface to support dynamic dispatch.
type IVariableDefListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_variableDef returns the _variableDef rule contexts.
	Get_variableDef() IVariableDefContext

	// Set_variableDef sets the _variableDef rule contexts.
	Set_variableDef(IVariableDefContext)

	// GetVariables returns the variables rule context list.
	GetVariables() []IVariableDefContext

	// SetVariables sets the variables rule context list.
	SetVariables([]IVariableDefContext)

	// IsVariableDefListContext differentiates from other interfaces.
	IsVariableDefListContext()
}

type VariableDefListContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	_variableDef IVariableDefContext
	variables    []IVariableDefContext
}

func NewEmptyVariableDefListContext() *VariableDefListContext {
	var p = new(VariableDefListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParserRULE_variableDefList
	return p
}

func (*VariableDefListContext) IsVariableDefListContext() {}

func NewVariableDefListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefListContext {
	var p = new(VariableDefListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserRULE_variableDefList

	return p
}

func (s *VariableDefListContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefListContext) Get_variableDef() IVariableDefContext { return s._variableDef }

func (s *VariableDefListContext) Set_variableDef(v IVariableDefContext) { s._variableDef = v }

func (s *VariableDefListContext) GetVariables() []IVariableDefContext { return s.variables }

func (s *VariableDefListContext) SetVariables(v []IVariableDefContext) { s.variables = v }

func (s *VariableDefListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ParserLPAREN, 0)
}

func (s *VariableDefListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ParserRPAREN, 0)
}

func (s *VariableDefListContext) AllVariableDef() []IVariableDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDefContext); ok {
			len++
		}
	}

	tst := make([]IVariableDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDefContext); ok {
			tst[i] = t.(IVariableDefContext)
			i++
		}
	}

	return tst
}

func (s *VariableDefListContext) VariableDef(i int) IVariableDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDefContext); ok {
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

	return t.(IVariableDefContext)
}

func (s *VariableDefListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ParserCOMMA)
}

func (s *VariableDefListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ParserCOMMA, i)
}

func (s *VariableDefListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterVariableDefList(s)
	}
}

func (s *VariableDefListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitVariableDefList(s)
	}
}

func (p *Parser) VariableDefList() (localctx IVariableDefListContext) {
	this := p
	_ = this

	localctx = NewVariableDefListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParserRULE_variableDefList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(ParserLPAREN)
	}
	{
		p.SetState(34)

		var _x = p.VariableDef()

		localctx.(*VariableDefListContext)._variableDef = _x
	}
	localctx.(*VariableDefListContext).variables = append(localctx.(*VariableDefListContext).variables, localctx.(*VariableDefListContext)._variableDef)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ParserCOMMA {
		{
			p.SetState(35)
			p.Match(ParserCOMMA)
		}
		{
			p.SetState(36)

			var _x = p.VariableDef()

			localctx.(*VariableDefListContext)._variableDef = _x
		}
		localctx.(*VariableDefListContext).variables = append(localctx.(*VariableDefListContext).variables, localctx.(*VariableDefListContext)._variableDef)

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(42)
		p.Match(ParserRPAREN)
	}

	return localctx
}

// IVariableDefContext is an interface to support dynamic dispatch.
type IVariableDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDataType returns the dataType token.
	GetDataType() antlr.Token

	// SetDataType sets the dataType token.
	SetDataType(antlr.Token)

	// IsVariableDefContext differentiates from other interfaces.
	IsVariableDefContext()
}

type VariableDefContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	dataType antlr.Token
}

func NewEmptyVariableDefContext() *VariableDefContext {
	var p = new(VariableDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParserRULE_variableDef
	return p
}

func (*VariableDefContext) IsVariableDefContext() {}

func NewVariableDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefContext {
	var p = new(VariableDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserRULE_variableDef

	return p
}

func (s *VariableDefContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefContext) GetDataType() antlr.Token { return s.dataType }

func (s *VariableDefContext) SetDataType(v antlr.Token) { s.dataType = v }

func (s *VariableDefContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VariableDefContext) COLON() antlr.TerminalNode {
	return s.GetToken(ParserCOLON, 0)
}

func (s *VariableDefContext) IDEN() antlr.TerminalNode {
	return s.GetToken(ParserIDEN, 0)
}

func (s *VariableDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterVariableDef(s)
	}
}

func (s *VariableDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitVariableDef(s)
	}
}

func (p *Parser) VariableDef() (localctx IVariableDefContext) {
	this := p
	_ = this

	localctx = NewVariableDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ParserRULE_variableDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Variable()
	}
	{
		p.SetState(45)
		p.Match(ParserCOLON)
	}
	{
		p.SetState(46)

		var _m = p.Match(ParserIDEN)

		localctx.(*VariableDefContext).dataType = _m
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(ParserDOLLAR, 0)
}

func (s *VariableContext) IDEN() antlr.TerminalNode {
	return s.GetToken(ParserIDEN, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *Parser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(ParserDOLLAR)
	}
	{
		p.SetState(49)
		p.Match(ParserIDEN)
	}

	return localctx
}

// ISelectionSetContext is an interface to support dynamic dispatch.
type ISelectionSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionSetContext differentiates from other interfaces.
	IsSelectionSetContext()
}

type SelectionSetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionSetContext() *SelectionSetContext {
	var p = new(SelectionSetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParserRULE_selectionSet
	return p
}

func (*SelectionSetContext) IsSelectionSetContext() {}

func NewSelectionSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionSetContext {
	var p = new(SelectionSetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserRULE_selectionSet

	return p
}

func (s *SelectionSetContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionSetContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ParserLBRACE, 0)
}

func (s *SelectionSetContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ParserRBRACE, 0)
}

func (s *SelectionSetContext) AllTagDef() []ITagDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITagDefContext); ok {
			len++
		}
	}

	tst := make([]ITagDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITagDefContext); ok {
			tst[i] = t.(ITagDefContext)
			i++
		}
	}

	return tst
}

func (s *SelectionSetContext) TagDef(i int) ITagDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagDefContext); ok {
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

	return t.(ITagDefContext)
}

func (s *SelectionSetContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ParserSTRING)
}

func (s *SelectionSetContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ParserSTRING, i)
}

func (s *SelectionSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterSelectionSet(s)
	}
}

func (s *SelectionSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitSelectionSet(s)
	}
}

func (p *Parser) SelectionSet() (localctx ISelectionSetContext) {
	this := p
	_ = this

	localctx = NewSelectionSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ParserRULE_selectionSet)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(ParserLBRACE)
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ParserSTRING || _la == ParserIDEN {
		p.SetState(54)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ParserIDEN:
			{
				p.SetState(52)
				p.TagDef()
			}

		case ParserSTRING:
			{
				p.SetState(53)
				p.Match(ParserSTRING)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(58)
		p.Match(ParserRBRACE)
	}

	return localctx
}

// ITagDefContext is an interface to support dynamic dispatch.
type ITagDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagDefContext differentiates from other interfaces.
	IsTagDefContext()
}

type TagDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagDefContext() *TagDefContext {
	var p = new(TagDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParserRULE_tagDef
	return p
}

func (*TagDefContext) IsTagDefContext() {}

func NewTagDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagDefContext {
	var p = new(TagDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserRULE_tagDef

	return p
}

func (s *TagDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TagDefContext) IDEN() antlr.TerminalNode {
	return s.GetToken(ParserIDEN, 0)
}

func (s *TagDefContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *TagDefContext) SelectionSet() ISelectionSetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectionSetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *TagDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterTagDef(s)
	}
}

func (s *TagDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitTagDef(s)
	}
}

func (p *Parser) TagDef() (localctx ITagDefContext) {
	this := p
	_ = this

	localctx = NewTagDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ParserRULE_tagDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(ParserIDEN)
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ParserLPAREN {
		{
			p.SetState(61)
			p.ArgumentList()
		}

	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ParserLBRACE {
		{
			p.SetState(64)
			p.SelectionSet()
		}

	}

	return localctx
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_argument returns the _argument rule contexts.
	Get_argument() IArgumentContext

	// Set_argument sets the _argument rule contexts.
	Set_argument(IArgumentContext)

	// GetArguments returns the arguments rule context list.
	GetArguments() []IArgumentContext

	// SetArguments sets the arguments rule context list.
	SetArguments([]IArgumentContext)

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	_argument IArgumentContext
	arguments []IArgumentContext
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParserRULE_argumentList
	return p
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) Get_argument() IArgumentContext { return s._argument }

func (s *ArgumentListContext) Set_argument(v IArgumentContext) { s._argument = v }

func (s *ArgumentListContext) GetArguments() []IArgumentContext { return s.arguments }

func (s *ArgumentListContext) SetArguments(v []IArgumentContext) { s.arguments = v }

func (s *ArgumentListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ParserLPAREN, 0)
}

func (s *ArgumentListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ParserRPAREN, 0)
}

func (s *ArgumentListContext) AllArgument() []IArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentContext); ok {
			len++
		}
	}

	tst := make([]IArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentContext); ok {
			tst[i] = t.(IArgumentContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentListContext) Argument(i int) IArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
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

	return t.(IArgumentContext)
}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (p *Parser) ArgumentList() (localctx IArgumentListContext) {
	this := p
	_ = this

	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ParserRULE_argumentList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(ParserLPAREN)
	}
	{
		p.SetState(68)

		var _x = p.Argument()

		localctx.(*ArgumentListContext)._argument = _x
	}
	localctx.(*ArgumentListContext).arguments = append(localctx.(*ArgumentListContext).arguments, localctx.(*ArgumentListContext)._argument)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ParserCOMMA {
		{
			p.SetState(69)
			p.Match(ParserCOMMA)
		}
		{
			p.SetState(70)

			var _x = p.Argument()

			localctx.(*ArgumentListContext)._argument = _x
		}
		localctx.(*ArgumentListContext).arguments = append(localctx.(*ArgumentListContext).arguments, localctx.(*ArgumentListContext)._argument)

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(76)
		p.Match(ParserRPAREN)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) IDEN() antlr.TerminalNode {
	return s.GetToken(ParserIDEN, 0)
}

func (s *ArgumentContext) COLON() antlr.TerminalNode {
	return s.GetToken(ParserCOLON, 0)
}

func (s *ArgumentContext) STRING() antlr.TerminalNode {
	return s.GetToken(ParserSTRING, 0)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *Parser) Argument() (localctx IArgumentContext) {
	this := p
	_ = this

	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ParserRULE_argument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(ParserIDEN)
	}
	{
		p.SetState(79)
		p.Match(ParserCOLON)
	}
	{
		p.SetState(80)
		p.Match(ParserSTRING)
	}

	return localctx
}
