package utils

import (
	"strings"
)

type MapStrFunc func(v string) interface{}

/* 对数组的每个元素进行处理 返回一个新的数组 */
func MapStrArr(arr []string, cb MapStrFunc) interface{} {
	newArr := []interface{}{}
	for _, v := range arr {
		newArr = append(newArr, cb(v))
	}
	return newArr
}

func ParseArr(str string, sep string, cb MapStrFunc) interface{} {
	str = strings.Trim(str, " ")
	str = strings.Trim(str, sep)
	return MapStrArr(strings.Split(str, sep), cb)
}

func ParseArrArr(str string, sep1 string, sep2 string, cb MapStrFunc) interface{} {
	result := []interface{}{}
	str = strings.Trim(str, " ")
	str = strings.Trim(str, sep1)
	var arr = strings.Split(str, sep1)
	for _, v := range arr {
		result = append(result, ParseArr(v, sep2, cb))
	}
	return result
}
