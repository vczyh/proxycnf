package proxycnf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type valueType uint8

const (
	valueTypeInt valueType = iota
	valueTypeString
	valueTypeBool
)

func TestConfiguration_GetScalar(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		key       string
		valueType valueType
		expected  any
	}{
		{
			name:      "Get scalar string",
			input:     `host: "localhost"`,
			key:       "host",
			valueType: valueTypeString,
			expected:  "localhost",
		},
		{
			name:      "Get scalar int",
			input:     `port: 3306`,
			key:       "port",
			valueType: valueTypeInt,
			expected:  3306,
		},
		{
			name:      "Get scalar bool(true)",
			input:     `enabled: true`,
			key:       "enabled",
			valueType: valueTypeBool,
			expected:  true,
		},
		{
			name:      "Get scalar bool(false)",
			input:     `enabled: false`,
			key:       "enabled",
			valueType: valueTypeBool,
			expected:  false,
		},
		{
			name:      "Get scalar float",
			input:     `threshold: 0.75`,
			key:       "threshold",
			valueType: valueTypeString,
			expected:  "0.75",
		},
	}

	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			c := Load(ts.input)

			value, err := c.GetScalar(ts.key)
			if err != nil {
				t.Fatal(err)
			}

			switch ts.valueType {
			case valueTypeString:
				assert.Equal(t, ts.expected, value.String())
			case valueTypeInt:
				assert.Equal(t, ts.expected, value.MInt())
			case valueTypeBool:
				assert.Equal(t, ts.expected, value.MBool())
			default:
				t.Fatalf("unsupported value type: %v", ts.valueType)
			}
		})
	}
}

func TestConfiguration_GetArray(t *testing.T) {
	input := `flags = [true, "a", 255]`
	c := Load(input)

	arr, err := c.GetArray("flags")
	assert.NoError(t, err)
	assert.Len(t, arr, 3)

	assert.Equal(t, true, arr[0].MBool())
	assert.Equal(t, "a", arr[1].String())
	assert.Equal(t, 255, arr[2].MInt())
}

func TestConfiguration_GetGroup(t *testing.T) {
	input := `
group1 = {
  host: "127.0.0.1"
  port: 3306
  enabled: true
}
`
	c := Load(input)

	group, err := c.GetGroup("group1")
	assert.NoError(t, err)
	assert.Len(t, group, 3)

	m := map[string]any{}
	for _, setting := range group {
		m[setting.Name] = setting.Value
	}

	host, ok := m["host"].(ScalarValue)
	assert.True(t, ok)
	assert.Equal(t, "127.0.0.1", host.String())

	port, ok := m["port"].(ScalarValue)
	assert.True(t, ok)
	assert.Equal(t, 3306, port.MInt())

	enabled, ok := m["enabled"].(ScalarValue)
	assert.True(t, ok)
	assert.Equal(t, true, enabled.MBool())
}

func TestConfiguration_GetList(t *testing.T) {
	input := `
servers = (
  "10.0.0.1",
  "10.0.0.2",
  "10.0.0.3"
)
`
	c := Load(input)

	list, err := c.GetList("servers")
	assert.NoError(t, err)
	assert.Len(t, list, 3)

	assert.Equal(t, "10.0.0.1", list[0].(ScalarValue).String())
	assert.Equal(t, "10.0.0.2", list[1].(ScalarValue).String())
	assert.Equal(t, "10.0.0.3", list[2].(ScalarValue).String())
}

func TestConfiguration_toMap(t *testing.T) {
	input := `
host: "localhost"
port: 3306
enabled: true
flags = [true, "a", 255]
servers = (
  "10.0.0.1",
  "10.0.0.2"
)
group1 = {
  foo: "bar"
  num: 42
}
`
	c := Load(input)
	m := c.ToMap()

	assert.Equal(t, "localhost", m["host"])
	assert.Equal(t, int64(3306), m["port"])
	assert.Equal(t, true, m["enabled"])

	arr, ok := m["flags"].([]any)
	assert.True(t, ok)
	assert.Equal(t, []any{true, "a", int64(255)}, arr)

	list, ok := m["servers"].([]any)
	assert.True(t, ok)
	assert.Equal(t, []any{"10.0.0.1", "10.0.0.2"}, list)

	group, ok := m["group1"].(map[any]any)
	assert.True(t, ok)
	assert.Equal(t, "bar", group["foo"])
	assert.Equal(t, int64(42), group["num"])
}
