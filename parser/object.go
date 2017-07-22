package parser

import (
	"fmt"
	"sort"
	"strings"
)

type MetaObject struct {
	//! package name
	Package   string
	GoPackage string
	//! model name
	Name string
	Tag  string
	//! dbs
	Db      string
	Dbs     []string
	comment string
	//! database
	DbName  string
	DbTable string
	DbView  string
	//! fields
	fields       []*Field
	fieldNameMap map[string]*Field
	//! primary
	primary *PrimaryKey
	//! indexes
	uniques []*Index
	indexes []*Index
	ranges  []*Index
	//! relation
	Relation *Relation
	//! importSQL
	ImportSQL string
	//! elastic
	ElasticIndexAll bool
}

func NewMetaObject(packageName string) *MetaObject {
	return &MetaObject{
		Package:      packageName,
		GoPackage:    packageName,
		fieldNameMap: make(map[string]*Field),
		uniques:      []*Index{},
		indexes:      []*Index{},
		ranges:       []*Index{},
	}
}

func (o *MetaObject) FieldByName(name string) *Field {
	if f, ok := o.fieldNameMap[name]; ok {
		return f
	}
	return nil
}

func (o *MetaObject) PrimaryField() *Field {
	for _, f := range o.Fields() {
		if f.IsPrimary() {
			return f
		}
	}
	return nil
}

func (o *MetaObject) PrimaryKey() *PrimaryKey {
	return o.primary
}

func (o *MetaObject) DbContains(db string) bool {
	for _, v := range o.Dbs {
		if strings.ToLower(v) == strings.ToLower(db) {
			return true
		}
	}
	return false
}

func (o *MetaObject) DbSource() string {
	if o.DbTable != "" {
		return o.DbTable
	}
	if o.DbView != "" {
		return o.DbView
	}
	return ""
}

func (o *MetaObject) FromDB() string {
	switch o.Db {
	case "mssql":
		return fmt.Sprintf("[dbo].[%s]", o.DbSource())
	}
	return fmt.Sprintf("%s", o.DbSource())
}

func (o *MetaObject) Fields() []*Field {
	if o.Relation != nil {
		return o.Relation.Fields()
	}
	return o.fields
}

func (o *MetaObject) NoneIncrementFields() []*Field {
	if o.Relation != nil {
		return o.Relation.NoneIncrementFields()
	}
	fields := make([]*Field, 0, len(o.fields))
	for _, f := range o.fields {
		if f.IsAutoIncrement() == false {
			fields = append(fields, f)
		}
	}
	return fields
}

func (o *MetaObject) Uniques() []*Index {
	sort.Sort(IndexArray(o.uniques))
	return o.uniques
}

func (o *MetaObject) Indexes() []*Index {
	sort.Sort(IndexArray(o.indexes))
	return o.indexes
}

func (o *MetaObject) Ranges() []*Index {
	sort.Sort(IndexArray(o.ranges))
	return o.ranges
}
func (o *MetaObject) LastField() *Field {
	return o.fields[len(o.fields)-1]
}

func (o *MetaObject) Read(name string, data map[string]interface{}) error {
	o.Name = name
	hasType := false
	for key, val := range data {
		switch key {
		case "db":
			o.Db = val.(string)
			dbs := []string{}
			dbs = append(dbs, o.Db)
			dbs = append(dbs, o.Dbs...)
			o.Dbs = dbs
			hasType = true
		case "dbs":
			o.Dbs = toStringSlice(val.([]interface{}))
			if len(o.Dbs) != 0 {
				o.Db = o.Dbs[0]
			}
			hasType = true
		}
	}
	if hasType {
		delete(data, "db")
		delete(data, "dbs")
	}

	for key, val := range data {
		switch key {
		case "tag":
			tag := val.(int)
			o.Tag = fmt.Sprint(tag)
		case "dbname":
			o.DbName = val.(string)
		case "dbtable":
			o.DbTable = val.(string)
		case "dbview":
			o.DbView = val.(string)
		case "comment":
			o.comment = val.(string)

		case "importSQL":
			o.ImportSQL = val.(string)
		case "fields":
			fieldData := val.([]interface{})
			o.fields = make([]*Field, len(fieldData))
			for i, field := range fieldData {
				f := NewField()
				f.Obj = o
				err := f.Read(field.(map[interface{}]interface{}))
				if err != nil {
					return fmt.Errorf("object (%s) %s", o.Name, err.Error())
				}
				o.fields[i] = f
				o.fieldNameMap[f.Name] = f
			}
		case "primary":
			o.primary = NewPrimaryKey(o)
			o.primary.FieldNames = toStringSlice(val.([]interface{}))
		case "uniques":
			for _, i := range val.([]interface{}) {
				if len(i.([]interface{})) == 0 {
					continue
				}
				index := NewIndex(o)
				index.FieldNames = toStringSlice(i.([]interface{}))
				o.uniques = append(o.uniques, index)
			}
		case "indexes":
			for _, i := range val.([]interface{}) {
				if len(i.([]interface{})) == 0 {
					continue
				}
				index := NewIndex(o)
				index.FieldNames = toStringSlice(i.([]interface{}))
				o.indexes = append(o.indexes, index)
			}
		case "ranges":
			for _, i := range val.([]interface{}) {
				if len(i.([]interface{})) == 0 {
					continue
				}
				index := NewIndex(o)
				index.FieldNames = toStringSlice(i.([]interface{}))
				o.ranges = append(o.ranges, index)
			}
		case "relation":
			relation := NewRelation(o)
			err := relation.Read(val.(map[interface{}]interface{}))
			if err != nil {
				return fmt.Errorf("object (%s) %s", o.Name, err.Error())
			}
			o.Relation = relation

		case "es_index_all":
			o.ElasticIndexAll = val.(bool)
		}
	}

	for _, field := range o.fields {
		if field.IsPrimary() {
			if o.primary == nil {
				o.primary = NewPrimaryKey(o)
				o.primary.FieldNames = []string{}
			}
			o.primary.FieldNames = append(o.primary.FieldNames, field.Name)
		}
		if field.HasIndex() && field.IsNullable() {
			return fmt.Errorf("object (%s) field (%s) should not be nullable for indexing", o.Name, field.Name)
		}
	}

	if o.Relation == nil {
		if o.primary == nil {
			if o.DbContains("mysql") || o.DbContains("mssql") {
				return fmt.Errorf("object (%s) needs a primary key declare.", o.Name)
			}
		} else {
			if err := o.primary.build(); err != nil {
				return fmt.Errorf("object (%s) %s", o.Name, err.Error())
			}

			if o.primary.IsRange() {
				index := NewIndex(o)
				index.FieldNames = o.primary.FieldNames
				o.ranges = append(o.ranges, index)
			}
		}
	}

	for _, unique := range o.uniques {
		if err := unique.buildUnique(); err != nil {
			return fmt.Errorf("object (%s) %s", o.Name, err.Error())
		}
	}
	for _, index := range o.indexes {
		if err := index.buildIndex(); err != nil {
			return fmt.Errorf("object (%s) %s", o.Name, err.Error())
		}
	}
	for _, rg := range o.ranges {
		if err := rg.buildRange(); err != nil {
			return fmt.Errorf("object (%s) %s", o.Name, err.Error())
		}
	}
	return nil
}

func (m *MetaObject) ElasticIndexTypeName() string {
	if m.DbTable != "" {
		return m.DbTable
	}

	return Camel2Name(m.Name)
}

func (m *MetaObject) Comment() string {
	if m.comment != "" {
		return m.comment
	}

	return m.DbTable
}
