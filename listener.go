package proxycnf

import (
	"fmt"
	"github.com/vczyh/proxycnf/parsing"
	"slices"
	"strconv"
)

type Listener struct {
	*parsing.BaseProxySQLListener
	config Configuration
	stack  []any
}

func newListener() *Listener {
	base := new(parsing.BaseProxySQLListener)
	return &Listener{
		BaseProxySQLListener: base,
	}
}

func (l *Listener) ExitConfiguration(c *parsing.ConfigurationContext) {
	var settings []Setting
	for v := l.pop(); v != nil; v = l.pop() {
		settings = append(settings, v.(Setting))
	}
	slices.Reverse(settings)
	l.config.settings = settings
}

func (l *Listener) EnterSetting(c *parsing.SettingContext) {
	l.push(Setting{Name: c.Name().GetText()})
}

func (l *Listener) ExitSetting(c *parsing.SettingContext) {
	val := l.pop()
	setting := l.pop().(Setting)

	setting.Value = val
	l.push(setting)
}

func (l *Listener) EnterArray(c *parsing.ArrayContext) {
	l.push(Array{})
}

func (l *Listener) ExitArray(c *parsing.ArrayContext) {
	var valueList []ScalarValue
	for v := l.pop(); v != nil; v = l.pop() {
		if array, ok := v.(Array); ok {
			array = valueList
			l.push(array)
			break
		} else {
			valueList = append([]ScalarValue{v.(ScalarValue)}, valueList...)
		}
	}
}

func (l *Listener) EnterList(c *parsing.ListContext) {
	l.push(List{})
}
func (l *Listener) ExitList(c *parsing.ListContext) {
	var valueList []Value
	for v := l.pop(); v != nil; v = l.pop() {
		if list, ok := v.(List); ok {
			list = valueList
			l.push(list)
			break
		} else {
			valueList = append([]Value{v}, valueList...)
		}
	}
}

func (l *Listener) EnterGroup(c *parsing.GroupContext) {
	l.push(Group{})
}

func (l *Listener) ExitGroup(c *parsing.GroupContext) {
	var settingList []Setting
	for v := l.pop(); v != nil; v = l.pop() {
		if group, ok := v.(Group); ok {
			group = settingList
			l.push(group)
			break
		} else {
			settingList = append([]Setting{v.(Setting)}, settingList...)
		}
	}
}

func (l *Listener) EnterScalarValue(c *parsing.ScalarValueContext) {
	text := c.GetText()

	var val any
	switch {
	case c.BOOLEAN() != nil:
		if text == "true" || text == "TRUE" {
			val = true
		} else if text == "false" || text == "FALSE" {
			val = false
		}
	case c.NUMBER() != nil || c.HEX() != nil:
		n, err := strconv.ParseInt(text, 0, 64)
		if err != nil {
			val = fmt.Sprintf("invalid number: %s", text)
		}
		val = n
	case c.FLOAT() != nil:
		f, err := strconv.ParseFloat(text, 64)
		if err != nil {
			val = fmt.Sprintf("invalid float: %s", text)
		}
		val = f
	case c.STRING() != nil:
		s, err := strconv.Unquote(text)
		if err != nil {
			val = fmt.Sprintf("invalid string: %s", text)
		}
		val = s
	}

	l.push(ScalarValue{value: val})
}

func (l *Listener) push(v any) {
	l.stack = append(l.stack, v)
}

func (l *Listener) pop() any {
	if len(l.stack) == 0 {
		return nil
	}
	v := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return v
}

func (l *Listener) peek() any {
	if len(l.stack) == 0 {
		return nil
	}
	return l.stack[len(l.stack)-1]
}
