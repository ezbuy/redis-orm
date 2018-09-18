package cond

import (
	"bytes"
	"strings"
)

// MultiColumnsCond handle the multi columns conditions
type MultiColumnsCond interface {
	Key() string
	BuildMultiColumnsPlaceholderStr() string
	BuildMultiColumnsArgs() []interface{}
}

// DefaultMultiColumnsCond impl the default MultiColumnsCond
type DefaultMultiColumnsCond struct {
	combinedStr []string
}

// Key set the multi columns split key
func (dm DefaultMultiColumnsCond) Key() string {
	return ":"
}

// BuildMultiColumnsPlaceholderStr builds ?,?,?... placeholders
func (dm DefaultMultiColumnsCond) BuildMultiColumnsPlaceholderStr() string {
	return BuildMultiColumnsPlaceholderStrWithKey(dm.combinedStr, dm.Key())
}

// BuildMultiColumnsPlaceholderStrWithKey build multi-columns with a split key
func BuildMultiColumnsPlaceholderStrWithKey(strs []string, key string) string {
	if len(strs) == 0 {
		return ""
	}
	var holderStrSlice []string
	for _, rawparam := range strs {
		buffer := bytes.NewBufferString("")
		buffer.WriteByte('(')
		splittedParam := strings.Split(rawparam, key)
		var tmpHolders []string
		for range splittedParam {
			tmpHolders = append(tmpHolders, "?")
		}
		buffer.WriteString(strings.Join(tmpHolders, ","))
		buffer.WriteByte(')')
		holderStrSlice = append(holderStrSlice, buffer.String())
	}
	holderStr := strings.Join(holderStrSlice, ",")
	return holderStr
}

// BuildMultiColumnsArgs builds interface{} args
func (dm DefaultMultiColumnsCond) BuildMultiColumnsArgs() []interface{} {
	return BuildMultiColumnsArgsWithKey(dm.combinedStr, dm.Key())
}

// BuildMultiColumnsArgsWithKey builds interface{} args with split key
func BuildMultiColumnsArgsWithKey(strs []string, key string) []interface{} {
	if len(strs) == 0 {
		return []interface{}{}
	}
	var args []interface{}
	for _, s := range strs {
		splitterArgs := strings.Split(s, key)
		for _, ss := range splitterArgs {
			args = append(args, ss)
		}
	}
	return args
}

// NewMultiColumnsCond new splitter
func NewMultiColumnsCond(splittedStr []string) DefaultMultiColumnsCond {
	cond := DefaultMultiColumnsCond{}
	cond.combinedStr = splittedStr
	return cond
}
