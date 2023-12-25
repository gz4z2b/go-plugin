/*
 * @Author: p_hanxichen
 * @Date: 2023-12-25 18:05:23
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/slice/union_test.go
 * @Description: 求并集
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slice

import (
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
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
			name: "有交集的",
			args: args{
				a: []string{"a", "b", "c"},
				b: []string{"c", "d", "e"},
			},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "无交集的",
			args: args{
				a: []string{"a", "b", "c"},
				b: []string{"f", "d", "e"},
			},
			want: []string{"a", "b", "c", "f", "d", "e"},
		},
		{
			name: "其中一个为空",
			args: args{
				a: []string{"a", "b", "c"},
				b: []string{},
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}
