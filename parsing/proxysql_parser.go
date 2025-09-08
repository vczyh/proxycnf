// Code generated from ProxySQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // ProxySQL
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

type ProxySQLParser struct {
	*antlr.BaseParser
}

var ProxySQLParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func proxysqlParserInit() {
	staticData := &ProxySQLParserStaticData
	staticData.LiteralNames = []string{
		"", "':'", "'='", "';'", "','", "'['", "']'", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "BOOLEAN", "ID", "STRING",
		"NUMBER", "FLOAT", "HEX", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"configuration", "settingList", "setting", "name", "value", "valueList",
		"scalarValue", "scalarValueList", "array", "list", "group",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 99, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 3, 0, 24, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5,
		1, 33, 8, 1, 10, 1, 12, 1, 36, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 44, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 52, 8, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 60, 8, 5, 1, 5, 5, 5, 63, 8, 5,
		10, 5, 12, 5, 66, 9, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		5, 7, 76, 8, 7, 10, 7, 12, 7, 79, 9, 7, 1, 8, 1, 8, 3, 8, 83, 8, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 3, 9, 89, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 3, 10, 95,
		8, 10, 1, 10, 1, 10, 1, 10, 0, 3, 2, 10, 14, 11, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 0, 2, 1, 0, 1, 2, 2, 0, 11, 11, 13, 16, 100, 0, 23, 1,
		0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 37, 1, 0, 0, 0, 6, 45, 1, 0, 0, 0, 8, 51,
		1, 0, 0, 0, 10, 53, 1, 0, 0, 0, 12, 67, 1, 0, 0, 0, 14, 69, 1, 0, 0, 0,
		16, 80, 1, 0, 0, 0, 18, 86, 1, 0, 0, 0, 20, 92, 1, 0, 0, 0, 22, 24, 3,
		2, 1, 0, 23, 22, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25,
		26, 5, 0, 0, 1, 26, 1, 1, 0, 0, 0, 27, 28, 6, 1, -1, 0, 28, 29, 3, 4, 2,
		0, 29, 34, 1, 0, 0, 0, 30, 31, 10, 1, 0, 0, 31, 33, 3, 4, 2, 0, 32, 30,
		1, 0, 0, 0, 33, 36, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0,
		35, 3, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 37, 38, 3, 6, 3, 0, 38, 39, 7, 0,
		0, 0, 39, 43, 3, 8, 4, 0, 40, 44, 5, 3, 0, 0, 41, 44, 5, 4, 0, 0, 42, 44,
		1, 0, 0, 0, 43, 40, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 42, 1, 0, 0, 0,
		44, 5, 1, 0, 0, 0, 45, 46, 5, 12, 0, 0, 46, 7, 1, 0, 0, 0, 47, 52, 3, 12,
		6, 0, 48, 52, 3, 16, 8, 0, 49, 52, 3, 18, 9, 0, 50, 52, 3, 20, 10, 0, 51,
		47, 1, 0, 0, 0, 51, 48, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 50, 1, 0, 0,
		0, 52, 9, 1, 0, 0, 0, 53, 54, 6, 5, -1, 0, 54, 55, 3, 8, 4, 0, 55, 64,
		1, 0, 0, 0, 56, 59, 10, 1, 0, 0, 57, 60, 5, 4, 0, 0, 58, 60, 1, 0, 0, 0,
		59, 57, 1, 0, 0, 0, 59, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 63, 3,
		8, 4, 0, 62, 56, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64,
		65, 1, 0, 0, 0, 65, 11, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 68, 7, 1, 0,
		0, 68, 13, 1, 0, 0, 0, 69, 70, 6, 7, -1, 0, 70, 71, 3, 12, 6, 0, 71, 77,
		1, 0, 0, 0, 72, 73, 10, 1, 0, 0, 73, 74, 5, 4, 0, 0, 74, 76, 3, 12, 6,
		0, 75, 72, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78,
		1, 0, 0, 0, 78, 15, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 82, 5, 5, 0, 0,
		81, 83, 3, 14, 7, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 1,
		0, 0, 0, 84, 85, 5, 6, 0, 0, 85, 17, 1, 0, 0, 0, 86, 88, 5, 7, 0, 0, 87,
		89, 3, 10, 5, 0, 88, 87, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 90, 1, 0,
		0, 0, 90, 91, 5, 8, 0, 0, 91, 19, 1, 0, 0, 0, 92, 94, 5, 9, 0, 0, 93, 95,
		3, 2, 1, 0, 94, 93, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0,
		96, 97, 5, 10, 0, 0, 97, 21, 1, 0, 0, 0, 10, 23, 34, 43, 51, 59, 64, 77,
		82, 88, 94,
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

// ProxySQLParserInit initializes any static state used to implement ProxySQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewProxySQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ProxySQLParserInit() {
	staticData := &ProxySQLParserStaticData
	staticData.once.Do(proxysqlParserInit)
}

// NewProxySQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewProxySQLParser(input antlr.TokenStream) *ProxySQLParser {
	ProxySQLParserInit()
	this := new(ProxySQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ProxySQLParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ProxySQL.g4"

	return this
}

// ProxySQLParser tokens.
const (
	ProxySQLParserEOF     = antlr.TokenEOF
	ProxySQLParserT__0    = 1
	ProxySQLParserT__1    = 2
	ProxySQLParserT__2    = 3
	ProxySQLParserT__3    = 4
	ProxySQLParserT__4    = 5
	ProxySQLParserT__5    = 6
	ProxySQLParserT__6    = 7
	ProxySQLParserT__7    = 8
	ProxySQLParserT__8    = 9
	ProxySQLParserT__9    = 10
	ProxySQLParserBOOLEAN = 11
	ProxySQLParserID      = 12
	ProxySQLParserSTRING  = 13
	ProxySQLParserNUMBER  = 14
	ProxySQLParserFLOAT   = 15
	ProxySQLParserHEX     = 16
	ProxySQLParserCOMMENT = 17
	ProxySQLParserWS      = 18
)

// ProxySQLParser rules.
const (
	ProxySQLParserRULE_configuration   = 0
	ProxySQLParserRULE_settingList     = 1
	ProxySQLParserRULE_setting         = 2
	ProxySQLParserRULE_name            = 3
	ProxySQLParserRULE_value           = 4
	ProxySQLParserRULE_valueList       = 5
	ProxySQLParserRULE_scalarValue     = 6
	ProxySQLParserRULE_scalarValueList = 7
	ProxySQLParserRULE_array           = 8
	ProxySQLParserRULE_list            = 9
	ProxySQLParserRULE_group           = 10
)

// IConfigurationContext is an interface to support dynamic dispatch.
type IConfigurationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	SettingList() ISettingListContext

	// IsConfigurationContext differentiates from other interfaces.
	IsConfigurationContext()
}

type ConfigurationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigurationContext() *ConfigurationContext {
	var p = new(ConfigurationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_configuration
	return p
}

func InitEmptyConfigurationContext(p *ConfigurationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_configuration
}

func (*ConfigurationContext) IsConfigurationContext() {}

func NewConfigurationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigurationContext {
	var p = new(ConfigurationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProxySQLParserRULE_configuration

	return p
}

func (s *ConfigurationContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigurationContext) EOF() antlr.TerminalNode {
	return s.GetToken(ProxySQLParserEOF, 0)
}

func (s *ConfigurationContext) SettingList() ISettingListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISettingListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISettingListContext)
}

func (s *ConfigurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigurationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigurationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.EnterConfiguration(s)
	}
}

func (s *ConfigurationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.ExitConfiguration(s)
	}
}

func (p *ProxySQLParser) Configuration() (localctx IConfigurationContext) {
	localctx = NewConfigurationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ProxySQLParserRULE_configuration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ProxySQLParserID {
		{
			p.SetState(22)
			p.settingList(0)
		}

	}
	{
		p.SetState(25)
		p.Match(ProxySQLParserEOF)
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

// ISettingListContext is an interface to support dynamic dispatch.
type ISettingListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Setting() ISettingContext
	SettingList() ISettingListContext

	// IsSettingListContext differentiates from other interfaces.
	IsSettingListContext()
}

type SettingListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySettingListContext() *SettingListContext {
	var p = new(SettingListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_settingList
	return p
}

func InitEmptySettingListContext(p *SettingListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_settingList
}

func (*SettingListContext) IsSettingListContext() {}

func NewSettingListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SettingListContext {
	var p = new(SettingListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProxySQLParserRULE_settingList

	return p
}

func (s *SettingListContext) GetParser() antlr.Parser { return s.parser }

func (s *SettingListContext) Setting() ISettingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISettingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISettingContext)
}

func (s *SettingListContext) SettingList() ISettingListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISettingListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISettingListContext)
}

func (s *SettingListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SettingListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SettingListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.EnterSettingList(s)
	}
}

func (s *SettingListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.ExitSettingList(s)
	}
}

func (p *ProxySQLParser) SettingList() (localctx ISettingListContext) {
	return p.settingList(0)
}

