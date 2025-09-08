// Code generated from ProxySQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // ProxySQL
import "github.com/antlr4-go/antlr/v4"

// BaseProxySQLListener is a complete listener for a parse tree produced by ProxySQLParser.
type BaseProxySQLListener struct{}

var _ ProxySQLListener = &BaseProxySQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProxySQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProxySQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProxySQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProxySQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterConfiguration is called when production configuration is entered.
func (s *BaseProxySQLListener) EnterConfiguration(ctx *ConfigurationContext) {}

// ExitConfiguration is called when production configuration is exited.
func (s *BaseProxySQLListener) ExitConfiguration(ctx *ConfigurationContext) {}

// EnterSettingList is called when production settingList is entered.
func (s *BaseProxySQLListener) EnterSettingList(ctx *SettingListContext) {}

// ExitSettingList is called when production settingList is exited.
func (s *BaseProxySQLListener) ExitSettingList(ctx *SettingListContext) {}

// EnterSetting is called when production setting is entered.
func (s *BaseProxySQLListener) EnterSetting(ctx *SettingContext) {}

// ExitSetting is called when production setting is exited.
func (s *BaseProxySQLListener) ExitSetting(ctx *SettingContext) {}

// EnterName is called when production name is entered.
func (s *BaseProxySQLListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseProxySQLListener) ExitName(ctx *NameContext) {}

// EnterValue is called when production value is entered.
func (s *BaseProxySQLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseProxySQLListener) ExitValue(ctx *ValueContext) {}

// EnterValueList is called when production valueList is entered.
func (s *BaseProxySQLListener) EnterValueList(ctx *ValueListContext) {}

// ExitValueList is called when production valueList is exited.
func (s *BaseProxySQLListener) ExitValueList(ctx *ValueListContext) {}

// EnterScalarValue is called when production scalarValue is entered.
func (s *BaseProxySQLListener) EnterScalarValue(ctx *ScalarValueContext) {}

// ExitScalarValue is called when production scalarValue is exited.
func (s *BaseProxySQLListener) ExitScalarValue(ctx *ScalarValueContext) {}

// EnterScalarValueList is called when production scalarValueList is entered.
func (s *BaseProxySQLListener) EnterScalarValueList(ctx *ScalarValueListContext) {}

// ExitScalarValueList is called when production scalarValueList is exited.
func (s *BaseProxySQLListener) ExitScalarValueList(ctx *ScalarValueListContext) {}

// EnterArray is called when production array is entered.
func (s *BaseProxySQLListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseProxySQLListener) ExitArray(ctx *ArrayContext) {}

// EnterList is called when production list is entered.
func (s *BaseProxySQLListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseProxySQLListener) ExitList(ctx *ListContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseProxySQLListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseProxySQLListener) ExitGroup(ctx *GroupContext) {}
