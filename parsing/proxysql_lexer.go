// Code generated from ProxySQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ProxySQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ProxySQLLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func proxysqllexerLexerInit() {
	staticData := &ProxySQLLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "':'", "'='", "';'", "','", "'['", "']'", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "BOOLEAN", "ID", "STRING",
		"NUMBER", "FLOAT", "HEX", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "BOOLEAN", "ID", "STRING", "NUMBER", "FLOAT", "HEX", "COMMENT",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 147, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10,
		76, 8, 10, 1, 11, 1, 11, 5, 11, 80, 8, 11, 10, 11, 12, 11, 83, 9, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 5, 12, 89, 8, 12, 10, 12, 12, 12, 92, 9, 12, 1,
		12, 1, 12, 1, 13, 4, 13, 97, 8, 13, 11, 13, 12, 13, 98, 1, 14, 4, 14, 102,
		8, 14, 11, 14, 12, 14, 103, 1, 14, 1, 14, 5, 14, 108, 8, 14, 10, 14, 12,
		14, 111, 9, 14, 1, 14, 1, 14, 3, 14, 115, 8, 14, 1, 14, 4, 14, 118, 8,
		14, 11, 14, 12, 14, 119, 3, 14, 122, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		4, 15, 128, 8, 15, 11, 15, 12, 15, 129, 1, 16, 1, 16, 5, 16, 134, 8, 16,
		10, 16, 12, 16, 137, 9, 16, 1, 16, 1, 16, 1, 17, 4, 17, 142, 8, 17, 11,
		17, 12, 17, 143, 1, 17, 1, 17, 0, 0, 18, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 33, 17, 35, 18, 1, 0, 9, 3, 0, 65, 90, 95, 95, 97, 122, 4,
		0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 34, 34, 92, 92, 1, 0, 48, 57,
		2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 3, 0, 48, 57, 65, 70, 97,
		102, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 161, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 1, 37, 1, 0, 0, 0, 3, 39, 1, 0, 0, 0,
		5, 41, 1, 0, 0, 0, 7, 43, 1, 0, 0, 0, 9, 45, 1, 0, 0, 0, 11, 47, 1, 0,
		0, 0, 13, 49, 1, 0, 0, 0, 15, 51, 1, 0, 0, 0, 17, 53, 1, 0, 0, 0, 19, 55,
		1, 0, 0, 0, 21, 75, 1, 0, 0, 0, 23, 77, 1, 0, 0, 0, 25, 84, 1, 0, 0, 0,
		27, 96, 1, 0, 0, 0, 29, 101, 1, 0, 0, 0, 31, 123, 1, 0, 0, 0, 33, 131,
		1, 0, 0, 0, 35, 141, 1, 0, 0, 0, 37, 38, 5, 58, 0, 0, 38, 2, 1, 0, 0, 0,
		39, 40, 5, 61, 0, 0, 40, 4, 1, 0, 0, 0, 41, 42, 5, 59, 0, 0, 42, 6, 1,
		0, 0, 0, 43, 44, 5, 44, 0, 0, 44, 8, 1, 0, 0, 0, 45, 46, 5, 91, 0, 0, 46,
		10, 1, 0, 0, 0, 47, 48, 5, 93, 0, 0, 48, 12, 1, 0, 0, 0, 49, 50, 5, 40,
		0, 0, 50, 14, 1, 0, 0, 0, 51, 52, 5, 41, 0, 0, 52, 16, 1, 0, 0, 0, 53,
		54, 5, 123, 0, 0, 54, 18, 1, 0, 0, 0, 55, 56, 5, 125, 0, 0, 56, 20, 1,
		0, 0, 0, 57, 58, 5, 116, 0, 0, 58, 59, 5, 114, 0, 0, 59, 60, 5, 117, 0,
		0, 60, 76, 5, 101, 0, 0, 61, 62, 5, 102, 0, 0, 62, 63, 5, 97, 0, 0, 63,
		64, 5, 108, 0, 0, 64, 65, 5, 115, 0, 0, 65, 76, 5, 101, 0, 0, 66, 67, 5,
		84, 0, 0, 67, 68, 5, 82, 0, 0, 68, 69, 5, 85, 0, 0, 69, 76, 5, 69, 0, 0,
		70, 71, 5, 70, 0, 0, 71, 72, 5, 65, 0, 0, 72, 73, 5, 76, 0, 0, 73, 74,
		5, 83, 0, 0, 74, 76, 5, 69, 0, 0, 75, 57, 1, 0, 0, 0, 75, 61, 1, 0, 0,
		0, 75, 66, 1, 0, 0, 0, 75, 70, 1, 0, 0, 0, 76, 22, 1, 0, 0, 0, 77, 81,
		7, 0, 0, 0, 78, 80, 7, 1, 0, 0, 79, 78, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0,
		81, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 24, 1, 0, 0, 0, 83, 81, 1,
		0, 0, 0, 84, 90, 5, 34, 0, 0, 85, 89, 8, 2, 0, 0, 86, 87, 5, 92, 0, 0,
		87, 89, 9, 0, 0, 0, 88, 85, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 92, 1,
		0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 93, 1, 0, 0, 0, 92,
		90, 1, 0, 0, 0, 93, 94, 5, 34, 0, 0, 94, 26, 1, 0, 0, 0, 95, 97, 7, 3,
		0, 0, 96, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99,
		1, 0, 0, 0, 99, 28, 1, 0, 0, 0, 100, 102, 7, 3, 0, 0, 101, 100, 1, 0, 0,
		0, 102, 103, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104,
		105, 1, 0, 0, 0, 105, 109, 5, 46, 0, 0, 106, 108, 7, 3, 0, 0, 107, 106,
		1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0,
		0, 0, 110, 121, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 114, 7, 4, 0, 0,
		113, 115, 7, 5, 0, 0, 114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115,
		117, 1, 0, 0, 0, 116, 118, 7, 3, 0, 0, 117, 116, 1, 0, 0, 0, 118, 119,
		1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 122, 1, 0,
		0, 0, 121, 112, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 30, 1, 0, 0, 0,
		123, 124, 5, 48, 0, 0, 124, 125, 5, 120, 0, 0, 125, 127, 1, 0, 0, 0, 126,
		128, 7, 6, 0, 0, 127, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 127,
		1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 32, 1, 0, 0, 0, 131, 135, 5, 35,
		0, 0, 132, 134, 8, 7, 0, 0, 133, 132, 1, 0, 0, 0, 134, 137, 1, 0, 0, 0,
		135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 138, 1, 0, 0, 0, 137,
		135, 1, 0, 0, 0, 138, 139, 6, 16, 0, 0, 139, 34, 1, 0, 0, 0, 140, 142,
		7, 8, 0, 0, 141, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 141, 1, 0,
		0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 6, 17, 0, 0,
		146, 36, 1, 0, 0, 0, 14, 0, 75, 81, 88, 90, 98, 103, 109, 114, 119, 121,
		129, 135, 143, 1, 6, 0, 0,
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

