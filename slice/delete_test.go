/*
 * @Author: p_hanxichen
 * @Date: 2023-12-25 16:34:40
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/slice/delete_test.go
 * @Description: 删除指定位置的元素
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slice

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDelete(t *testing.T) {
	type args struct {
		slice []int
		i     int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr error
	}{
		{
			name: "删除中间的",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				i:     2,
			},
			want:    []int{1, 2, 4, 5},
			wantErr: nil,
		},
		{
			name: "删除最后一个",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				i:     4,
			},
			want:    []int{1, 2, 3, 4},
			wantErr: nil,
		},
		{
			name: "删除第一个",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				i:     0,
			},
			want:    []int{2, 3, 4, 5},
			wantErr: nil,
		},
		{
			name: "idx太小",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				i:     -1,
			},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: ErrIndexOutOfRange,
		},
		{
			name: "idx太大",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				i:     6,
			},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: ErrIndexOutOfRange,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Delete[int](tt.args.slice, tt.args.i)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestDeleteMulti(t *testing.T) {
	type args struct {
		slice []int
		idxs  []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr error
	}{
		{
			name: "删除中间",
			args: args{
				slice: []int{1, 2, 3, 4, 5, 6},
				idxs:  []int{1, 4, 3},
			},
			want:    []int{1, 3, 6},
			wantErr: nil,
		},
		{
			name: "删除头, 尾",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idxs:  []int{4, 0},
			},
			want:    []int{2, 3, 4},
			wantErr: nil,
		},
		{
			name: "删除头, 尾, 中间",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idxs:  []int{0, 3, 4},
			},
			want:    []int{2, 3},
			wantErr: nil,
		},
		{
			name: "都不删",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idxs:  []int{},
			},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: nil,
		},
		{
			name: "删除中间(一个)",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idxs:  []int{0},
			},
			want:    []int{2, 3, 4, 5},
			wantErr: nil,
		},
		{
			name: "越界(多个)",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idxs:  []int{-1, 6},
			},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: ErrIndexOutOfRange,
		},
		{
			name: "越界太小(一个)",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idxs:  []int{-1},
			},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: ErrIndexOutOfRange,
		},
		{
			name: "越界太大(一个)",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idxs:  []int{6},
			},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: ErrIndexOutOfRange,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteMulti[int](tt.args.slice, tt.args.idxs)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
