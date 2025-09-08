package proxycnf

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"proxycnf/parsing"
	"strconv"
	"strings"
)

type Configuration struct {
	settings []Setting
}

func Load(input string) *Configuration {
	inputStream := antlr.NewInputStream(input)
	lexer := parsing.NewProxySQLLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parsing.NewProxySQLParser(stream)

	l := newListener()
	antlr.ParseTreeWalkerDefault.Walk(l, p.Configuration())

	return &l.config
}

func (c *Configuration) Get(name string) Value {
	for _, setting := range c.settings {
		if setting.Name == name {
			return setting.Value
		}
	}

	return nil
}

func (c *Configuration) GetScalar(key string) (ScalarValue, error) {
	for _, v := range c.settings {
		if v.Name == key {
			if scalar, ok := v.Value.(ScalarValue); ok {
				return scalar, nil
			}
		}
	}

	return ScalarValue{}, fmt.Errorf("key %s not found or not a scalar", key)
}

func (c *Configuration) GetArray(key string) (Array, error) {
	for _, v := range c.settings {
		if v.Name == key {
			if arr, ok := v.Value.(Array); ok {
				return arr, nil
			}
		}
	}

	return nil, fmt.Errorf("key %s not found or not an array", key)
}

func (c *Configuration) GetList(key string) (List, error) {
	for _, v := range c.settings {
		if v.Name == key {
			if list, ok := v.Value.(List); ok {
				return list, nil
			}
		}
	}

	return nil, fmt.Errorf("key %s not found or not a list", key)
}

func (c *Configuration) GetGroup(name string) (Group, error) {
	for _, v := range c.settings {
		if v.Name == name {
			if group, ok := v.Value.(Group); ok {
				return group, nil
			}
		}
	}

	return nil, fmt.Errorf("key %s not found or not a group", name)
}

func (c *Configuration) ToMap() map[string]any {
	m := map[string]any{}
	for _, setting := range c.settings {
		m[setting.Name] = valueToAny(setting.Value)
	}
	return m
}

func (c *Configuration) Serialize() string {
	var s strings.Builder
	for _, setting := range c.settings {
		s.WriteString(setting.Serialize(0))
		s.WriteString("\n")
	}
	return s.String()
}

func valueToAny(v any) any {
	switch val := v.(type) {
	case ScalarValue:
		return val.Value()
	case Array:
		arr := make([]any, len(val))
		for i, item := range val {
			arr[i] = item.Value()
		}
		return arr
	case List:
		list := make([]any, len(val))
		for i, item := range val {
			list[i] = valueToAny(item)
		}
		return list
	case Group:
		group := make(map[any]any, len(val))
		for _, setting := range val {
			group[setting.Name] = valueToAny(setting.Value)
		}
		return group
	default:
		return val
	}
}

type ScalarValue struct {
	value any
}

// Value returns int64 for number.
func (v ScalarValue) Value() any {
	return v.value
}

func (v ScalarValue) Bool() (bool, error) {
	if b, ok := (v.value).(bool); ok {
		return b, nil
	}
	return false, fmt.Errorf("value is not a bool: %v", v)
}

func (v ScalarValue) MBool() bool {
	b, err := v.Bool()
	if err != nil {
		panic(err)
	}
	return b
}

func (v ScalarValue) Int() (int, error) {
	switch val := v.value.(type) {
	case int64:
		return int(val), nil
	case uint64:
		return int(val), nil
	default:
		return 0, fmt.Errorf("value is not int64 or uint64: %v", v)
	}
}

func (v ScalarValue) MInt() int {
	i, err := v.Int()
	if err != nil {
		panic(err)
	}
	return i
}

func (v ScalarValue) Int64() (int64, error) {
	switch val := v.value.(type) {
	case int64:
		return val, nil
	case uint64:
		return int64(val), nil
	default:
		return 0, fmt.Errorf("value is not int64 or uint64: %v", v)
	}
}
func (v ScalarValue) MInt64() int64 {
	i, err := v.Int64()
	if err != nil {
		panic(err)
	}
	return i
}
func (v ScalarValue) Float64() (float64, error) {
	if f, ok := (v.value).(float64); ok {
		return f, nil
	}
	return 0, fmt.Errorf("value is not a float64: %v", v)
}

