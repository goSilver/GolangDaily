package demo

import (
	"errors"
	"reflect"
)

// IterateFuncs 输出方法信息，并执行调用
func IterateFuncs(val any) (map[string]*FuncInfo, error) {

	if val == nil {
		return nil, errors.New("不能为nil")
	}
	typ := reflect.TypeOf(val)
	if typ.Kind() != reflect.Struct && typ.Kind() != reflect.Ptr {
		return nil, errors.New("不支持的类型")
	}
	numMethod := typ.NumMethod()
	res := make(map[string]*FuncInfo, numMethod)
	for i := 0; i < numMethod; i++ {
		method := typ.Method(i)
		mt := method.Type
		// 入参
		numIn := mt.NumIn()
		in := make([]reflect.Type, 0, numIn)
		for j := 0; j < numIn; j++ {
			in = append(in, mt.In(j))
		}
		// 出参
		numOut := mt.NumOut()
		out := make([]reflect.Type, 0, numOut)
		for k := 0; k < numOut; k++ {
			out = append(out, mt.Out(k))
		}

		callRes := method.Func.Call([]reflect.Value{reflect.ValueOf(val), reflect.ValueOf("233")})
		retVals := make([]any, 0, len(callRes))
		for _, cr := range callRes {
			retVals = append(retVals, cr.Interface())
		}

		res[method.Name] = &FuncInfo{
			Name:   method.Name,
			In:     in,
			Out:    out,
			Result: retVals,
		}
	}

	return res, nil
}

type FuncInfo struct {
	Name   string
	In     []reflect.Type
	Out    []reflect.Type
	Result []any // 反射调用得到的结果
}
