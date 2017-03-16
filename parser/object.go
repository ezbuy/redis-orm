package parser

import (
	"errors"
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
	Db  string
	Dbs []string
	//! database
	DbName  string
	DbTable string
	DbView  string
	//! fields
	fields       []*Field
	fieldNameMap map[string]*Field
	//! primary
	primary []*Field
	//! indexes
	uniques []*Index
	indexes []*Index
	ranges  []*Index
	//! relation
	Relation *Relation
	//! importSQL
	ImportSQL string
}

func NewMetaObject(packageName string) *MetaObject {
	return &MetaObject{
		Package:      packageName,
		GoPackage:    packageName,
		fieldNameMap: make(map[string]*Field),
		primary:      []*Field{},
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

func (o *MetaObject) PrimaryKey() []*Field {
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

func (o *MetaObject) Fields() []*Field {
	if o.Relation != nil {
		return o.Relation.Fields
	}
	return o.fields
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
					return errors.New(o.Name + " obj has " + err.Error())
				}
				o.fields[i] = f
				o.fieldNameMap[f.Name] = f
			}
		case "primary":
			fields := toStringSlice(val.([]interface{}))
			for _, field := range fields {
				if f := o.FieldByName(field); f != nil {
					o.primary = append(o.primary, f)
				}
			}
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
				return errors.New(o.Name + " obj has " + err.Error())
			}
			o.Relation = relation
		}
	}

	for _, field := range o.fields {
		if field.IsPrimary() {
			if len(o.primary) != 0 {
				return fmt.Errorf("primary key already defined: <%s> ", field.Name)
			}
			o.primary = append(o.primary, field)
		}
		if field.HasIndex() && field.IsNullable() {
			return fmt.Errorf("field <%s> should not be nullable for indexing", field.Name)
		}
	}

	for _, unique := range o.uniques {
		if err := unique.buildUnique(); err != nil {
			return err
		}
	}
	for _, index := range o.indexes {
		if err := index.buildIndex(); err != nil {
			return err
		}
	}
	for _, rg := range o.ranges {
		if err := rg.buildRange(); err != nil {
			return err
		}
	}
	return nil
}
