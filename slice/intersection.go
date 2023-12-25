/*
 * @Author: p_hanxichen
 * @Date: 2023-12-25 16:50:16
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/slice/intersection.go
 * @Description: 求交集
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slice

// Intersection 求交集
func Intersection[T compareable](a, b []T) []T {
	res := []T{}
	for _, v := range a {
		if Find(b, v) != -1 {
			res = append(res, v)
		}
	}
	return res
}
