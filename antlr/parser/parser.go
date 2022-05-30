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
		"", "", "", "", "", "", "", "'$'", "'{'", "'}'", "'('", "')'", "','",
		"':'", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "STRING", "MODULE", "COMPONENT", "FROM", "IMPORT", "COMMENT", "DOLLAR",
		"LBRACE", "RBRACE", "LPAREN", "RPAREN", "COMMA", "COLON", "DOT", "IDEN",
		"ID", "INT", "FLOAT", "BOOLEAN", "WS",
	}
	staticData.ruleNames = []string{
		"file", "moduleDef", "importDef", "componentDef", "variableDefList",
		"variableDef", "variable", "selectionSet", "selection", "tagDef", "argumentList",
		"argument",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 20, 113, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 5, 0, 27, 8, 0, 10, 0, 12, 0, 30, 9, 0, 1,
		0, 4, 0, 33, 8, 0, 11, 0, 12, 0, 34, 1, 0, 3, 0, 38, 8, 0, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 49, 8, 2, 10, 2, 12, 2,
		52, 9, 2, 1, 3, 1, 3, 1, 3, 3, 3, 57, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 5, 4, 65, 8, 4, 10, 4, 12, 4, 68, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 4, 7, 81, 8, 7, 11, 7, 12, 7,
		82, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 89, 8, 8, 1, 9, 1, 9, 3, 9, 93, 8, 9,
		1, 9, 3, 9, 96, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 102, 8, 10, 10,
		10, 12, 10, 105, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		0, 0, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 0, 111, 0, 37,
		1, 0, 0, 0, 2, 39, 1, 0, 0, 0, 4, 42, 1, 0, 0, 0, 6, 53, 1, 0, 0, 0, 8,
		60, 1, 0, 0, 0, 10, 71, 1, 0, 0, 0, 12, 75, 1, 0, 0, 0, 14, 78, 1, 0, 0,
		0, 16, 88, 1, 0, 0, 0, 18, 90, 1, 0, 0, 0, 20, 97, 1, 0, 0, 0, 22, 108,
		1, 0, 0, 0, 24, 28, 3, 2, 1, 0, 25, 27, 3, 4, 2, 0, 26, 25, 1, 0, 0, 0,
		27, 30, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 32, 1,
		0, 0, 0, 30, 28, 1, 0, 0, 0, 31, 33, 3, 6, 3, 0, 32, 31, 1, 0, 0, 0, 33,
		34, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 38, 1, 0, 0,
		0, 36, 38, 5, 0, 0, 1, 37, 24, 1, 0, 0, 0, 37, 36, 1, 0, 0, 0, 38, 1, 1,
		0, 0, 0, 39, 40, 5, 2, 0, 0, 40, 41, 5, 15, 0, 0, 41, 3, 1, 0, 0, 0, 42,
		43, 5, 4, 0, 0, 43, 44, 5, 15, 0, 0, 44, 45, 5, 5, 0, 0, 45, 50, 5, 15,
		0, 0, 46, 47, 5, 12, 0, 0, 47, 49, 5, 15, 0, 0, 48, 46, 1, 0, 0, 0, 49,
		52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 5, 1, 0, 0,
		0, 52, 50, 1, 0, 0, 0, 53, 54, 5, 3, 0, 0, 54, 56, 5, 15, 0, 0, 55, 57,
		3, 8, 4, 0, 56, 55, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0,
		58, 59, 3, 14, 7, 0, 59, 7, 1, 0, 0, 0, 60, 61, 5, 10, 0, 0, 61, 66, 3,
		10, 5, 0, 62, 63, 5, 12, 0, 0, 63, 65, 3, 10, 5, 0, 64, 62, 1, 0, 0, 0,
		65, 68, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 69, 1,
		0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 70, 5, 11, 0, 0, 70, 9, 1, 0, 0, 0, 71,
		72, 3, 12, 6, 0, 72, 73, 5, 13, 0, 0, 73, 74, 5, 15, 0, 0, 74, 11, 1, 0,
		0, 0, 75, 76, 5, 7, 0, 0, 76, 77, 5, 15, 0, 0, 77, 13, 1, 0, 0, 0, 78,
		80, 5, 8, 0, 0, 79, 81, 3, 16, 8, 0, 80, 79, 1, 0, 0, 0, 81, 82, 1, 0,
		0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85,
		5, 9, 0, 0, 85, 15, 1, 0, 0, 0, 86, 89, 5, 1, 0, 0, 87, 89, 3, 18, 9, 0,
		88, 86, 1, 0, 0, 0, 88, 87, 1, 0, 0, 0, 89, 17, 1, 0, 0, 0, 90, 92, 5,
		15, 0, 0, 91, 93, 3, 20, 10, 0, 92, 91, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0,
		93, 95, 1, 0, 0, 0, 94, 96, 3, 14, 7, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1,
		0, 0, 0, 96, 19, 1, 0, 0, 0, 97, 98, 5, 10, 0, 0, 98, 103, 3, 22, 11, 0,
		99, 100, 5, 12, 0, 0, 100, 102, 3, 22, 11, 0, 101, 99, 1, 0, 0, 0, 102,
		105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 106,
		1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 107, 5, 11, 0, 0, 107, 21, 1, 0,
		0, 0, 108, 109, 5, 15, 0, 0, 109, 110, 5, 13, 0, 0, 110, 111, 5, 1, 0,
		0, 111, 23, 1, 0, 0, 0, 11, 28, 34, 37, 50, 56, 66, 82, 88, 92, 95, 103,
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
	ParserFROM      = 4
	ParserIMPORT    = 5
	ParserCOMMENT   = 6
	ParserDOLLAR    = 7
	ParserLBRACE    = 8
	ParserRBRACE    = 9
	ParserLPAREN    = 10
	ParserRPAREN    = 11
	ParserCOMMA     = 12
	ParserCOLON     = 13
	ParserDOT       = 14
	ParserIDEN      = 15
	ParserID        = 16
	ParserINT       = 17
	ParserFLOAT     = 18
	ParserBOOLEAN   = 19
	ParserWS        = 20
)

