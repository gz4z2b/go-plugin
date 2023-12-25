/*
 * @Author: p_hanxichen
 * @Date: 2023-12-25 16:50:16
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/slice/intersection_test.go
 * @Description: 求交集
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slice

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIntersection(t *testing.T) {
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
			name: "有交集",
			args: args{
				a: []string{"a", "b", "c"},
				b: []string{"b", "c", "d"},
			},
			want: []string{"b", "c"},
		},
		{
			name: "无交集",
			args: args{
				a: []string{"a", "b", "c"},
				b: []string{"d", "e", "f"},
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersection[string](tt.args.a, tt.args.b)
			assert.Equal(t, got, tt.want)
		})
	}
}
