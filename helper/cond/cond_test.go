// Package cond provides orm-liked sql builder
package cond

import "testing"

func TestIsCondOverMaxQuestions(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"over", args{1 << 12}, true},
		{"eq", args{1 << 11}, false},
		{"below", args{1 << 10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCondOverMaxQuestions(tt.args.count); got != tt.want {
				t.Errorf("IsCondOverMaxQuestions() = %v, want %v", got, tt.want)
			}
		})
	}
}
