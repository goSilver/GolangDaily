package reflect

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name string
	age  int
}

func TestIterateFields(t *testing.T) {
	u1 := &User{
		Name: "少华",
	}
	u2 := &u1

	tests := []struct {
		name    string         // 测试用例名称
		val     any            // 入参
		wantRes map[string]any // 期望的结果
		wantErr error          // 期望的error
	}{
		{
			name:    "nil",
			val:     nil,
			wantErr: errors.New("不能为 nil"),
		},
		{
			name:    "user",
			val:     User{Name: "Tom"},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "Tom",
			},
		},
		{
			name:    "pointer",
			val:     &User{Name: "Jerry"},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "Jerry",
			},
		},
		{
			name:    "multiple pointer",
			val:     u2,
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "少华",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := iterateFields(tt.val)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantRes, res)
		})
	}
}

func TestSetField(t *testing.T) {
	testCases := []struct {
		name string

		field  string
		entity any
		newVal any

		wantErr error
	}{
		{
			name:    "struct",
			entity:  User{},
			field:   "Name",
			wantErr: errors.New("非法类型"),
		},
		{
			name:    "private field",
			entity:  &User{},
			field:   "age",
			wantErr: errors.New("不可修改字段"),
		},
		{
			name:    "invalid field",
			entity:  &User{},
			field:   "invalid_field",
			wantErr: errors.New("字段不存在"),
		},
		{
			name: "pass",
			entity: &User{
				Name: "",
			},
			field:  "Name",
			newVal: "Tom",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := SetField(tc.entity, tc.field, tc.newVal)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
