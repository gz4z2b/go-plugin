/*
 * @Author: p_hanxichen
 * @Date: 2023-12-25 16:44:36
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/slice/find_test.go
 * @Description: 查抄切片中指定值的位置
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slice

import "testing"

func TestFind(t *testing.T) {
	type args struct {
		slice []string
		val   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "找到",
			args: args{
				slice: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
				val:   "c",
			},
			want: 2,
		},
		{
			name: "没找到",
			args: args{
				slice: []string{"a", "b"},
				val:   "c",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.slice, tt.args.val); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
