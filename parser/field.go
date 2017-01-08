package parser

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/ezbuy/utils/container/set"
)

var (
	nullablePrimitiveSet = map[string]bool{
		"uint8":   true,
		"uint16":  true,
		"uint32":  true,
		"uint64":  true,
		"int8":    true,
		"int16":   true,
		"int32":   true,
		"int64":   true,
		"float32": true,
		"float64": true,
		"bool":    true,
		"string":  true,
	}
)

type Field struct {
	Name    string
	Type    string
	Flags   set.Set
	Attrs   map[string]string
	Comment string
	Obj     *MetaObject
}

func NewField() *Field {
	return &Field{
		Flags: set.NewStringSet(),
	}
}

var SupportedFieldTypes = map[string]string{
	"bool":      "bool",
	"int":       "int32",
	"int8":      "int8",
	"int16":     "int16",
	"int32":     "int32",
	"float32":   "float32",
	"float64":   "float64",
	"string":    "string",
	"datetime":  "datetime",
	"timestamp": "timestamp",
	"timeint":   "timeint",
}

func (f *Field) SetType(t string) error {
	st, ok := SupportedFieldTypes[t]
	if !ok {
		return fmt.Errorf("%s type not support.", t)
	}
	//! special type convert
	switch f.Obj.Db {
	case "mysql":
	case "mssql":
	case "redis":
	case "mongo":
	case "elastic":
	}
	f.Type = st
	return nil
}

func (f *Field) IsPrimary() bool {
	return f.Flags.Contains("primary")
}

func (f *Field) IsAutoIncrement() bool {
	if f.IsPrimary() {
		if f.Flags.Contains("autoinc") {
			return true
		}
		if !f.Flags.Contains("noinc") {
			return true
		}
	}
	return false
}

func (f *Field) IsNullable() bool {
	return !f.IsPrimary() && f.Flags.Contains("nullable")
}

func (f *Field) IsUnique() bool {
	return f.Flags.Contains("unique")
}

func (f *Field) IsRange() bool {
	return f.Flags.Contains("range")
}

func (f *Field) IsOrder() bool {
	return f.Flags.Contains("order")
}

func (f *Field) IsIndex() bool {
	return f.Flags.Contains("index")
}

func (f *Field) IsFullText() bool {
	return f.Flags.Contains("fulltext")
}

func (f *Field) HasIndex() bool {
	return f.Flags.Contains("unique") ||
		f.Flags.Contains("index") ||
		f.Flags.Contains("range") ||
		f.Flags.Contains("order")
}

func (f *Field) GetType() string {
	st := f.Type
	if transform := f.GetTransformType(); transform != nil {
		st = transform.TypeTarget
	}

	if f.IsNullable() {
		if st == "time.Time" {
			st = "*time.Time"
		}
	}
	return st
}

func (f *Field) IsNullablePrimitive() bool {
	return f.IsNullable() && nullablePrimitiveSet[f.GetType()]
}

func (f *Field) GetNullSQLType() string {
	t := f.GetType()
	if t == "bool" {
		return "NullBool"
	} else if t == "string" {
		return "NullString"
	} else if strings.HasPrefix(t, "int") {
		return "NullInt64"
	} else if strings.HasPrefix(t, "float") {
		return "NullFloat64"
	}
	return t
}

func (f *Field) NullSQLTypeValue() string {
	t := f.GetType()
	if t == "bool" {
		return "Bool"
	} else if t == "string" {
		return "String"
	} else if strings.HasPrefix(t, "int") {
		return "Int64"
	} else if strings.HasPrefix(t, "float") {
		return "Float64"
	}
	panic("unsupported null sql type: " + t)
}

func (f *Field) NullSQLTypeNeedCast() bool {
	t := f.GetType()
	if strings.HasPrefix(t, "int") && t != "int64" {
		return true
	} else if strings.HasPrefix(t, "float") && t != "float64" {
		return true
	}
	return false
}

type Transform struct {
	TypeOrigin  string
	ConvertTo   string
	TypeTarget  string
	ConvertBack string
}

