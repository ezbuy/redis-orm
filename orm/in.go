package orm

import (
	"fmt"
	"strings"
)

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
