/*
 * @Author: p_hanxichen
 * @Date: 2023-12-25 16:44:36
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/slice/find.go
 * @Description: 查抄切片中指定值的位置
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slice

// Find 查抄切片中指定值的首次出现的位置, 若未找到返回-1
func Find[T compareable](slice []T, val T) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}
