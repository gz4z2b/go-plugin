/*
 * @Author: p_hanxichen
 * @Date: 2023-12-26 10:06:24
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/mapx/map_test.go
 * @Description: map辅助方法
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package mapx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "不为空的",
			args: args{
				m: map[string]int{
					"1": 1,
					"2": 2,
					"3": 3,
				},
			},
			want: []string{"1", "2", "3"},
		},
		{
			name: "为空的",
			args: args{
				m: map[string]int{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Keys[string, int](tt.args.m)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestValues(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "不为空的",
			args: args{
				m: map[string]int{
					"1": 1,
					"2": 2,
					"3": 3,
				},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "为空的",
			args: args{
				m: map[string]int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Values[string, int](tt.args.m)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestKeyValues(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []int
	}{
		{
			name: "不为空的",
			args: args{
				m: map[string]int{
					"1": 1,
					"2": 2,
					"3": 3,
				},
			},
			want:  []string{"1", "2", "3"},
			want1: []int{1, 2, 3},
		},
		{
			name: "为空的",
			args: args{
				m: map[string]int{},
			},
			want:  []string{},
			want1: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := KeyValues[string, int](tt.args.m)

			assert.ElementsMatch(t, tt.want, got)
			assert.ElementsMatch(t, tt.want1, got1)
		})
	}
}

func TestKSortValues(t *testing.T) {
	type args struct {
		m   map[string]int
		asc bool
	}
	tests := []struct {
		name   string
		args   args
		keys   []string
		values []int
	}{
		{
			name: "不为空的",
			args: args{
				m: map[string]int{
					"a": 1,
					"c": 2,
					"b": 3,
				},
				asc: true,
			},
			keys:   []string{"a", "b", "c"},
			values: []int{1, 3, 2},
		},
		{
			name: "为空的",
			args: args{
				m:   map[string]int{},
				asc: true,
			},
			keys:   []string{},
			values: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys, values := KSortValues[string, int](tt.args.m, tt.args.asc)
			assert.Equal(t, tt.keys, keys)
			assert.Equal(t, tt.values, values)
		})
	}
}
