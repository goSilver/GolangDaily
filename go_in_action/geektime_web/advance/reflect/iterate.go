package reflect

import (
	"errors"
	"reflect"
)

// Iterate 迭代数组、切片、字符串
func Iterate(input any) ([]any, error) {
	val := reflect.ValueOf(input)
	typ := val.Type()
	kind := typ.Kind()
	if kind != reflect.Array && kind != reflect.Slice && kind != reflect.String {
		return nil, errors.New("非法类型")
	}
	res := make([]any, 0, val.Len())
	for i := 0; i < val.Len(); i++ {
		ele := val.Index(i)
		res = append(res, ele.Interface())
	}
	return res, nil
}

// IterateMapV1 迭代map
func IterateMapV1(input any) ([]any, []any, error) {
	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Map {
		return nil, nil, errors.New("非法类型")
	}
	keys := make([]any, 0, val.Len())
	values := make([]any, 0, val.Len())
	for _, k := range val.MapKeys() {
		keys = append(keys, k.Interface())
		v := val.MapIndex(k)
		values = append(values, v.Interface())
	}
	return keys, values, nil
}

func IterateMapV2(input any) ([]any, []any, error) {
	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Map {
		return nil, nil, errors.New("非法类型")
	}
	keys := make([]any, 0, val.Len())
	values := make([]any, 0, val.Len())
	itr := val.MapRange()
	for itr.Next() {
		keys = append(keys, itr.Key().Interface())
		values = append(values, itr.Value().Interface())
	}
	return keys, values, nil
}
