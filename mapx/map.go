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

func Keys[KeyType comparable, ValueType any](m map[KeyType]ValueType) []KeyType {
	keys := make([]KeyType, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func Values[KeyType comparable, ValueType any](m map[KeyType]ValueType) []ValueType {
	values := make([]ValueType, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

func KeyValues[KeyType comparable, ValueType any](m map[KeyType]ValueType) ([]KeyType, []ValueType) {
	keys := make([]KeyType, 0, len(m))
	values := make([]ValueType, 0, len(m))

	for key, value := range m {
		keys = append(keys, key)
		values = append(values, value)
	}
	return keys, values
}
