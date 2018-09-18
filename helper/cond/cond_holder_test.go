// Package cond provides orm-liked sql builder
package cond

import (
	"testing"
)

func TestCond_BuildPlaceholderStr(t *testing.T) {

	tests := []struct {
		name       string
		cond       Cond
		wantHolder string
	}{
		{
			name:       "TestEmptyPlaceholder",
			cond:       NewIntCond([]int{}),
			wantHolder: "",
		},
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

func TestCond_BuildPlaceholderStrWithBrackets(t *testing.T) {

	tests := []struct {
		name       string
		cond       Cond
		wantHolder string
	}{
		{
			name:       "TestEmptyPlaceholderWithBrackets",
			cond:       NewIntCond([]int{}),
			wantHolder: "",
		},
		{
			name:       "TestIntPlaceholderWithBrackets",
			cond:       NewIntCond([]int{1, 2, 3}),
			wantHolder: "(?,?,?)",
		},
		{
			name:       "TestInt32PlaceholderWithBrackets",
			cond:       NewInt32Cond([]int32{1, 2, 3}),
			wantHolder: "(?,?,?)",
		},
		{
			name:       "TestInt64PlaceholderWithBrackets",
			cond:       NewInt64Cond([]int64{1, 2, 3}),
			wantHolder: "(?,?,?)",
		},
		{
			name:       "TestStrPlaceholderWithBrackets",
			cond:       NewStrCond([]string{"1", "2", "3"}),
			wantHolder: "(?,?,?)",
		},
		{
			name:       "TestDefaultMultiColumnsWithBrackets",
			cond:       NewDefaultMultiColumns([]string{"1:2", "2:3"}),
			wantHolder: "((?,?),(?,?))",
		},
		{
			name:       "TestCustomizedMultoColumnsWithBrackets",
			cond:       NewMultiColumns(NewC([]string{"1|2", "3|4"})),
			wantHolder: "((?,?),(?,?))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHolder := tt.cond.BuildPlaceholderStrWithBrackets(); gotHolder != tt.wantHolder {
				t.Errorf("Cond.BuildPlaceholderStrWithBrackets() = %v, want %v", gotHolder, tt.wantHolder)
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
