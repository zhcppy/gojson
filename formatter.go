package gojson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Formatter struct {
	KeyColor        *color.Color
	StringColor     *color.Color
	BoolColor       *color.Color
	NumberColor     *color.Color
	NullColor       *color.Color
	StringMaxLength int
	DisabledColor   bool
	Indent          int
	Newline         string
}

func NewFormatter() *Formatter {
	return &Formatter{
		KeyColor:        color.New(color.FgBlue, color.Bold),
		StringColor:     color.New(color.FgGreen, color.Bold),
		BoolColor:       color.New(color.FgYellow, color.Bold),
		NumberColor:     color.New(color.FgCyan, color.Bold),
		NullColor:       color.New(color.FgBlack, color.Bold),
		StringMaxLength: 0,
		DisabledColor:   false,
		Indent:          2,
		Newline:         "\n",
	}
}

func Marshal(v interface{}) (string, error) {
	return NewFormatter().Marshal(v)
}

func MustMarshal(v interface{}) string {
	res, err := NewFormatter().Marshal(v)
	if err != nil {
		return err.Error()
	}
	return res
}

func Format(data []byte) (string, error) {
	return NewFormatter().Format(data)
}

func (f *Formatter) Marshal(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return f.Format(data)
}

func (f *Formatter) Format(data []byte) (string, error) {
	var v interface{}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(&v); err != nil {
		return "", err
	}
	return f.addColor(v, 1), nil
}

func (f *Formatter) sprintColor(c *color.Color, s string) string {
	if f.DisabledColor || c == nil {
		return fmt.Sprint(s)
	}
	return c.SprintFunc()(s)
}

func (f *Formatter) addColor(v interface{}, depth int) string {
	switch val := v.(type) {
	case string:
		return f.processString(val)
	case float64:
		return f.sprintColor(f.NumberColor, strconv.FormatFloat(val, 'f', -1, 64))
	case json.Number:
		return f.sprintColor(f.NumberColor, string(val))
	case bool:
		return f.sprintColor(f.BoolColor, strconv.FormatBool(val))
	case nil:
		return f.sprintColor(f.NullColor, "null")
	case map[string]interface{}:
		return f.processMap(val, depth)
	case []interface{}:
		return f.processArray(val, depth)
	default:
		return ""
	}
}

func (f *Formatter) processString(res string) string {
	sr := []rune(res)
	if f.StringMaxLength != 0 && len(sr) >= f.StringMaxLength {
		res = string(sr[0:f.StringMaxLength]) + "..."
	}

	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(res); err != nil {
		return err.Error()
	}

	return f.sprintColor(f.StringColor, strings.TrimSuffix(string(buf.Bytes()), "\n"))
}

func (f *Formatter) processMap(m map[string]interface{}, depth int) string {
	if len(m) == 0 {
		return "{}"
	}

	currentIndent := f.generateIndent(depth - 1)
	nextIndent := f.generateIndent(depth)

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var rows []string
	for _, key := range keys {
		val := m[key]
		k := f.sprintColor(f.KeyColor, fmt.Sprintf(`"%s"`, key))
		v := f.addColor(val, depth+1)

		valueIndent := " "
		if f.Newline == "" {
			valueIndent = ""
		}
		rows = append(rows, fmt.Sprintf("%s%s:%s%s", nextIndent, k, valueIndent, v))
	}

	return fmt.Sprintf("{%s%s%s%s}", f.Newline, strings.Join(rows, ","+f.Newline), f.Newline, currentIndent)
}

func (f *Formatter) processArray(a []interface{}, depth int) string {
	if len(a) == 0 {
		return "[]"
	}

	currentIndent := f.generateIndent(depth - 1)
	nextIndent := f.generateIndent(depth)

	var rows []string
	for _, val := range a {
		c := f.addColor(val, depth+1)
		row := nextIndent + c
		rows = append(rows, row)
	}

	return fmt.Sprintf("[%s%s%s%s]", f.Newline, strings.Join(rows, ","+f.Newline), f.Newline, currentIndent)
}

func (f *Formatter) generateIndent(depth int) string {
	return strings.Repeat(" ", f.Indent*depth)
}
