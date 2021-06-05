package orm

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

type FieldMultiIN struct {
	rawFieldsLen int
	in           *FieldIN
}

func NewMultiFieldIN(fields []string) *FieldMultiIN {
	var b bytes.Buffer
	b.WriteByte('(')
	b.WriteString(strings.Join(fields, ","))
	b.WriteByte(')')
	return &FieldMultiIN{
		rawFieldsLen: len(fields),
		in: &FieldIN{
			Field: b.String(),
		},
	}
}

func (in *FieldMultiIN) Add(v []interface{}) error {
	if in.rawFieldsLen == 0 || len(v) == 0 {
		return errors.New("builder: fields length and passed-in value length should above zero")
	}
	if len(v)%in.rawFieldsLen != 0 {
		return errors.New("builder: passed-in value length should be integer multiple than fields length")
	}
	in.in.Params = append(in.in.Params, v...)
	var b bytes.Buffer
	b.WriteByte('(')
	for index := range v {
		if index == len(v)-1 {
			b.WriteByte('?')
			continue
		}
		b.WriteString("?,")
	}
	b.WriteByte(')')
	in.in.holders = append(in.in.holders, b.String())
	return nil
}

func (in *FieldMultiIN) SQLFormat() string {
	return in.in.SQLFormat()
}

func (in *FieldMultiIN) SQLParams() []interface{} {
	return in.in.SQLParams()
}

type FieldIN struct {
	Field   string
	Params  []interface{}
	holders []string
}

func NewFieldIN(field string) *FieldIN {
	return &FieldIN{
		Field:   field,
		Params:  []interface{}{},
		holders: []string{},
	}
}

func (in *FieldIN) Add(v interface{}) *FieldIN {
	in.Params = append(in.Params, v)
	in.holders = append(in.holders, "?")
	return in
}

func (in *FieldIN) SQLFormat() string {
	if len(in.Params) == 0 {
		return ""
	}
	return fmt.Sprintf("%s IN (%s)", in.Field, strings.Join(in.holders, ","))
}

func (in *FieldIN) SQLParams() []interface{} {
	return in.Params
}
