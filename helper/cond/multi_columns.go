package cond

import (
	"bytes"
	"fmt"
	"strings"
)

// MultiColumnsCond handle the multi columns conditions
type MultiColumnsCond interface {
	SetSplitKey(key string)
	SetMultiColumns(strs []string)
	BuildMultiColumnsPlaceholderStr() string
	BuildMultiColumnsArgs() []interface{}
	BuildMultiColumnsPlaceholderStrWithQuote() string
}

// DefaultMultiColumnsCond impl the default MultiColumnsCond
type DefaultMultiColumnsCond struct {
	key         string
	combinedStr []string
}

// SetSplitKey set the multi columns split key
func (dm DefaultMultiColumnsCond) SetSplitKey(key string) {
	dm.key = key
}

// SetMultiColumns set the combineStr
func (dm DefaultMultiColumnsCond) SetMultiColumns(strs []string) {
	dm.combinedStr = strs
}

// BuildMultiColumnsPlaceholderStr builds ?,?,?... placeholders
func (dm DefaultMultiColumnsCond) BuildMultiColumnsPlaceholderStr() string {
	if len(dm.combinedStr) == 0 {
		return ""
	}
	var holderStrSlice []string
	for _, rawparam := range dm.combinedStr {
		buffer := bytes.NewBufferString("")
		buffer.WriteByte('(')
		splittedParam := strings.Split(rawparam, dm.key)
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

// BuildMultiColumnsPlaceholderStrWithQuote builds (?,?,?...) placeholders
func (dm DefaultMultiColumnsCond) BuildMultiColumnsPlaceholderStrWithQuote() string {
	str := dm.BuildMultiColumnsPlaceholderStr()
	if str == "" {
		return ""
	}

	return fmt.Sprintf("(%s)", str)
}

// BuildMultiColumnsArgs builds interface{} args
func (dm DefaultMultiColumnsCond) BuildMultiColumnsArgs() []interface{} {
	if len(dm.combinedStr) == 0 {
		return []interface{}{}
	}
	var args []interface{}
	for _, s := range dm.combinedStr {
		splitterArgs := strings.Split(s, dm.key)
		for _, ss := range splitterArgs {
			args = append(args, ss)
		}
	}

	return args

}

// NewMultiColumnsCond new splitter
func NewMultiColumnsCond(key string, splittedStr []string) DefaultMultiColumnsCond {
	cond := DefaultMultiColumnsCond{}
	cond.SetMultiColumns(splittedStr)
	cond.SetSplitKey(key)
	return cond
}
