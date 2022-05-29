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
		"", "", "'module'", "'component'", "", "'$'", "'{'", "'}'", "'('", "')'",
		"','", "':'", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "STRING", "MODULE", "COMPONENT", "COMMENT", "DOLLAR", "LBRACE",
		"RBRACE", "LPAREN", "RPAREN", "COMMA", "COLON", "DOT", "IDEN", "ID",
		"INT", "FLOAT", "BOOLEAN", "WS",
	}
	staticData.ruleNames = []string{
		"MODULE", "COMPONENT", "DQUOTE", "COMMENT", "DOLLAR", "LBRACE", "RBRACE",
		"LPAREN", "RPAREN", "COMMA", "COLON", "DOT", "IDEN", "ID", "INT", "FLOAT",
		"BOOLEAN", "DIGIT", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 145, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 5, 2, 59, 8, 2, 10, 2, 12, 2, 62, 9, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 5, 3, 70, 8, 3, 10, 3, 12, 3, 73, 9, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 5, 12, 95, 8, 12, 10, 12, 12,
		12, 98, 9, 12, 1, 13, 1, 13, 1, 13, 5, 13, 103, 8, 13, 10, 13, 12, 13,
		106, 9, 13, 1, 13, 1, 13, 1, 14, 4, 14, 111, 8, 14, 11, 14, 12, 14, 112,
		1, 15, 4, 15, 116, 8, 15, 11, 15, 12, 15, 117, 1, 15, 1, 15, 4, 15, 122,
		8, 15, 11, 15, 12, 15, 123, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 3, 16, 135, 8, 16, 1, 17, 1, 17, 1, 18, 4, 18, 140, 8,
		18, 11, 18, 12, 18, 141, 1, 18, 1, 18, 1, 60, 0, 19, 1, 2, 3, 3, 5, 0,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 35, 0, 37, 18, 1, 0, 7, 2, 0, 10, 10, 13,
		13, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 5, 0,
		45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 44, 44, 46, 46, 1, 0, 48,
		57, 3, 0, 9, 10, 13, 13, 32, 32, 152, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 1, 39, 1, 0, 0, 0, 3, 46, 1, 0, 0, 0, 5, 56, 1, 0, 0, 0, 7,
		67, 1, 0, 0, 0, 9, 76, 1, 0, 0, 0, 11, 78, 1, 0, 0, 0, 13, 80, 1, 0, 0,
		0, 15, 82, 1, 0, 0, 0, 17, 84, 1, 0, 0, 0, 19, 86, 1, 0, 0, 0, 21, 88,
		1, 0, 0, 0, 23, 90, 1, 0, 0, 0, 25, 92, 1, 0, 0, 0, 27, 99, 1, 0, 0, 0,
		29, 110, 1, 0, 0, 0, 31, 115, 1, 0, 0, 0, 33, 134, 1, 0, 0, 0, 35, 136,
		1, 0, 0, 0, 37, 139, 1, 0, 0, 0, 39, 40, 5, 109, 0, 0, 40, 41, 5, 111,
		0, 0, 41, 42, 5, 100, 0, 0, 42, 43, 5, 117, 0, 0, 43, 44, 5, 108, 0, 0,
		44, 45, 5, 101, 0, 0, 45, 2, 1, 0, 0, 0, 46, 47, 5, 99, 0, 0, 47, 48, 5,
		111, 0, 0, 48, 49, 5, 109, 0, 0, 49, 50, 5, 112, 0, 0, 50, 51, 5, 111,
		0, 0, 51, 52, 5, 110, 0, 0, 52, 53, 5, 101, 0, 0, 53, 54, 5, 110, 0, 0,
		54, 55, 5, 116, 0, 0, 55, 4, 1, 0, 0, 0, 56, 60, 5, 34, 0, 0, 57, 59, 9,
		0, 0, 0, 58, 57, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 60,
		58, 1, 0, 0, 0, 61, 63, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 63, 64, 5, 34,
		0, 0, 64, 65, 1, 0, 0, 0, 65, 66, 6, 2, 0, 0, 66, 6, 1, 0, 0, 0, 67, 71,
		5, 35, 0, 0, 68, 70, 8, 0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0,
		71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 74, 1, 0, 0, 0, 73, 71, 1,
		0, 0, 0, 74, 75, 6, 3, 1, 0, 75, 8, 1, 0, 0, 0, 76, 77, 5, 36, 0, 0, 77,
		10, 1, 0, 0, 0, 78, 79, 5, 123, 0, 0, 79, 12, 1, 0, 0, 0, 80, 81, 5, 125,
		0, 0, 81, 14, 1, 0, 0, 0, 82, 83, 5, 40, 0, 0, 83, 16, 1, 0, 0, 0, 84,
		85, 5, 41, 0, 0, 85, 18, 1, 0, 0, 0, 86, 87, 5, 44, 0, 0, 87, 20, 1, 0,
		0, 0, 88, 89, 5, 58, 0, 0, 89, 22, 1, 0, 0, 0, 90, 91, 5, 46, 0, 0, 91,
		24, 1, 0, 0, 0, 92, 96, 7, 1, 0, 0, 93, 95, 7, 2, 0, 0, 94, 93, 1, 0, 0,
		0, 95, 98, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 26,
		1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99, 100, 5, 34, 0, 0, 100, 104, 7, 1, 0,
		0, 101, 103, 7, 3, 0, 0, 102, 101, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104,
		102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 107, 1, 0, 0, 0, 106, 104,
		1, 0, 0, 0, 107, 108, 5, 34, 0, 0, 108, 28, 1, 0, 0, 0, 109, 111, 3, 35,
		17, 0, 110, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0,
		112, 113, 1, 0, 0, 0, 113, 30, 1, 0, 0, 0, 114, 116, 3, 35, 17, 0, 115,
		114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118,
		1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 121, 7, 4, 0, 0, 120, 122, 3, 35,
		17, 0, 121, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0,
		123, 124, 1, 0, 0, 0, 124, 32, 1, 0, 0, 0, 125, 126, 5, 116, 0, 0, 126,
		127, 5, 114, 0, 0, 127, 128, 5, 117, 0, 0, 128, 135, 5, 101, 0, 0, 129,
		130, 5, 102, 0, 0, 130, 131, 5, 97, 0, 0, 131, 132, 5, 108, 0, 0, 132,
		133, 5, 115, 0, 0, 133, 135, 5, 101, 0, 0, 134, 125, 1, 0, 0, 0, 134, 129,
		1, 0, 0, 0, 135, 34, 1, 0, 0, 0, 136, 137, 7, 5, 0, 0, 137, 36, 1, 0, 0,
		0, 138, 140, 7, 6, 0, 0, 139, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141,
		139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144,
		6, 18, 1, 0, 144, 38, 1, 0, 0, 0, 10, 0, 60, 71, 96, 104, 112, 117, 123,
		134, 141, 2, 7, 1, 0, 0, 1, 0,
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
	LexerMODULE    = 2
	LexerCOMPONENT = 3
	LexerCOMMENT   = 4
	LexerDOLLAR    = 5
	LexerLBRACE    = 6
	LexerRBRACE    = 7
	LexerLPAREN    = 8
	LexerRPAREN    = 9
	LexerCOMMA     = 10
	LexerCOLON     = 11
	LexerDOT       = 12
	LexerIDEN      = 13
	LexerID        = 14
	LexerINT       = 15
	LexerFLOAT     = 16
	LexerBOOLEAN   = 17
	LexerWS        = 18
)
