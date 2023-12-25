/*
 * @Author: p_hanxichen
 * @Date: 2023-12-25 17:08:00
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/slice/sort_test.go
 * @Description: 排序
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slice

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args[T any] struct {
		slice []T
		asc   bool
	}
	type cases[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []cases[string]{
		{
			name: "正序",
			args: args[string]{
				slice: []string{"c", "d", "a", "A"},
				asc:   true,
			},
			want: []string{"A", "a", "c", "d"},
		},
		{
			name: "倒序",
			args: args[string]{
				slice: []string{"c", "d", "a", "A"},
				asc:   false,
			},
			want: []string{"d", "c", "a", "A"},
		},
		{
			name: "空切片",
			args: args[string]{
				slice: []string{},
				asc:   true,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sort[string](tt.args.slice, tt.args.asc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}

	casesInt := []cases[int]{
		{
			name: "数字",
			args: args[int]{
				slice: []int{1, 2, 6, 4, 5},
				asc:   true,
			},
			want: []int{1, 2, 4, 5, 6},
		},
	}
	for _, tt := range casesInt {
		t.Run(tt.name, func(t *testing.T) {
			got := Sort[int](tt.args.slice, tt.args.asc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
