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
		"", "", "'module'", "'component'", "", "'$'", "'{'", "'}'", "'('", "')'",
		"','", "':'", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "STRING", "MODULE", "COMPONENT", "COMMENT", "DOLLAR", "LBRACE",
		"RBRACE", "LPAREN", "RPAREN", "COMMA", "COLON", "DOT", "IDEN", "ID",
		"INT", "FLOAT", "BOOLEAN", "WS",
	}
	staticData.ruleNames = []string{
		"file", "moduleDef", "componentDef", "variableDefList", "variableDef",
		"variable", "selectionSet", "selection", "tagDef", "argumentList", "argument",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 94, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 4, 0, 25, 8, 0, 11, 0, 12, 0, 26, 1, 0, 3, 0, 30, 8, 0,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 3, 2, 38, 8, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 5, 3, 46, 8, 3, 10, 3, 12, 3, 49, 9, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 4, 6, 62, 8, 6, 11, 6,
		12, 6, 63, 1, 6, 1, 6, 1, 7, 1, 7, 3, 7, 70, 8, 7, 1, 8, 1, 8, 3, 8, 74,
		8, 8, 1, 8, 3, 8, 77, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 83, 8, 9, 10,
		9, 12, 9, 86, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 0, 0,
		11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 0, 91, 0, 29, 1, 0, 0, 0,
		2, 31, 1, 0, 0, 0, 4, 34, 1, 0, 0, 0, 6, 41, 1, 0, 0, 0, 8, 52, 1, 0, 0,
		0, 10, 56, 1, 0, 0, 0, 12, 59, 1, 0, 0, 0, 14, 69, 1, 0, 0, 0, 16, 71,
		1, 0, 0, 0, 18, 78, 1, 0, 0, 0, 20, 89, 1, 0, 0, 0, 22, 24, 3, 2, 1, 0,
		23, 25, 3, 4, 2, 0, 24, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 24, 1,
		0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 30, 1, 0, 0, 0, 28, 30, 5, 0, 0, 1, 29,
		22, 1, 0, 0, 0, 29, 28, 1, 0, 0, 0, 30, 1, 1, 0, 0, 0, 31, 32, 5, 2, 0,
		0, 32, 33, 5, 13, 0, 0, 33, 3, 1, 0, 0, 0, 34, 35, 5, 3, 0, 0, 35, 37,
		5, 13, 0, 0, 36, 38, 3, 6, 3, 0, 37, 36, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0,
		38, 39, 1, 0, 0, 0, 39, 40, 3, 12, 6, 0, 40, 5, 1, 0, 0, 0, 41, 42, 5,
		8, 0, 0, 42, 47, 3, 8, 4, 0, 43, 44, 5, 10, 0, 0, 44, 46, 3, 8, 4, 0, 45,
		43, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0,
		0, 48, 50, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 50, 51, 5, 9, 0, 0, 51, 7, 1,
		0, 0, 0, 52, 53, 3, 10, 5, 0, 53, 54, 5, 11, 0, 0, 54, 55, 5, 13, 0, 0,
		55, 9, 1, 0, 0, 0, 56, 57, 5, 5, 0, 0, 57, 58, 5, 13, 0, 0, 58, 11, 1,
		0, 0, 0, 59, 61, 5, 6, 0, 0, 60, 62, 3, 14, 7, 0, 61, 60, 1, 0, 0, 0, 62,
		63, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 65, 1, 0, 0,
		0, 65, 66, 5, 7, 0, 0, 66, 13, 1, 0, 0, 0, 67, 70, 5, 1, 0, 0, 68, 70,
		3, 16, 8, 0, 69, 67, 1, 0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 15, 1, 0, 0, 0,
		71, 73, 5, 13, 0, 0, 72, 74, 3, 18, 9, 0, 73, 72, 1, 0, 0, 0, 73, 74, 1,
		0, 0, 0, 74, 76, 1, 0, 0, 0, 75, 77, 3, 12, 6, 0, 76, 75, 1, 0, 0, 0, 76,
		77, 1, 0, 0, 0, 77, 17, 1, 0, 0, 0, 78, 79, 5, 8, 0, 0, 79, 84, 3, 20,
		10, 0, 80, 81, 5, 10, 0, 0, 81, 83, 3, 20, 10, 0, 82, 80, 1, 0, 0, 0, 83,
		86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 1, 0, 0,
		0, 86, 84, 1, 0, 0, 0, 87, 88, 5, 9, 0, 0, 88, 19, 1, 0, 0, 0, 89, 90,
		5, 13, 0, 0, 90, 91, 5, 11, 0, 0, 91, 92, 5, 1, 0, 0, 92, 21, 1, 0, 0,
		0, 9, 26, 29, 37, 47, 63, 69, 73, 76, 84,
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
	ParserMODULE    = 2
	ParserCOMPONENT = 3
	ParserCOMMENT   = 4
	ParserDOLLAR    = 5
	ParserLBRACE    = 6
	ParserRBRACE    = 7
	ParserLPAREN    = 8
	ParserRPAREN    = 9
	ParserCOMMA     = 10
	ParserCOLON     = 11
	ParserDOT       = 12
	ParserIDEN      = 13
	ParserID        = 14
	ParserINT       = 15
	ParserFLOAT     = 16
	ParserBOOLEAN   = 17
	ParserWS        = 18
)