// ProxySQLLexerInit initializes any static state used to implement ProxySQLLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewProxySQLLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ProxySQLLexerInit() {
	staticData := &ProxySQLLexerLexerStaticData
	staticData.once.Do(proxysqllexerLexerInit)
}

// NewProxySQLLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewProxySQLLexer(input antlr.CharStream) *ProxySQLLexer {
	ProxySQLLexerInit()
	l := new(ProxySQLLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ProxySQLLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ProxySQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ProxySQLLexer tokens.
const (
	ProxySQLLexerT__0    = 1
	ProxySQLLexerT__1    = 2
	ProxySQLLexerT__2    = 3
	ProxySQLLexerT__3    = 4
	ProxySQLLexerT__4    = 5
	ProxySQLLexerT__5    = 6
	ProxySQLLexerT__6    = 7
	ProxySQLLexerT__7    = 8
	ProxySQLLexerT__8    = 9
	ProxySQLLexerT__9    = 10
	ProxySQLLexerBOOLEAN = 11
	ProxySQLLexerID      = 12
	ProxySQLLexerSTRING  = 13
	ProxySQLLexerNUMBER  = 14
	ProxySQLLexerFLOAT   = 15
	ProxySQLLexerHEX     = 16
	ProxySQLLexerCOMMENT = 17
	ProxySQLLexerWS      = 18
)
