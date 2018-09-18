// Package cond provides orm-liked sql builder
package cond

import (
	"fmt"
	"testing"
)

func TestCond_BuildPlaceholderStr(t *testing.T) {

	tests := []struct {
		name       string
		cond       Cond
		wantHolder string
	}{
		{
			name:       "TestIntPlaceholder",
			cond:       NewIntCond([]int{1, 2, 3}),
			wantHolder: "?,?,?",
		},
		{
			name:       "TestInt32Placeholder",
			cond:       NewInt32Cond([]int32{1, 2, 3}),
			wantHolder: "?,?,?",
		},
		{
			name:       "TestInt64Placeholder",
			cond:       NewInt64Cond([]int64{1, 2, 3}),
			wantHolder: "?,?,?",
		},
		{
			name:       "TestStrPlaceholder",
			cond:       NewStrCond([]string{"1", "2", "3"}),
			wantHolder: "?,?,?",
		},
		{
			name:       "TestDefaultMultiColumns",
			cond:       NewDefaultMultiColumns([]string{"1:2", "2:3"}),
			wantHolder: "(?,?),(?,?)",
		},
		{
			name:       "TestCustomizedMultoColumns",
			cond:       NewMultiColumns(NewC([]string{"1|2", "3|4"})),
			wantHolder: "(?,?),(?,?)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHolder := tt.cond.BuildPlaceholderStr(); gotHolder != tt.wantHolder {
				t.Errorf("Cond.BuildPlaceholderStr() = %v, want %v", gotHolder, tt.wantHolder)
			}
		})
	}
}

type CustomizedMultiColumns struct {
	strs []string
}

func (c CustomizedMultiColumns) Key() string {
	return "|"
}

func NewC(strs []string) CustomizedMultiColumns {
	return CustomizedMultiColumns{
		strs: strs,
	}
}

func (c CustomizedMultiColumns) BuildMultiColumnsPlaceholderStr() string {
	return BuildMultiColumnsPlaceholderStrWithKey(c.strs, c.Key())
}

func (c CustomizedMultiColumns) BuildMultiColumnsArgs() []interface{} {
	return BuildMultiColumnsArgsWithKey(c.strs, c.Key())
}

func (c CustomizedMultiColumns) BuildMultiColumnsPlaceholderStrWithQuote() string {
	str := c.BuildMultiColumnsPlaceholderStr()
	if str == "" {
		return ""
	}

	return fmt.Sprintf("(%s)", str)
}
