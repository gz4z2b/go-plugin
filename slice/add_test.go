/*
 * @Author: p_hanxichen
 * @Date: 2023-12-25 15:18:46
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/slice/add_test.go
 * @Description: 切片指定位置插入元素操作
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slice

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAdd(t *testing.T) {
	type args struct {
		slice []int
		idx   int
		val   int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "插入中间",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idx:   2,
				val:   6,
			},
			want:    []int{1, 2, 6, 3, 4, 5},
			wantErr: nil,
		},
		{
			name: "插入末尾",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idx:   5,
				val:   6,
			},
			want:    []int{1, 2, 3, 4, 5, 6},
			wantErr: nil,
		},
		{
			name: "插入开头",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idx:   0,
				val:   6,
			},
			want:    []int{6, 1, 2, 3, 4, 5},
			wantErr: nil,
		},
		{
			name: "idx越界(太小)",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idx:   -1,
				val:   6,
			},
			want:    nil,
			wantErr: ErrIndexOutOfRange,
		},
		{
			name: "idx越界(太大)",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				idx:   6,
				val:   6,
			},
			want:    nil,
			wantErr: ErrIndexOutOfRange,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add[int](tt.args.slice, tt.args.idx, tt.args.val)
			assert.Equal(t, err, tt.wantErr)
			assert.Equal(t, got, tt.want)

		})
	}
}