// convert `TypeOrigin` in datebase to `TypeTarget` when quering
// convert `TypeTarget` back to `TypeOrigin` when updating/inserting
var transformMap = map[string]Transform{
	"mysql_timestamp": { // TIMESTAMP (string, UTC)
		"string", `orm.TimeParse(%v)`,
		"time.Time", `orm.TimeFormat(%v)`,
	},
	"mysql_timeint": { // INT(11)
		"int64", "time.Unix(%v, 0)",
		"time.Time", "%v.Unix()",
	},
	"mysql_datetime": { // DATETIME (string, localtime)
		"string", "orm.TimeParseLocalTime(%v)",
		"time.Time", "orm.TimeToLocalTime(%v)",
	},
	"redis_timestamp": { // TIMESTAMP (string, UTC)
		"string", `orm.TimeParse(%v)`,
		"time.Time", `orm.TimeFormat(%v)`,
	},
	"redis_timeint": { // INT(11)
		"int64", "time.Unix(%v, 0)",
		"time.Time", "%v.Unix()",
	},
	"redis_datetime": { // DATETIME (string, localtime)
		"string", "orm.TimeParseLocalTime(%v)",
		"time.Time", "orm.TimeToLocalTime(%v)",
	},
}

func (f *Field) IsNeedTransform() bool {
	return f.GetTransformType() != nil
}

func (f *Field) GetTransformType() *Transform {
	key := fmt.Sprintf("%v_%v", f.Obj.Db, f.Type)
	t, ok := transformMap[key]
	if !ok {
		return nil
	}
	return &t
}

func (f *Field) GetTransformValue(prefix string) string {
	t := f.GetTransformType()
	if t == nil {
		return prefix + f.Name
	}
	return fmt.Sprintf(t.ConvertBack, prefix+f.Name)
}

func (f *Field) GetTag() string {
	tags := map[string]bool{}
	for _, db := range f.Obj.Dbs {
		switch db {
		case "mongo":
			tags["bson"] = true
			tags["json"] = true
		case "redis":
			tags["json"] = false
		case "elastic":
			tags["json"] = false
		case "mysql":
			tags["db"] = false
		case "mssql":
			tags["db"] = true
		}
	}

	tagstr := []string{}
	for tag, camel := range tags {
		if val, ok := f.Attrs[tag+"Tag"]; ok {
			tagstr = append(tagstr, fmt.Sprintf("%s:\"%s\"", tag, val))
		} else {
			if camel {
				tagstr = append(tagstr, fmt.Sprintf("%s:\"%s\"", tag, f.Name))
			} else {
				tagstr = append(tagstr, fmt.Sprintf("%s:\"%s\"", tag, Camel2Name(f.Name)))
			}
		}
	}
	sortstr := sort.StringSlice(tagstr)
	sort.Sort(sortstr)
	if len(sortstr) != 0 {
		return "`" + strings.Join(sortstr, " ") + "`"
	}
	return ""
}

func (f *Field) Read(data map[interface{}]interface{}) error {
	foundName := false
	for k, v := range data {
		key := k.(string)

		if isUpperCase(key[0:1]) {
			if foundName {
				return errors.New("invalid field name: " + key)
			}
			f.Name = key
			if err := f.SetType(v.(string)); err != nil {
				return err
			}
			continue
		}

		switch key {
		case "comment":
			f.Comment = v.(string)
		case "attrs":
			attrs := make(map[string]string)
			for ki, vi := range v.(map[interface{}]interface{}) {
				attrs[ki.(string)] = vi.(string)
			}
			f.Attrs = attrs
		case "flags":

			for _, flag := range v.([]interface{}) {
				f.Flags.Add(flag.(string))
			}
		default:
			return errors.New("invalid field name: " + key)
		}
	}
	if f.IsUnique() {
		index := NewIndex(f.Obj)
		index.FieldNames = []string{f.Name}
		f.Obj.Uniques = append(f.Obj.Uniques, index)
	}
	if f.IsIndex() {
		index := NewIndex(f.Obj)
		index.FieldNames = []string{f.Name}
		f.Obj.Indexes = append(f.Obj.Indexes, index)
	}
	if f.IsRange() {
		index := NewIndex(f.Obj)
		index.FieldNames = []string{f.Name}
		f.Obj.Ranges = append(f.Obj.Ranges, index)
	}
	if f.IsOrder() {
		index := NewIndex(f.Obj)
		index.FieldNames = []string{f.Name}
		f.Obj.Orders = append(f.Obj.Orders, index)
	}
	return nil
}
