// Code generated from krpc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // krpc
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

type krpcParser struct {
	*antlr.BaseParser
}

var KrpcParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func krpcParserInit() {
	staticData := &KrpcParserStaticData
	staticData.RuleNames = []string{
		"krpc",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 0, 5, 2, 0, 7, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 3, 0, 2, 1,
		0, 0, 0, 2, 3, 5, 0, 0, 1, 3, 1, 1, 0, 0, 0, 0,
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

// krpcParserInit initializes any static state used to implement krpcParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewkrpcParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func KrpcParserInit() {
	staticData := &KrpcParserStaticData
	staticData.once.Do(krpcParserInit)
}

// NewkrpcParser produces a new parser instance for the optional input antlr.TokenStream.
func NewkrpcParser(input antlr.TokenStream) *krpcParser {
	KrpcParserInit()
	this := new(krpcParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &KrpcParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "krpc.g4"

	return this
}

// krpcParserEOF is the krpcParser token.
const krpcParserEOF = antlr.TokenEOF

// krpcParserRULE_krpc is the krpcParser rule.
const krpcParserRULE_krpc = 0

// IKrpcContext is an interface to support dynamic dispatch.
type IKrpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode

	// IsKrpcContext differentiates from other interfaces.
	IsKrpcContext()
}

type KrpcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKrpcContext() *KrpcContext {
	var p = new(KrpcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = krpcParserRULE_krpc
	return p
}

func InitEmptyKrpcContext(p *KrpcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = krpcParserRULE_krpc
}

func (*KrpcContext) IsKrpcContext() {}

func NewKrpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KrpcContext {
	var p = new(KrpcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = krpcParserRULE_krpc

	return p
}

func (s *KrpcContext) GetParser() antlr.Parser { return s.parser }

func (s *KrpcContext) EOF() antlr.TerminalNode {
	return s.GetToken(krpcParserEOF, 0)
}

func (s *KrpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KrpcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KrpcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(krpcListener); ok {
		listenerT.EnterKrpc(s)
	}
}

func (s *KrpcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(krpcListener); ok {
		listenerT.ExitKrpc(s)
	}
}

func (p *krpcParser) Krpc() (localctx IKrpcContext) {
	localctx = NewKrpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, krpcParserRULE_krpc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(2)
		p.Match(krpcParserEOF)
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
