package parser

import (
	"fmt"
	"strings"
)

type Index struct {
	Name       string
	Fields     []*Field
	FieldNames []string
	Obj        *MetaObject
}

func NewIndex(obj *MetaObject) *Index {
	return &Index{Obj: obj}
}

func (idx *Index) build() error {
	idx.Name = fmt.Sprintf("%sOf%s", strings.Join(idx.FieldNames, ""), idx.Obj.Name)
	for _, name := range idx.FieldNames {
		f := idx.Obj.FieldByName(name)
		if f == nil {
			return fmt.Errorf("%s field not exist", name)
		}
		idx.Fields = append(idx.Fields, f)
	}
	return nil
}

func (idx *Index) GetRelation(storetype, valuetype, modeltype string) *Relation {
	relation := NewRelation(idx.Obj)
	relation.Name = idx.Name
	relation.StoreType = storetype
	relation.ValueType = valuetype
	relation.ModelType = modeltype
	relation.build()
	return nil
}
