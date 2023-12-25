/*
 * @Author: p_hanxichen
 * @Date: 2023-12-25 18:05:23
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/slice/union.go
 * @Description: 求并集
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slice

// Union 求并集
func Union[T compareable](a, b []T) []T {
	var result []T
	for _, v := range a {
		if Find[T](result, v) == -1 {
			result = append(result, v)
		}
	}
	for _, v := range b {
		if Find[T](result, v) == -1 {
			result = append(result, v)
		}
	}
	return result
}
