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

func (pk *PrimaryKey) build() error {
	pk.Name = strings.Join(pk.FieldNames, "_")
	for _, name := range pk.FieldNames {
		f := pk.Obj.FieldByName(name)
		if f == nil {
			return fmt.Errorf("%s field not exist", name)
		}
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