func (p *ProxySQLParser) settingList(_p int) (localctx ISettingListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewSettingListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISettingListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ProxySQLParserRULE_settingList, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Setting()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSettingListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ProxySQLParserRULE_settingList)
			p.SetState(30)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(31)
				p.Setting()
			}

		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISettingContext is an interface to support dynamic dispatch.
type ISettingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Value() IValueContext

	// IsSettingContext differentiates from other interfaces.
	IsSettingContext()
}

type SettingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySettingContext() *SettingContext {
	var p = new(SettingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_setting
	return p
}

func InitEmptySettingContext(p *SettingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_setting
}

func (*SettingContext) IsSettingContext() {}

func NewSettingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SettingContext {
	var p = new(SettingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProxySQLParserRULE_setting

	return p
}

func (s *SettingContext) GetParser() antlr.Parser { return s.parser }

func (s *SettingContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *SettingContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SettingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SettingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SettingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.EnterSetting(s)
	}
}

func (s *SettingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.ExitSetting(s)
	}
}

func (p *ProxySQLParser) Setting() (localctx ISettingContext) {
	localctx = NewSettingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ProxySQLParserRULE_setting)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Name()
	}
	{
		p.SetState(38)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ProxySQLParserT__0 || _la == ProxySQLParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(39)
		p.Value()
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(40)
			p.Match(ProxySQLParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(41)
			p.Match(ProxySQLParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:

	case antlr.ATNInvalidAltNumber:
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

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProxySQLParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) ID() antlr.TerminalNode {
	return s.GetToken(ProxySQLParserID, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *ProxySQLParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ProxySQLParserRULE_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(ProxySQLParserID)
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ScalarValue() IScalarValueContext
	Array() IArrayContext
	List() IListContext
	Group() IGroupContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProxySQLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) ScalarValue() IScalarValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarValueContext)
}

func (s *ValueContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ValueContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ValueContext) Group() IGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *ProxySQLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ProxySQLParserRULE_value)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ProxySQLParserBOOLEAN, ProxySQLParserSTRING, ProxySQLParserNUMBER, ProxySQLParserFLOAT, ProxySQLParserHEX:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.ScalarValue()
		}

	case ProxySQLParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.Array()
		}

	case ProxySQLParserT__6:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(49)
			p.List()
		}

	case ProxySQLParserT__8:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(50)
			p.Group()
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

// IValueListContext is an interface to support dynamic dispatch.
type IValueListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	ValueList() IValueListContext

	// IsValueListContext differentiates from other interfaces.
	IsValueListContext()
}

type ValueListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueListContext() *ValueListContext {
	var p = new(ValueListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_valueList
	return p
}

func InitEmptyValueListContext(p *ValueListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_valueList
}

func (*ValueListContext) IsValueListContext() {}

func NewValueListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueListContext {
	var p = new(ValueListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProxySQLParserRULE_valueList

	return p
}

func (s *ValueListContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueListContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValueListContext) ValueList() IValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *ValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.EnterValueList(s)
	}
}

func (s *ValueListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.ExitValueList(s)
	}
}

func (p *ProxySQLParser) ValueList() (localctx IValueListContext) {
	return p.valueList(0)
}

func (p *ProxySQLParser) valueList(_p int) (localctx IValueListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewValueListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IValueListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, ProxySQLParserRULE_valueList, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Value()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewValueListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ProxySQLParserRULE_valueList)
			p.SetState(56)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(59)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case ProxySQLParserT__3:
				{
					p.SetState(57)
					p.Match(ProxySQLParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case ProxySQLParserT__4, ProxySQLParserT__6, ProxySQLParserT__8, ProxySQLParserBOOLEAN, ProxySQLParserSTRING, ProxySQLParserNUMBER, ProxySQLParserFLOAT, ProxySQLParserHEX:

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}
			{
				p.SetState(61)
				p.Value()
			}

		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScalarValueContext is an interface to support dynamic dispatch.
type IScalarValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOLEAN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	HEX() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsScalarValueContext differentiates from other interfaces.
	IsScalarValueContext()
}

type ScalarValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarValueContext() *ScalarValueContext {
	var p = new(ScalarValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_scalarValue
	return p
}

func InitEmptyScalarValueContext(p *ScalarValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_scalarValue
}

func (*ScalarValueContext) IsScalarValueContext() {}

func NewScalarValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarValueContext {
	var p = new(ScalarValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProxySQLParserRULE_scalarValue

	return p
}

func (s *ScalarValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarValueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ProxySQLParserBOOLEAN, 0)
}

func (s *ScalarValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ProxySQLParserNUMBER, 0)
}