// Parser rules.
const (
	ParserRULE_file            = 0
	ParserRULE_moduleDef       = 1
	ParserRULE_importDef       = 2
	ParserRULE_componentDef    = 3
	ParserRULE_variableDefList = 4
	ParserRULE_variableDef     = 5
	ParserRULE_variable        = 6
	ParserRULE_selectionSet    = 7
	ParserRULE_selection       = 8
	ParserRULE_tagDef          = 9
	ParserRULE_argumentList    = 10
	ParserRULE_argument        = 11
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetImports returns the imports rule contexts.
	GetImports() IImportDefContext

	// SetImports sets the imports rule contexts.
	SetImports(IImportDefContext)

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	imports IImportDefContext
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

func (s *FileContext) GetImports() IImportDefContext { return s.imports }

func (s *FileContext) SetImports(v IImportDefContext) { s.imports = v }

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

func (s *FileContext) AllImportDef() []IImportDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportDefContext); ok {
			len++
		}
	}

	tst := make([]IImportDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportDefContext); ok {
			tst[i] = t.(IImportDefContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) ImportDef(i int) IImportDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportDefContext); ok {
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

	return t.(IImportDefContext)
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

	p.SetState(37)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ParserMODULE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.ModuleDef()
		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ParserFROM {
			{
				p.SetState(25)

				var _x = p.ImportDef()

				localctx.(*FileContext).imports = _x
			}

			p.SetState(30)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ParserCOMPONENT {
			{
				p.SetState(31)
				p.ComponentDef()
			}

			p.SetState(34)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case ParserEOF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(36)
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
		p.SetState(39)
		p.Match(ParserMODULE)
	}
	{
		p.SetState(40)

		var _m = p.Match(ParserIDEN)

		localctx.(*ModuleDefContext).name = _m
	}

	return localctx
}

// IImportDefContext is an interface to support dynamic dispatch.
type IImportDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetModule returns the module token.
	GetModule() antlr.Token

	// Get_IDEN returns the _IDEN token.
	Get_IDEN() antlr.Token

	// SetModule sets the module token.
	SetModule(antlr.Token)

	// Set_IDEN sets the _IDEN token.
	Set_IDEN(antlr.Token)

	// GetImports returns the imports token list.
	GetImports() []antlr.Token

	// SetImports sets the imports token list.
	SetImports([]antlr.Token)

	// IsImportDefContext differentiates from other interfaces.
	IsImportDefContext()
}

type ImportDefContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	module  antlr.Token
	_IDEN   antlr.Token
	imports []antlr.Token
}

func NewEmptyImportDefContext() *ImportDefContext {
	var p = new(ImportDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParserRULE_importDef
	return p
}

func (*ImportDefContext) IsImportDefContext() {}

func NewImportDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDefContext {
	var p = new(ImportDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserRULE_importDef

	return p
}

func (s *ImportDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDefContext) GetModule() antlr.Token { return s.module }

func (s *ImportDefContext) Get_IDEN() antlr.Token { return s._IDEN }

func (s *ImportDefContext) SetModule(v antlr.Token) { s.module = v }

func (s *ImportDefContext) Set_IDEN(v antlr.Token) { s._IDEN = v }

func (s *ImportDefContext) GetImports() []antlr.Token { return s.imports }

func (s *ImportDefContext) SetImports(v []antlr.Token) { s.imports = v }

func (s *ImportDefContext) FROM() antlr.TerminalNode {
	return s.GetToken(ParserFROM, 0)
}

func (s *ImportDefContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(ParserIMPORT, 0)
}

func (s *ImportDefContext) AllIDEN() []antlr.TerminalNode {
	return s.GetTokens(ParserIDEN)
}

func (s *ImportDefContext) IDEN(i int) antlr.TerminalNode {
	return s.GetToken(ParserIDEN, i)
}

func (s *ImportDefContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ParserCOMMA)
}

func (s *ImportDefContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ParserCOMMA, i)
}