// Parser rules.
const (
	ParserRULE_file            = 0
	ParserRULE_moduleDef       = 1
	ParserRULE_componentDef    = 2
	ParserRULE_variableDefList = 3
	ParserRULE_variableDef     = 4
	ParserRULE_variable        = 5
	ParserRULE_selectionSet    = 6
	ParserRULE_selection       = 7
	ParserRULE_tagDef          = 8
	ParserRULE_argumentList    = 9
	ParserRULE_argument        = 10
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) ModuleDef() IModuleDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModuleDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModuleDefContext)
}

func (s *FileContext) AllComponentDef() []IComponentDefContext {
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

func (s *FileContext) ComponentDef(i int) IComponentDefContext {
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

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(ParserEOF, 0)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitFile(s)
	}
}

func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ParserVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Parser) File() (localctx IFileContext) {
	this := p
	_ = this

	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParserRULE_file)
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

	p.SetState(29)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ParserMODULE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.ModuleDef()
		}
		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ParserCOMPONENT {
			{
				p.SetState(23)
				p.ComponentDef()
			}

			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case ParserEOF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(28)
			p.Match(ParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IModuleDefContext is an interface to support dynamic dispatch.
type IModuleDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsModuleDefContext differentiates from other interfaces.
	IsModuleDefContext()
}

type ModuleDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyModuleDefContext() *ModuleDefContext {
	var p = new(ModuleDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParserRULE_moduleDef
	return p
}

func (*ModuleDefContext) IsModuleDefContext() {}

func NewModuleDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleDefContext {
	var p = new(ModuleDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserRULE_moduleDef

	return p
}

func (s *ModuleDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleDefContext) GetName() antlr.Token { return s.name }

func (s *ModuleDefContext) SetName(v antlr.Token) { s.name = v }

func (s *ModuleDefContext) MODULE() antlr.TerminalNode {
	return s.GetToken(ParserMODULE, 0)
}

func (s *ModuleDefContext) IDEN() antlr.TerminalNode {
	return s.GetToken(ParserIDEN, 0)
}

func (s *ModuleDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterModuleDef(s)
	}
}

func (s *ModuleDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitModuleDef(s)
	}
}

func (s *ModuleDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ParserVisitor:
		return t.VisitModuleDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Parser) ModuleDef() (localctx IModuleDefContext) {
	this := p
	_ = this

	localctx = NewModuleDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ParserRULE_moduleDef)

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
		p.SetState(31)
		p.Match(ParserMODULE)
	}
	{
		p.SetState(32)

		var _m = p.Match(ParserIDEN)

		localctx.(*ModuleDefContext).name = _m
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

func (s *ComponentDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ParserVisitor:
		return t.VisitComponentDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Parser) ComponentDef() (localctx IComponentDefContext) {
	this := p
	_ = this

	localctx = NewComponentDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParserRULE_componentDef)
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
		p.SetState(34)
		p.Match(ParserCOMPONENT)
	}
	{
		p.SetState(35)

		var _m = p.Match(ParserIDEN)

		localctx.(*ComponentDefContext).name = _m
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ParserLPAREN {
		{
			p.SetState(36)
			p.VariableDefList()
		}

	}
	{
		p.SetState(39)
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

func (s *VariableDefListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ParserVisitor:
		return t.VisitVariableDefList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Parser) VariableDefList() (localctx IVariableDefListContext) {
	this := p
	_ = this

	localctx = NewVariableDefListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ParserRULE_variableDefList)
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
		p.SetState(41)
		p.Match(ParserLPAREN)
	}
	{
		p.SetState(42)

		var _x = p.VariableDef()

		localctx.(*VariableDefListContext)._variableDef = _x
	}
	localctx.(*VariableDefListContext).variables = append(localctx.(*VariableDefListContext).variables, localctx.(*VariableDefListContext)._variableDef)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ParserCOMMA {
		{
			p.SetState(43)
			p.Match(ParserCOMMA)
		}
		{
			p.SetState(44)

			var _x = p.VariableDef()

			localctx.(*VariableDefListContext)._variableDef = _x
		}
		localctx.(*VariableDefListContext).variables = append(localctx.(*VariableDefListContext).variables, localctx.(*VariableDefListContext)._variableDef)

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
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

func (s *VariableDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ParserVisitor:
		return t.VisitVariableDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Parser) VariableDef() (localctx IVariableDefContext) {
	this := p
	_ = this

	localctx = NewVariableDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ParserRULE_variableDef)

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
		p.SetState(52)
		p.Variable()
	}
	{
		p.SetState(53)
		p.Match(ParserCOLON)
	}
	{
		p.SetState(54)

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

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ParserVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Parser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ParserRULE_variable)

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
		p.SetState(56)
		p.Match(ParserDOLLAR)
	}
	{
		p.SetState(57)
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

func (s *SelectionSetContext) AllSelection() []ISelectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectionContext); ok {
			len++
		}
	}

	tst := make([]ISelectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectionContext); ok {
			tst[i] = t.(ISelectionContext)
			i++
		}
	}

	return tst
}

