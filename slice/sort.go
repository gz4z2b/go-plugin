/*
 * @Author: p_hanxichen
 * @Date: 2023-12-25 17:08:00
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/slice/sort.go
 * @Description: 排序
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slice

import goplugin "github.com/gz4z2b/go-plugin"

// 排序
func Sort[T goplugin.Compareable](slice []T, asc bool) []T {
	if len(slice) == 0 {
		return slice
	}

	if asc {
		for i := 0; i < len(slice)-1; i++ {
			for j := 0; j < len(slice)-i-1; j++ {
				if slice[j] > slice[j+1] {
					slice[j], slice[j+1] = slice[j+1], slice[j]
				}
			}
		}
	} else {
		for i := 0; i < len(slice)-1; i++ {
			for j := 0; j < len(slice)-i-1; j++ {
				if slice[j] < slice[j+1] {
					slice[j], slice[j+1] = slice[j+1], slice[j]
				}
			}
		}
	}
	return slice
}
