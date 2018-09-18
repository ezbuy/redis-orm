// Package in provides `IN` cond help functions
package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	testSrc := []interface{}{1, 2, 3, 4}
	type args struct {
		src      []interface{}
		splittor Splittor
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]interface{}
	}{
		{
			name: "TestZeroSize",
			args: args{
				src:      testSrc,
				splittor: NewSplittor(0),
			},
			wantRes: [][]interface{}(nil),
		},
		{
			name: "TestLowerThanSize",
			args: args{
				src:      testSrc,
				splittor: NewSplittor(1),
			},
			wantRes: [][]interface{}{
				[]interface{}{1},
				[]interface{}{2},
				[]interface{}{3},
				[]interface{}{4},
			},
		},
		{
			name: "TestEqualSize",
			args: args{
				src:      testSrc,
				splittor: NewSplittor(4),
			},
			wantRes: [][]interface{}{
				[]interface{}{1, 2, 3, 4},
			},
		},
		{
			name: "TestMoreThanSize",
			args: args{
				src:      testSrc,
				splittor: NewSplittor(5),
			},
			wantRes: [][]interface{}{
				[]interface{}{1, 2, 3, 4},
			},
		},
		{
			name: "TestDoubleEqualSize",
			args: args{
				src:      testSrc,
				splittor: NewSplittor(2),
			},
			wantRes: [][]interface{}{
				[]interface{}{1, 2},
				[]interface{}{3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Split(tt.args.src, tt.args.splittor); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Split() = %#v, want %#v", gotRes, tt.wantRes)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split([]interface{}{1, 2, 3, 4}, NewSplittor(2))
	}
}
