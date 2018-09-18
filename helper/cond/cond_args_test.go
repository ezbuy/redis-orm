package cond

import (
	"reflect"
	"testing"
)

func TestCond_BuildArgs(t *testing.T) {

	tests := []struct {
		name     string
		cond     Cond
		wantArgs []interface{}
	}{
		{
			name:     "TestIntArgs",
			cond:     NewIntCond([]int{1, 2, 3}),
			wantArgs: []interface{}{1, 2, 3},
		},
		{
			name:     "TestInt64Args",
			cond:     NewInt64Cond([]int64{int64(1), int64(2)}),
			wantArgs: []interface{}{int64(1), int64(2)},
		},
		{
			name:     "TestInt32Args",
			cond:     NewInt32Cond([]int32{int32(1), int32(2)}),
			wantArgs: []interface{}{int32(1), int32(2)},
		},
		{
			name:     "TestStrArgs",
			cond:     NewStrCond([]string{"1", "2", "3"}),
			wantArgs: []interface{}{"1", "2", "3"},
		},
		{
			name:     "TestDefaultArgs",
			cond:     NewDefaultMultiColumns([]string{"1:2", "3:4"}),
			wantArgs: []interface{}{"1", "2", "3", "4"},
		},
		{
			name:     "TestCustomizedArgs",
			cond:     NewMultiColumns(NewC([]string{"1|2", "3|4"})),
			wantArgs: []interface{}{"1", "2", "3", "4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotArgs := tt.cond.BuildArgs()
			if ok := reflect.DeepEqual(gotArgs, tt.wantArgs); !ok {
				t.Errorf("Cond.BuildArgs() = %+v,want %+v", gotArgs, tt.wantArgs)
			}
		})
	}
}
