package parser

import (
	"errors"
	"fmt"
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
	//! indexes
	Uniques []*Index
	Indexes []*Index
	Ranges  []*Index
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
		Uniques:      []*Index{},
		Indexes:      []*Index{},
		Ranges:       []*Index{},
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
		case "uniques":
			for _, i := range val.([]interface{}) {
				if len(i.([]interface{})) == 0 {
					continue
				}
				index := NewIndex(o)
				index.FieldNames = toStringSlice(i.([]interface{}))
				o.Uniques = append(o.Uniques, index)
			}
		case "indexes":
			for _, i := range val.([]interface{}) {
				if len(i.([]interface{})) == 0 {
					continue
				}
				index := NewIndex(o)
				index.FieldNames = toStringSlice(i.([]interface{}))
				o.Indexes = append(o.Indexes, index)
			}
		case "ranges":
			for _, i := range val.([]interface{}) {
				if len(i.([]interface{})) == 0 {
					continue
				}
				index := NewIndex(o)
				index.FieldNames = toStringSlice(i.([]interface{}))
				o.Ranges = append(o.Ranges, index)
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

	for _, unique := range o.Uniques {
		if err := unique.buildUnique(); err != nil {
			return err
		}
	}
	for _, index := range o.Indexes {
		if err := index.buildIndex(); err != nil {
			return err
		}
	}
	for _, rg := range o.Ranges {
		if err := rg.buildRange(); err != nil {
			return err
		}
	}
	return nil
}
