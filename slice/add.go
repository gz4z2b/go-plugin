/*
 * @Author: p_hanxichen
 * @Date: 2023-12-25 15:18:46
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/slice/add.go
 * @Description: 切片指定位置插入元素操作
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slice

import goplugin "github.com/gz4z2b/go-plugin"

// Add 切片指定位置插入元素操作
func Add[T goplugin.Number](slice []T, idx int, val T) ([]T, error) {
	if idx > len(slice) {
		return nil, ErrIndexOutOfRange
	}
	if idx < 0 {
		return nil, ErrIndexOutOfRange
	}
	var zeroVal T
	slice = append(slice, zeroVal)

	for i := len(slice) - 1; i > idx; i-- {
		slice[i] = slice[i-1]
	}
	slice[idx] = val

	return slice, nil
}
