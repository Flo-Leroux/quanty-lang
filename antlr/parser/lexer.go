// Code generated from ./Lexer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lexerLexerInit() {
	staticData := &lexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"COMPONENT", "DQUOTE", "COMMENT", "DOLLAR", "LBRACE", "RBRACE", "LPAREN",
		"RPAREN", "COMMA", "COLON", "DOT", "IDEN", "ID", "INT", "FLOAT", "BOOLEAN",
		"DIGIT", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 136, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 50, 8, 1, 10, 1, 12, 1, 53, 9,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 5, 2, 61, 8, 2, 10, 2, 12, 2, 64,
		9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 5, 11, 86, 8,
		11, 10, 11, 12, 11, 89, 9, 11, 1, 12, 1, 12, 1, 12, 5, 12, 94, 8, 12, 10,
		12, 12, 12, 97, 9, 12, 1, 12, 1, 12, 1, 13, 4, 13, 102, 8, 13, 11, 13,
		12, 13, 103, 1, 14, 4, 14, 107, 8, 14, 11, 14, 12, 14, 108, 1, 14, 1, 14,
		4, 14, 113, 8, 14, 11, 14, 12, 14, 114, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 126, 8, 15, 1, 16, 1, 16, 1, 17,
		4, 17, 131, 8, 17, 11, 17, 12, 17, 132, 1, 17, 1, 17, 1, 51, 0, 18, 1,
		2, 3, 0, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 0, 35, 17, 1, 0, 7, 2, 0, 10,
		10, 13, 13, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 44, 44, 46, 46, 1,
		0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 143, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		1, 37, 1, 0, 0, 0, 3, 47, 1, 0, 0, 0, 5, 58, 1, 0, 0, 0, 7, 67, 1, 0, 0,
		0, 9, 69, 1, 0, 0, 0, 11, 71, 1, 0, 0, 0, 13, 73, 1, 0, 0, 0, 15, 75, 1,
		0, 0, 0, 17, 77, 1, 0, 0, 0, 19, 79, 1, 0, 0, 0, 21, 81, 1, 0, 0, 0, 23,
		83, 1, 0, 0, 0, 25, 90, 1, 0, 0, 0, 27, 101, 1, 0, 0, 0, 29, 106, 1, 0,
		0, 0, 31, 125, 1, 0, 0, 0, 33, 127, 1, 0, 0, 0, 35, 130, 1, 0, 0, 0, 37,
		38, 5, 99, 0, 0, 38, 39, 5, 111, 0, 0, 39, 40, 5, 109, 0, 0, 40, 41, 5,
		112, 0, 0, 41, 42, 5, 111, 0, 0, 42, 43, 5, 110, 0, 0, 43, 44, 5, 101,
		0, 0, 44, 45, 5, 110, 0, 0, 45, 46, 5, 116, 0, 0, 46, 2, 1, 0, 0, 0, 47,
		51, 5, 34, 0, 0, 48, 50, 9, 0, 0, 0, 49, 48, 1, 0, 0, 0, 50, 53, 1, 0,
		0, 0, 51, 52, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 54, 1, 0, 0, 0, 53, 51,
		1, 0, 0, 0, 54, 55, 5, 34, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 6, 1, 0, 0,
		57, 4, 1, 0, 0, 0, 58, 62, 5, 35, 0, 0, 59, 61, 8, 0, 0, 0, 60, 59, 1,
		0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63,
		65, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 66, 6, 2, 1, 0, 66, 6, 1, 0, 0,
		0, 67, 68, 5, 36, 0, 0, 68, 8, 1, 0, 0, 0, 69, 70, 5, 123, 0, 0, 70, 10,
		1, 0, 0, 0, 71, 72, 5, 125, 0, 0, 72, 12, 1, 0, 0, 0, 73, 74, 5, 40, 0,
		0, 74, 14, 1, 0, 0, 0, 75, 76, 5, 41, 0, 0, 76, 16, 1, 0, 0, 0, 77, 78,
		5, 44, 0, 0, 78, 18, 1, 0, 0, 0, 79, 80, 5, 58, 0, 0, 80, 20, 1, 0, 0,
		0, 81, 82, 5, 46, 0, 0, 82, 22, 1, 0, 0, 0, 83, 87, 7, 1, 0, 0, 84, 86,
		7, 2, 0, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0,
		87, 88, 1, 0, 0, 0, 88, 24, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 91, 5,
		34, 0, 0, 91, 95, 7, 1, 0, 0, 92, 94, 7, 3, 0, 0, 93, 92, 1, 0, 0, 0, 94,
		97, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 98, 1, 0, 0,
		0, 97, 95, 1, 0, 0, 0, 98, 99, 5, 34, 0, 0, 99, 26, 1, 0, 0, 0, 100, 102,
		3, 33, 16, 0, 101, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 101, 1,
		0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 28, 1, 0, 0, 0, 105, 107, 3, 33, 16,
		0, 106, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108,
		109, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 112, 7, 4, 0, 0, 111, 113,
		3, 33, 16, 0, 112, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 112, 1,
		0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 30, 1, 0, 0, 0, 116, 117, 5, 116, 0,
		0, 117, 118, 5, 114, 0, 0, 118, 119, 5, 117, 0, 0, 119, 126, 5, 101, 0,
		0, 120, 121, 5, 102, 0, 0, 121, 122, 5, 97, 0, 0, 122, 123, 5, 108, 0,
		0, 123, 124, 5, 115, 0, 0, 124, 126, 5, 101, 0, 0, 125, 116, 1, 0, 0, 0,
		125, 120, 1, 0, 0, 0, 126, 32, 1, 0, 0, 0, 127, 128, 7, 5, 0, 0, 128, 34,
		1, 0, 0, 0, 129, 131, 7, 6, 0, 0, 130, 129, 1, 0, 0, 0, 131, 132, 1, 0,
		0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0,
		134, 135, 6, 17, 1, 0, 135, 36, 1, 0, 0, 0, 10, 0, 51, 62, 87, 95, 103,
		108, 114, 125, 132, 2, 7, 1, 0, 0, 1, 0,
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

// LexerInit initializes any static state used to implement Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LexerInit() {
	staticData := &lexerLexerStaticData
	staticData.once.Do(lexerLexerInit)
}

// NewLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLexer(input antlr.CharStream) *Lexer {
	LexerInit()
	l := new(Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &lexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Lexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Lexer tokens.
const (
	LexerSTRING    = 1
	LexerCOMPONENT = 2
	LexerCOMMENT   = 3
	LexerDOLLAR    = 4
	LexerLBRACE    = 5
	LexerRBRACE    = 6
	LexerLPAREN    = 7
	LexerRPAREN    = 8
	LexerCOMMA     = 9
	LexerCOLON     = 10
	LexerDOT       = 11
	LexerIDEN      = 12
	LexerID        = 13
	LexerINT       = 14
	LexerFLOAT     = 15
	LexerBOOLEAN   = 16
	LexerWS        = 17
)
