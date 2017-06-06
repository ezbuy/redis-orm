package parser

import "fmt"

var SupportedESFieldTypes = map[string]string{
	"bool":      "boolean",
	"int":       "integer",
	"int8":      "byte",
	"int16":     "short",
	"int32":     "integer",
	"int64":     "long",
	"float32":   "float",
	"float64":   "double",
	"string":    "string",
	"datetime":  "date",
	"timestamp": "date",
	"timeint":   "long",
}

var ESAnalyzableFields = map[string]bool{
	"string": true,
}

type TplESIndexMappingField struct {
	Field string
	Value string
}

type ESIndex struct {
	Type       string
	DoIndex    bool
	DateFormat string
	DoAnalyze  bool
	Analyzer   string
}

func (e *ESIndex) SetType(t string) error {
	esType, ok := SupportedESFieldTypes[t]
	if !ok {
		return fmt.Errorf("invalid elastic type " + t)
	}

	e.Type = esType
	return nil
}

func (e *ESIndex) ShouldIndex() bool {
	return e.DoIndex || e.ShouldAnalyze() || e.DateFormat != ""
}

func (e *ESIndex) ShouldAnalyze() bool {
	return e.Analyzer != "" || e.DoAnalyze
}

func (e *ESIndex) IndexType() string {
	if e.ShouldAnalyze() {
		return "analyzed"
	}

	return "not_analyzed"
}

func (e *ESIndex) TplMappingSettings() []TplESIndexMappingField {
	res := make([]TplESIndexMappingField, 0, 3)

	res = append(res, TplESIndexMappingField{
		"\"type\"",
		fmt.Sprintf("%q", e.Type),
	})

	switch e.Type {
	case "string":
		res = append(res, TplESIndexMappingField{
			"\"index\"",
			fmt.Sprintf("%q", e.IndexType()),
		})

		if e.Analyzer != "" {
			res = append(res, TplESIndexMappingField{
				"\"analyzer\"",
				fmt.Sprintf("%q", e.Analyzer),
			})
		}

	case "date":
		if e.DateFormat != "" {
			res = append(res, TplESIndexMappingField{
				"\"format\"",
				fmt.Sprintf("%q", e.DateFormat),
			})
		}
	}

	return res
}
