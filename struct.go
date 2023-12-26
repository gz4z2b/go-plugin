/*
 * @Author: p_hanxichen
 * @Date: 2023-12-25 15:20:20
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/struct.go
 * @Description: 结构定义
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package goplugin

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Compareable interface {
	~string | Number
}
