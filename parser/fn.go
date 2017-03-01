package parser

import (
	"bytes"
	"io/ioutil"
	"strings"
	"unicode"

	yaml "gopkg.in/yaml.v2"
)

func ReadYaml(packageName string, yamlFile string) ([]*MetaObject, error) {
	var models map[string]map[string]interface{}

	data, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal([]byte(data), &models); err != nil {
		return nil, err
	}

	objs := []*MetaObject{}
	for name, model := range models {
		metaObj := NewMetaObject(packageName)
		if err := metaObj.Read(name, model); err != nil {
			return nil, err
		}

		objs = append(objs, metaObj)
	}
	return objs, nil
}

func toStringSlice(val []interface{}) (result []string) {
	result = make([]string, len(val))
	for i, v := range val {
		result[i] = v.(string)
	}
	return
}

func isUpperCase(c string) bool {
	return c == strings.ToUpper(c)
}

////////////////////////////////////////////////////////////////////////
func Camel2Name(s string) string {
	nameBuf := bytes.NewBuffer(nil)
	before := false
	for i := range s {
		n := rune(s[i]) // always ASCII?
		if unicode.IsUpper(n) {
			if !before && i > 0 {
				nameBuf.WriteRune('_')
			}
			n = unicode.ToLower(n)
			before = true
		} else {
			before = false
		}
		nameBuf.WriteRune(n)
	}
	return nameBuf.String()
}

func ToIds(bufName, typeName, name string) string {
	switch typeName {
	case "int":
		return "intToIds(" + bufName + "," + name + ")"
	case "int32":
		return "int32ToIds(" + bufName + "," + name + ")"
	case "bool":
		return "boolToIds(" + bufName + "," + name + ")"
	case "string":
		return "stringToIds(" + bufName + "," + name + ")"
	}
	return name
}