func (s *SelectionSetContext) Selection(i int) ISelectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectionContext); ok {
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

	return t.(ISelectionContext)
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

func (s *SelectionSetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ParserVisitor:
		return t.VisitSelectionSet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Parser) SelectionSet() (localctx ISelectionSetContext) {
	this := p
	_ = this

	localctx = NewSelectionSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ParserRULE_selectionSet)
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
		p.SetState(59)
		p.Match(ParserLBRACE)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ParserSTRING || _la == ParserIDEN {
		{
			p.SetState(60)
			p.Selection()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(65)
		p.Match(ParserRBRACE)
	}

	return localctx
}

// ISelectionContext is an interface to support dynamic dispatch.
type ISelectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionContext differentiates from other interfaces.
	IsSelectionContext()
}

type SelectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionContext() *SelectionContext {
	var p = new(SelectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParserRULE_selection
	return p
}

func (*SelectionContext) IsSelectionContext() {}

func NewSelectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionContext {
	var p = new(SelectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserRULE_selection

	return p
}

func (s *SelectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionContext) CopyFrom(ctx *SelectionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SelectTagContext struct {
	*SelectionContext
}

func NewSelectTagContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectTagContext {
	var p = new(SelectTagContext)

	p.SelectionContext = NewEmptySelectionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SelectionContext))

	return p
}

func (s *SelectTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectTagContext) TagDef() ITagDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagDefContext)
}

func (s *SelectTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterSelectTag(s)
	}
}

func (s *SelectTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitSelectTag(s)
	}
}

func (s *SelectTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ParserVisitor:
		return t.VisitSelectTag(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelectStringContext struct {
	*SelectionContext
}

func NewSelectStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectStringContext {
	var p = new(SelectStringContext)

	p.SelectionContext = NewEmptySelectionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SelectionContext))

	return p
}

func (s *SelectStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(ParserSTRING, 0)
}

func (s *SelectStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterSelectString(s)
	}
}

func (s *SelectStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitSelectString(s)
	}
}

func (s *SelectStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ParserVisitor:
		return t.VisitSelectString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Parser) Selection() (localctx ISelectionContext) {
	this := p
	_ = this

	localctx = NewSelectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ParserRULE_selection)

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

	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ParserSTRING:
		localctx = NewSelectStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Match(ParserSTRING)
		}

	case ParserIDEN:
		localctx = NewSelectTagContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.TagDef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *TagDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ParserVisitor:
		return t.VisitTagDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Parser) TagDef() (localctx ITagDefContext) {
	this := p
	_ = this

	localctx = NewTagDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ParserRULE_tagDef)
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
		p.SetState(71)
		p.Match(ParserIDEN)
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ParserLPAREN {
		{
			p.SetState(72)
			p.ArgumentList()
		}

	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ParserLBRACE {
		{
			p.SetState(75)
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

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ParserVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Parser) ArgumentList() (localctx IArgumentListContext) {
	this := p
	_ = this

	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ParserRULE_argumentList)
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
		p.SetState(78)
		p.Match(ParserLPAREN)
	}
	{
		p.SetState(79)

		var _x = p.Argument()

		localctx.(*ArgumentListContext)._argument = _x
	}
	localctx.(*ArgumentListContext).arguments = append(localctx.(*ArgumentListContext).arguments, localctx.(*ArgumentListContext)._argument)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ParserCOMMA {
		{
			p.SetState(80)
			p.Match(ParserCOMMA)
		}
		{
			p.SetState(81)

			var _x = p.Argument()

			localctx.(*ArgumentListContext)._argument = _x
		}
		localctx.(*ArgumentListContext).arguments = append(localctx.(*ArgumentListContext).arguments, localctx.(*ArgumentListContext)._argument)

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(ParserRPAREN)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
	value  antlr.Token
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

func (s *ArgumentContext) GetKey() antlr.Token { return s.key }

func (s *ArgumentContext) GetValue() antlr.Token { return s.value }

func (s *ArgumentContext) SetKey(v antlr.Token) { s.key = v }

func (s *ArgumentContext) SetValue(v antlr.Token) { s.value = v }

func (s *ArgumentContext) COLON() antlr.TerminalNode {
	return s.GetToken(ParserCOLON, 0)
}

func (s *ArgumentContext) IDEN() antlr.TerminalNode {
	return s.GetToken(ParserIDEN, 0)
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

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ParserVisitor:
		return t.VisitArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Parser) Argument() (localctx IArgumentContext) {
	this := p
	_ = this

	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ParserRULE_argument)

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
		p.SetState(89)

		var _m = p.Match(ParserIDEN)

		localctx.(*ArgumentContext).key = _m
	}
	{
		p.SetState(90)
		p.Match(ParserCOLON)
	}
	{
		p.SetState(91)

		var _m = p.Match(ParserSTRING)

		localctx.(*ArgumentContext).value = _m
	}

	return localctx
}