func (v ScalarValue) MFloat64() float64 {
	f, err := v.Float64()
	if err != nil {
		panic(err)
	}
	return f
}

func (v ScalarValue) String() string {
	return fmt.Sprintf("%v", v.value)
}

func (v ScalarValue) Serialize(level int) string {
	var s strings.Builder
	s.WriteString(indent(level))

	switch val := v.value.(type) {
	case string:
		s.WriteString(strconv.Quote(val))
	case bool:
		s.WriteString(strconv.FormatBool(val))
	case uint8:
		s.WriteString(strconv.FormatUint(uint64(val), 10))
	case uint16:
		s.WriteString(strconv.FormatUint(uint64(val), 10))
	case uint32:
		s.WriteString(strconv.FormatUint(uint64(val), 10))
	case uint64:
		s.WriteString(strconv.FormatUint(val, 10))
	case int8:
		s.WriteString(strconv.FormatInt(int64(val), 10))
	case int16:
		s.WriteString(strconv.FormatInt(int64(val), 10))
	case int32:
		s.WriteString(strconv.FormatInt(int64(val), 10))
	case int64:
		s.WriteString(strconv.FormatInt(val, 10))
	case float32:
		s.WriteString(strconv.FormatFloat(float64(val), 'f', -1, 32))
	case float64:
		s.WriteString(strconv.FormatFloat(val, 'f', -1, 64))
	default:
		s.WriteString(fmt.Sprintf("%v", val))
	}

	return s.String()
}

// Value can be ScalarValue, Array, List, Group
type Value any

func serializeValue(v Value, level int) string {
	switch val := v.(type) {
	case ScalarValue:
		return val.Serialize(0)
	case Array:
		return val.Serialize(level)
	case List:
		return val.Serialize(level)
	case Group:
		return val.Serialize(level)
	default:
		return ""
	}
}

// List '( Value, ... )'
type List []Value

func (l List) Serialize(level int) string {
	prefix := indent(level)

	var s strings.Builder
	s.WriteString("\n")
	s.WriteString(prefix + "(\n")

	var values []string
	for _, v := range l {
		values = append(values, serializeValue(v, level+1))
	}
	s.WriteString(strings.Join(values, ","))
	s.WriteString(prefix + "\n)\n")

	return s.String()
}

// Array '[ ScalarValue, ... ]'
type Array []ScalarValue

func (a Array) Serialize(level int) string {
	var s strings.Builder
	s.WriteString("[\n")
	s.WriteString(indent(level))

	var values []string
	for _, v := range a {
		values = append(values, v.Serialize(level+1))
	}

	s.WriteString(strings.Join(values, ",\n"))
	s.WriteString("\n]\n")

	return s.String()
}

// Group '{ Setting, ... }'
type Group []Setting

func (g Group) lllGet(key string) Value {
	for _, setting := range g {
		if setting.Name == key {
			return setting.Value
		}
	}

	return nil
}

func (g Group) GetScalar(key string) (ScalarValue, error) {
	for _, setting := range g {
		if setting.Name == key {
			if scalar, ok := setting.Value.(ScalarValue); ok {
				return scalar, nil
			}
		}
	}

	return ScalarValue{}, fmt.Errorf("key %s not found or not a scalar", key)
}

func (g Group) Serialize(level int) string {
	prefix := indent(level)
	var s strings.Builder

	s.WriteString("\n")
	s.WriteString(prefix + "{\n")

	var settings []string
	for _, s := range g {
		settings = append(settings, s.Serialize(level+1))
	}
	s.WriteString(strings.Join(settings, "\n"))

	s.WriteString("\n")
	s.WriteString(prefix + "}")

	return s.String()
}

type Setting struct {
	Name  string
	Value Value
}

func (s Setting) Serialize(level int) string {
	var sb strings.Builder
	sb.WriteString(indent(level))
	sb.WriteString(fmt.Sprintf("%s = %s", s.Name, serializeValue(s.Value, level)))
	return sb.String()
}

type Serializer interface {
	Serialize(level int) string
}

func indent(level int) string {
	return strings.Repeat("    ", level)
}
