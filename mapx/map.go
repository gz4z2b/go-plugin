/*
 * @Author: p_hanxichen
 * @Date: 2023-12-26 10:06:24
 * @LastEditors: p_hanxichen
 * @FilePath: /go-plugin/mapx/map.go
 * @Description: map辅助方法
 *
 * Copyright (c) 2023 by gdtengnan, All Rights Reserved.
 */
package mapx

import (
	goplugin "github.com/gz4z2b/go-plugin"
	"github.com/gz4z2b/go-plugin/slice"
)

// Keys 返回map的所有key
func Keys[KeyType comparable, ValueType any](m map[KeyType]ValueType) []KeyType {
	keys := make([]KeyType, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// Values 返回map的所有value
func Values[KeyType comparable, ValueType any](m map[KeyType]ValueType) []ValueType {
	values := make([]ValueType, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

// KeyValues 返回map的所有key和value
func KeyValues[KeyType comparable, ValueType any](m map[KeyType]ValueType) ([]KeyType, []ValueType) {
	keys := make([]KeyType, 0, len(m))
	values := make([]ValueType, 0, len(m))

	for key, value := range m {
		keys = append(keys, key)
		values = append(values, value)
	}
	return keys, values
}

// KSortValues 按key排序,返回所有按照此顺序的key和value
func KSortValues[KeyType goplugin.Compareable, ValueType any](m map[KeyType]ValueType, asc bool) ([]KeyType, []ValueType) {
	keys := Keys(m)
	keys = slice.Sort[KeyType](keys, asc)

	values := make([]ValueType, 0, len(m))
	for _, value := range keys {
		values = append(values, m[value])
	}
	return keys, values
}
