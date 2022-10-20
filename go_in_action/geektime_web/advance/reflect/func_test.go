package reflect

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIterateFuncs(t *testing.T) {
	type args struct {
		val any
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*FuncInfo
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name:    "nil",
			wantErr: errors.New("不能为nil"),
		},
		{
			name: "basic type",
			args: args{
				val: 123,
			},
			wantErr: errors.New("不支持的类型"),
		},
		{
			name: "结构体",
			args: args{
				val: &Order{
					OrderNum: "123",
				},
			},
			want: map[string]*FuncInfo{
				"SetOrderNum": {
					Name:   "SetOrderNum",
					In:     []reflect.Type{reflect.TypeOf(&Order{}), reflect.TypeOf("")},
					Out:    []reflect.Type{reflect.TypeOf("")},
					Result: []any{"233"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IterateFuncs(tt.args.val)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

type Order struct {
	OrderNum string
}

func (o *Order) SetOrderNum(orderNum string) string {
	o.OrderNum = orderNum
	return orderNum
}
