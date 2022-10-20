package reflect

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name string
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
