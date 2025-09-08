// Code generated from ProxySQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // ProxySQL
import "github.com/antlr4-go/antlr/v4"

// ProxySQLListener is a complete listener for a parse tree produced by ProxySQLParser.
type ProxySQLListener interface {
	antlr.ParseTreeListener

	// EnterConfiguration is called when entering the configuration production.
	EnterConfiguration(c *ConfigurationContext)

	// EnterSettingList is called when entering the settingList production.
	EnterSettingList(c *SettingListContext)

	// EnterSetting is called when entering the setting production.
	EnterSetting(c *SettingContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterValueList is called when entering the valueList production.
	EnterValueList(c *ValueListContext)

	// EnterScalarValue is called when entering the scalarValue production.
	EnterScalarValue(c *ScalarValueContext)

	// EnterScalarValueList is called when entering the scalarValueList production.
	EnterScalarValueList(c *ScalarValueListContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// ExitConfiguration is called when exiting the configuration production.
	ExitConfiguration(c *ConfigurationContext)

	// ExitSettingList is called when exiting the settingList production.
	ExitSettingList(c *SettingListContext)

	// ExitSetting is called when exiting the setting production.
	ExitSetting(c *SettingContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitValueList is called when exiting the valueList production.
	ExitValueList(c *ValueListContext)

	// ExitScalarValue is called when exiting the scalarValue production.
	ExitScalarValue(c *ScalarValueContext)

	// ExitScalarValueList is called when exiting the scalarValueList production.
	ExitScalarValueList(c *ScalarValueListContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)
}