func (s *ImportDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterImportDef(s)
	}
}

func (s *ImportDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitImportDef(s)
	}
}

func (s *ImportDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ParserVisitor:
		return t.VisitImportDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Parser) ImportDef() (localctx IImportDefContext) {
	this := p
	_ = this

	localctx = NewImportDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParserRULE_importDef)
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
		p.SetState(42)
		p.Match(ParserFROM)
	}
	{
		p.SetState(43)

		var _m = p.Match(ParserIDEN)

		localctx.(*ImportDefContext).module = _m
	}
	{
		p.SetState(44)
		p.Match(ParserIMPORT)
	}
	{
		p.SetState(45)

		var _m = p.Match(ParserIDEN)

		localctx.(*ImportDefContext)._IDEN = _m
	}
	localctx.(*ImportDefContext).imports = append(localctx.(*ImportDefContext).imports, localctx.(*ImportDefContext)._IDEN)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ParserCOMMA {
		{
			p.SetState(46)
			p.Match(ParserCOMMA)
		}
		{
			p.SetState(47)

			var _m = p.Match(ParserIDEN)

			localctx.(*ImportDefContext)._IDEN = _m
		}
		localctx.(*ImportDefContext).imports = append(localctx.(*ImportDefContext).imports, localctx.(*ImportDefContext)._IDEN)

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 6, ParserRULE_componentDef)
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
		p.SetState(53)
		p.Match(ParserCOMPONENT)
	}
	{
		p.SetState(54)

		var _m = p.Match(ParserIDEN)

		localctx.(*ComponentDefContext).name = _m
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ParserLPAREN {
		{
			p.SetState(55)
			p.VariableDefList()
		}

	}
	{
		p.SetState(58)
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
	p.EnterRule(localctx, 8, ParserRULE_variableDefList)
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
		p.Match(ParserLPAREN)
	}
	{
		p.SetState(61)

		var _x = p.VariableDef()

		localctx.(*VariableDefListContext)._variableDef = _x
	}
	localctx.(*VariableDefListContext).variables = append(localctx.(*VariableDefListContext).variables, localctx.(*VariableDefListContext)._variableDef)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ParserCOMMA {
		{
			p.SetState(62)
			p.Match(ParserCOMMA)
		}
		{
			p.SetState(63)

			var _x = p.VariableDef()

			localctx.(*VariableDefListContext)._variableDef = _x
		}
		localctx.(*VariableDefListContext).variables = append(localctx.(*VariableDefListContext).variables, localctx.(*VariableDefListContext)._variableDef)

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(69)
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
	p.EnterRule(localctx, 10, ParserRULE_variableDef)

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
		p.Variable()
	}
	{
		p.SetState(72)
		p.Match(ParserCOLON)
	}
	{
		p.SetState(73)

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
	p.EnterRule(localctx, 12, ParserRULE_variable)

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
		p.SetState(75)
		p.Match(ParserDOLLAR)
	}
	{
		p.SetState(76)
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
	p.EnterRule(localctx, 14, ParserRULE_selectionSet)
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
		p.Match(ParserLBRACE)
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ParserSTRING || _la == ParserIDEN {
		{
			p.SetState(79)
			p.Selection()
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(84)
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
	p.EnterRule(localctx, 16, ParserRULE_selection)

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

	p.SetState(88)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ParserSTRING:
		localctx = NewSelectStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(ParserSTRING)
		}

	case ParserIDEN:
		localctx = NewSelectTagContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
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
	p.EnterRule(localctx, 18, ParserRULE_tagDef)
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
		p.SetState(90)
		p.Match(ParserIDEN)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ParserLPAREN {
		{
			p.SetState(91)
			p.ArgumentList()
		}

	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ParserLBRACE {
		{
			p.SetState(94)
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
	p.EnterRule(localctx, 20, ParserRULE_argumentList)
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
		p.SetState(97)
		p.Match(ParserLPAREN)
	}
	{
		p.SetState(98)

		var _x = p.Argument()

		localctx.(*ArgumentListContext)._argument = _x
	}
	localctx.(*ArgumentListContext).arguments = append(localctx.(*ArgumentListContext).arguments, localctx.(*ArgumentListContext)._argument)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ParserCOMMA {
		{
			p.SetState(99)
			p.Match(ParserCOMMA)
		}
		{
			p.SetState(100)

			var _x = p.Argument()

			localctx.(*ArgumentListContext)._argument = _x
		}
		localctx.(*ArgumentListContext).arguments = append(localctx.(*ArgumentListContext).arguments, localctx.(*ArgumentListContext)._argument)

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(106)
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
	p.EnterRule(localctx, 22, ParserRULE_argument)

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
		p.SetState(108)

		var _m = p.Match(ParserIDEN)

		localctx.(*ArgumentContext).key = _m
	}
	{
		p.SetState(109)
		p.Match(ParserCOLON)
	}
	{
		p.SetState(110)

		var _m = p.Match(ParserSTRING)

		localctx.(*ArgumentContext).value = _m
	}

	return localctx
}
