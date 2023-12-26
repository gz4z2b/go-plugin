/*
 * @Author: p_hanxichen
 * @Date: 2023-12-26 09:33:40
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/slice/diff_test.go
 * @Description: 求差集
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slice

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "结果为空",
			args: args{
				a: []string{"a", "b", "c"},
				b: []string{"a", "b", "c"},
			},
			want: []string{},
		},
		{
			name: "a不为空，b为空",
			args: args{
				a: []string{"a", "b", "c"},
				b: []string{},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "a为空，b不为空",
			args: args{
				a: []string{},
				b: []string{"a", "b", "c"},
			},
			want: []string{},
		},
		{
			name: "结果不为空",
			args: args{
				a: []string{"a", "b", "c"},
				b: []string{"d", "b", "f"},
			},
			want: []string{"a", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff[string](tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
