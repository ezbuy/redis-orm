package parser

import (
	"errors"
	"fmt"
	"log"
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
	Name      string
	Type      string
	sqlType   string
	sqlColumn string
	Size      int
	Flags     set.Set
	Attrs     map[string]string
	Comment   string
	Validator string
	Obj       *MetaObject
	ESIndex   ESIndex
	Default   interface{}
	PBField   *PBField
}

type PBField struct {
	Name string
	Type string
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
	"int64":     "int64",
	"uint":      "uint32",
	"uint8":     "uint8",
	"uint16":    "uint16",
	"uint32":    "uint32",
	"uint64":    "uint64",
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

func (f *Field) FieldName() string {
	if f.Obj.DbContains("mysql") {
		return fmt.Sprintf("`%s`", f.ColumnName())
	}
	return f.ColumnName()
}

func (f *Field) ColumnName() string {
	if f.sqlColumn != "" {
		return f.sqlColumn
	}
	return Camel2Name(f.Name)
}

func (f *Field) IsPrimary() bool {
	return f.Flags.Contains("primary")
}

func (f *Field) IsAutoIncrement() bool {
	if f.Flags.Contains("autoinc") {
		return true
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

func (f *Field) IsIndex() bool {
	return f.Flags.Contains("index")
}

func (f *Field) IsFullText() bool {
	return f.Flags.Contains("fulltext")
}

func (f *Field) IsSetPBMapping() bool {
	return f.PBField != nil
}

func (f *Field) IsEncode() bool {
	if f.IsString() {
		return f.Flags.Contains("encode") || f.Flags.Contains("base64")
	}
	return false
}

func (f *Field) IsNumber() bool {
	if transform := f.GetTransform(); transform != nil {
		if strings.HasPrefix(transform.TypeOrigin, "uint") ||
			strings.HasPrefix(transform.TypeOrigin, "int") ||
			strings.HasPrefix(transform.TypeOrigin, "bool") ||
			strings.HasPrefix(transform.TypeOrigin, "float") {
			return true
		}
	}
	if strings.HasPrefix(f.Type, "uint") ||
		strings.HasPrefix(f.Type, "int") ||
		strings.HasPrefix(f.Type, "bool") ||
		strings.HasPrefix(f.Type, "float") {
		return true
	}
	return false
}

func (f *Field) IsBool() bool {
	if transform := f.GetTransform(); transform != nil {
		return strings.HasPrefix(transform.TypeOrigin, "bool")
	}
	return strings.HasPrefix(f.Type, "bool")
}

func (f *Field) IsString() bool {
	if transform := f.GetTransform(); transform != nil {
		if strings.HasPrefix(transform.TypeOrigin, "string") {
			return true
		}
	}
	if strings.HasPrefix(f.Type, "string") {
		return true
	}
	return false
}

func (f *Field) IsTime() bool {
	switch f.Type {
	case "datetime", "timestamp", "timeint":
		return true
	}
	return false
}

func (f *Field) HasIndex() bool {
	return f.Flags.Contains("unique") ||
		f.Flags.Contains("index") ||
		f.Flags.Contains("range")
}

func (f *Field) GetType() string {
	st := f.Type
	if transform := f.GetTransform(); transform != nil {
		st = transform.TypeTarget
	}

	if f.IsNullable() {
		if st == "time.Time" {
			st = "*time.Time"
		}
	}
	return st
}

func (f *Field) GetNames() string {
	return CamelName(f.Name) + "s"
}

func (f *Field) IsNullablePrimitive() bool {
	return f.IsNullable() && nullablePrimitiveSet[f.GetType()]
}

func (f *Field) GetNullSQLType() string {
	origin_type := f.Type
	if transform := f.GetTransform(); transform != nil {
		origin_type = transform.TypeOrigin
	}

	if f.IsNullable() {
		if origin_type == "bool" {
			return "NullBool"
		} else if origin_type == "string" {
			return "NullString"
		} else if strings.HasPrefix(origin_type, "int") {
			return "NullInt64"
		} else if strings.HasPrefix(origin_type, "float") {
			return "NullFloat64"
		}
	}
	return origin_type
}

func (f *Field) NullSQLTypeValue() string {
	origin_type := f.Type
	if transform := f.GetTransform(); transform != nil {
		origin_type = transform.TypeOrigin
	}
	if origin_type == "bool" {
		return "Bool"
	} else if origin_type == "string" {
		return "String"
	} else if strings.HasPrefix(origin_type, "int") {
		return "Int64"
	} else if strings.HasPrefix(origin_type, "float") {
		return "Float64"
	}
	panic("unsupported null sql type: " + origin_type)
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
	"mssql_timestamp": { // TIMESTAMP (string, UTC)
		"string", `orm.MsSQLTimeParse(%v)`,
		"time.Time", `orm.MsSQLTimeFormat(%v)`,
	},
	"mssql_timeint": { // INT(11)
		"int64", "time.Unix(%v, 0)",
		"time.Time", "%v.Unix()",
	},
	"mssql_datetime": { // DATETIME (string, localtime)
		"string", "orm.MsSQLTimeParse(%v)",
		"time.Time", "orm.MsSQLTimeFormat(%v)",
	},
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
	"elastic_timestamp": { // TIMESTAMP (string, UTC)
		"string", `orm.TimeParse(%v)`,
		"time.Time", `orm.TimeFormat(%v)`,
	},
	"elastic_timeint": { // INT(11)
		"int64", "time.Unix(%v, 0)",
		"time.Time", "%v.Unix()",
	},
	"elastic_datetime": { // DATETIME (string, localtime)
		"string", "orm.TimeParseLocalTime(%v)",
		"time.Time", "orm.TimeToLocalTime(%v)",
	},
}

func (f *Field) IsNeedTransform() bool {
	return f.GetTransform() != nil
}

func (f *Field) GetTransform() *Transform {
	key := fmt.Sprintf("%v_%v", f.Obj.Db, f.Type)
	t, ok := transformMap[key]
	if !ok {
		return nil
	}
	return &t
}

func (f *Field) GetTransformValue(prefix string) string {
	t := f.GetTransform()
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
			tags["db"] = false
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
	if f.Validator != "" {
		tagstr = append(tagstr, fmt.Sprintf("validate:\"%s\"", f.Validator))
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
	f.ESIndex.DoIndex = f.Obj.ElasticIndexAll

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
		case "size":
			f.Size = v.(int)
		case "sqltype":
			f.sqlType = v.(string)
		case "sqlcolumn":
			f.sqlColumn = v.(string)
		case "comment":
			f.Comment = v.(string)
		case "validator":
			f.Validator = strings.ToLower(v.(string))
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

		case "es_do_index":
			f.ESIndex.DoIndex = v.(bool)

		case "es_do_analyze":
			f.ESIndex.DoAnalyze = v.(bool)

		case "es_analyzer":
			f.ESIndex.Analyzer = v.(string)

		case "es_date_format":
			f.ESIndex.DateFormat = v.(string)
		case "default":
			f.Default = v
		case "pb":
			f.extractPBFields(v)
		default:
			return errors.New("invalid field name: " + key)
		}
	}

	if f.Obj.DbContains("elastic") && f.ESIndex.ShouldIndex() {
		if err := f.ESIndex.SetType(f.Type); err != nil {
			return err
		}
	}

	//! single field primary adjust for redis ops
	if f.IsUnique() {
		index := NewIndex(f.Obj)
		index.FieldNames = []string{f.Name}
		f.Obj.uniques = append(f.Obj.uniques, index)
	}
	if f.IsIndex() {
		index := NewIndex(f.Obj)
		index.FieldNames = []string{f.Name}
		f.Obj.indexes = append(f.Obj.indexes, index)
	}
	if f.IsRange() {
		index := NewIndex(f.Obj)
		index.FieldNames = []string{f.Name}
		f.Obj.ranges = append(f.Obj.ranges, index)
	}
	return nil
}

func (f *Field) extractPBFields(v interface{}) {
	fds, ok := v.(map[interface{}]interface{})
	if !ok {
		log.Printf("pb plugin: field %s has defined pb but not defined any pb mapping", f.Name)
		return
	}
	if len(fds) != 1 {
		log.Printf("pb plugin: field %s has defined multi pb mapping,but only support one", f.Name)
		return
	}

	f.PBField = new(PBField)
	for n, t := range fds {
		f.PBField.Name = n.(string)
		f.PBField.Type = t.(string)
	}

	return
}

//! field SQL script functions
func (f *Field) SQLColumn(driver string) string {
	switch strings.ToLower(driver) {
	case "mysql":
		columns := make([]string, 0, 6)
		columns = append(columns, f.SQLName(driver))
		columns = append(columns, f.SQLType(driver))
		columns = append(columns, f.SQLNull(driver))
		if f.IsAutoIncrement() {
			columns = append(columns, "AUTO_INCREMENT")
		} else {
			columns = append(columns, f.SQLDefault(driver))
		}
		if f.Comment != "" {
			columns = append(columns, "COMMENT", "'"+f.Comment+"'")
		}
		return strings.Join(columns, " ")
	}
	return ""
}
func (f *Field) SQLName(driver string) string {
	switch strings.ToLower(driver) {
	case "mysql":
		return "`" + Camel2Name(f.Name) + "`"
	}
	return ""
}

func (f *Field) SQLType(driver string) string {
	if f.sqlType != "" {
		return strings.ToUpper(f.sqlType)
	}
	switch strings.ToLower(driver) {
	case "mysql":
		if f.IsNumber() {
			switch f.GetType() {
			case "bool":
				return "TINYINT(1) UNSIGNED"
			case "uint8":
				return "SMALLINT UNSIGNED"
			case "uint16":
				return "MEDIUMINT UNSIGNED"
			case "uint32":
				return "INT(11) UNSIGNED"
			case "uint64":
				return "BIGINT UNSIGNED"
			case "int8":
				return "SMALLINT"
			case "int16":
				return "MEDIUMINT"
			case "int32", "int":
				return "INT(11)"
			case "int64":
				return "BIGINT(20)"
			case "float32", "float64":
				return "FLOAT"
			case "time.Time", "*time.Time":
				return "BIGINT(20)"
			}
		}
		if f.IsString() {
			switch f.Type {
			case "datetime":
				return "DATETIME"
			case "timestamp", "timeint":
				return "TIMESTAMP"
			}
			if f.Size == 0 {
				return "VARCHAR(100)"
			}
			return fmt.Sprintf("VARCHAR(%d)", f.Size)
		}
		return f.GetType()
	}
	return ""
}

func (f *Field) SQLNull(driver string) string {
	switch strings.ToLower(driver) {
	case "mysql":
		if f.IsNullable() {
			return "NULL"
		}
		return "NOT NULL"
	}
	return ""
}

func (f *Field) SQLDefault(driver string) string {
	if f.IsNullable() {
		return ""
	}
	switch strings.ToLower(driver) {
	case "mysql":
		if f.IsTime() {
			if f.IsString() {
				return "DEFAULT CURRENT_TIMESTAMP"
			}
			if f.IsNumber() {
				return "DEFAULT '0'"
			}
		}

		if f.IsBool() {
			switch v, _ := f.Default.(bool); v {
			case true:
				return "DEFAULT '1'"
			default:
				return "DEFAULT '0'"
			}
		}

		if f.IsNumber() {
			return "DEFAULT '0'"
		}
		if f.IsString() {
			return "DEFAULT ''"
		}
		return ""
	}
	return ""
}

type Fields []*Field

func (fs Fields) GetFuncParam() string {
	var params []string
	for _, f := range fs {
		params = append(params, CamelName(f.Name)+" "+f.GetType())
	}
	return strings.Join(params, ", ")
}

func (fs Fields) GetObjectParam() string {
	var params []string
	for _, f := range fs {
		params = append(params, "obj."+f.Name)
	}
	return strings.Join(params, ", ")
}

func (fs Fields) GetConstructor() string {
	params := make([]string, 0, len(fs)+1)
	for _, f := range fs {
		params = append(params, f.Name+" : "+CamelName(f.Name))
	}
	params = append(params, "")
	return strings.Join(params, ",\n")
}