func (s *ScalarValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ProxySQLParserFLOAT, 0)
}

func (s *ScalarValueContext) HEX() antlr.TerminalNode {
	return s.GetToken(ProxySQLParserHEX, 0)
}

func (s *ScalarValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(ProxySQLParserSTRING, 0)
}

func (s *ScalarValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.EnterScalarValue(s)
	}
}

func (s *ScalarValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.ExitScalarValue(s)
	}
}

func (p *ProxySQLParser) ScalarValue() (localctx IScalarValueContext) {
	localctx = NewScalarValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ProxySQLParserRULE_scalarValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&124928) != 0) {
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

// IScalarValueListContext is an interface to support dynamic dispatch.
type IScalarValueListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ScalarValue() IScalarValueContext
	ScalarValueList() IScalarValueListContext

	// IsScalarValueListContext differentiates from other interfaces.
	IsScalarValueListContext()
}

type ScalarValueListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarValueListContext() *ScalarValueListContext {
	var p = new(ScalarValueListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_scalarValueList
	return p
}

func InitEmptyScalarValueListContext(p *ScalarValueListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_scalarValueList
}

func (*ScalarValueListContext) IsScalarValueListContext() {}

func NewScalarValueListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarValueListContext {
	var p = new(ScalarValueListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProxySQLParserRULE_scalarValueList

	return p
}

func (s *ScalarValueListContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarValueListContext) ScalarValue() IScalarValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarValueContext)
}

func (s *ScalarValueListContext) ScalarValueList() IScalarValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarValueListContext)
}

func (s *ScalarValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarValueListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarValueListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.EnterScalarValueList(s)
	}
}

func (s *ScalarValueListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.ExitScalarValueList(s)
	}
}

func (p *ProxySQLParser) ScalarValueList() (localctx IScalarValueListContext) {
	return p.scalarValueList(0)
}

func (p *ProxySQLParser) scalarValueList(_p int) (localctx IScalarValueListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewScalarValueListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IScalarValueListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, ProxySQLParserRULE_scalarValueList, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.ScalarValue()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewScalarValueListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ProxySQLParserRULE_scalarValueList)
			p.SetState(72)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(73)
				p.Match(ProxySQLParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(74)
				p.ScalarValue()
			}

		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ScalarValueList() IScalarValueListContext

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_array
	return p
}

func InitEmptyArrayContext(p *ArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_array
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProxySQLParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) ScalarValueList() IScalarValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarValueListContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.ExitArray(s)
	}
}

func (p *ProxySQLParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ProxySQLParserRULE_array)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(ProxySQLParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&124928) != 0 {
		{
			p.SetState(81)
			p.scalarValueList(0)
		}

	}
	{
		p.SetState(84)
		p.Match(ProxySQLParserT__5)
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

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ValueList() IValueListContext

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProxySQLParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) ValueList() IValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.ExitList(s)
	}
}

func (p *ProxySQLParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ProxySQLParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(ProxySQLParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&125600) != 0 {
		{
			p.SetState(87)
			p.valueList(0)
		}

	}
	{
		p.SetState(90)
		p.Match(ProxySQLParserT__7)
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

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SettingList() ISettingListContext

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_group
	return p
}

func InitEmptyGroupContext(p *GroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProxySQLParserRULE_group
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProxySQLParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) SettingList() ISettingListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISettingListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISettingListContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProxySQLListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (p *ProxySQLParser) Group() (localctx IGroupContext) {
	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ProxySQLParserRULE_group)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(ProxySQLParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ProxySQLParserID {
		{
			p.SetState(93)
			p.settingList(0)
		}

	}
	{
		p.SetState(96)
		p.Match(ProxySQLParserT__9)
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

func (p *ProxySQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *SettingListContext = nil
		if localctx != nil {
			t = localctx.(*SettingListContext)
		}
		return p.SettingList_Sempred(t, predIndex)

	case 5:
		var t *ValueListContext = nil
		if localctx != nil {
			t = localctx.(*ValueListContext)
		}
		return p.ValueList_Sempred(t, predIndex)

	case 7:
		var t *ScalarValueListContext = nil
		if localctx != nil {
			t = localctx.(*ScalarValueListContext)
		}
		return p.ScalarValueList_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ProxySQLParser) SettingList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ProxySQLParser) ValueList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ProxySQLParser) ScalarValueList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
