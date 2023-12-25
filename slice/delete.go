/*
 * @Author: p_hanxichen
 * @Date: 2023-12-25 16:34:40
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/slice/delete.go
 * @Description: 删除指定位置的元素
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slice

// Delete 删除指定位置的元素
func Delete[T number](slice []T, i int) ([]T, error) {
	if i < 0 || i >= len(slice) {
		return slice, ErrIndexOutOfRange
	}
	return append(slice[:i], slice[i+1:]...), nil
}

// DeleteMulti 删除多个指定位置的元素
func DeleteMulti[T number](slice []T, idxs []int) ([]T, error) {
	if len(idxs) == 0 {
		return slice, nil
	}
	if len(idxs) == 1 {
		return Delete[T](slice, idxs[0])
	}
	for _, i := range idxs {
		if i < 0 || i >= len(slice) {
			return slice, ErrIndexOutOfRange
		}
	}

	idxs = Sort[int](idxs, true)

	res := slice[:idxs[0]]
	for i := 1; i <= len(idxs)-1; i++ {
		res = append(res, slice[idxs[i-1]+1:idxs[i]]...)
	}
	res = append(res, slice[idxs[len(idxs)-1]+1:]...)

	return res, nil
}
