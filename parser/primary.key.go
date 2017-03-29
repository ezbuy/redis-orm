package parser

import (
	"fmt"
	"strings"
)

type PrimaryKey struct {
	Name       string
	FieldNames []string
	Fields     []*Field
	Obj        *MetaObject
}

func NewPrimaryKey(obj *MetaObject) *PrimaryKey {
	return &PrimaryKey{Obj: obj}
}

func (pk *PrimaryKey) IsSingleField() bool {
	if len(pk.Fields) == 1 {
		return true
	}
	return false
}

func (pk *PrimaryKey) FirstField() *Field {
	if len(pk.Fields) > 0 {
		return pk.Fields[0]
	}
	return nil
}

func (pk *PrimaryKey) IsAutocrement() bool {
	if len(pk.Fields) == 1 {
		return pk.Fields[0].Flags.Contains("autoinc")
	}
	return false
}

func (pk *PrimaryKey) IsRange() bool {
	c := len(pk.Fields)
	if c > 0 {
		return pk.Fields[c-1].IsNumber()
	}
	return false
}

func (pk *PrimaryKey) build() error {
	pk.Name = fmt.Sprintf("%sOf%sPK", strings.Join(pk.FieldNames, ""), pk.Obj.Name)
	for _, name := range pk.FieldNames {
		f := pk.Obj.FieldByName(name)
		if f == nil {
			return fmt.Errorf("%s field not exist", name)
		}
		f.Flags.Add("primary")
		pk.Fields = append(pk.Fields, f)
	}
	if len(pk.Fields) == 0 {
		return fmt.Errorf("primary key  not declare.")
	}
	return nil
}

func (pk *PrimaryKey) SQLColumn(driver string) string {
	switch strings.ToLower(driver) {
	case "mysql":
		columns := make([]string, 0, len(pk.Fields))
		for _, f := range pk.Fields {
			columns = append(columns, f.SQLName(driver))
		}
		return fmt.Sprintf("PRIMARY KEY(%s)", strings.Join(columns, ","))
	}
	return ""
}
