/*
 * @Author: p_hanxichen
 * @Date: 2023-12-26 09:33:40
 * @LastEditors: p_hanxichen
 * @FilePath: /basic/Users/chenhanxi/dev/go/src/go-plugin/slice/diff.go
 * @Description: 求差集
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package slice

import goplugin "github.com/gz4z2b/go-plugin"

// Diff 求差集
func Diff[T goplugin.Compareable](src, dst []T) []T {
	diff := []T{}
	for _, v := range src {
		if Find(dst, v) == -1 {
			diff = append(diff, v)
		}
	}
	return diff
}
