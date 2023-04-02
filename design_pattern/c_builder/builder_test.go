package c_builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		name    string
		builder *ResourcePoolConfigBuilder
		want    *ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "名称为空",
			builder: &ResourcePoolConfigBuilder{
				name:     "",
				maxTotal: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "maxIdle小于minIdle",
			builder: &ResourcePoolConfigBuilder{
				name:     "test",
				maxTotal: 0,
				maxIdle:  10,
				minIdle:  20,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			builder: &ResourcePoolConfigBuilder{
				name:     "test",
				maxTotal: defaultMaxTotal,
				maxIdle:  defaultMaxIdle,
				minIdle:  defaultMinIdle,
			},
			want: &ResourcePoolConfig{
				name:     "test",
				maxTotal: defaultMaxTotal,
				maxIdle:  defaultMaxIdle,
				minIdle:  defaultMinIdle,
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.builder.Build()
			assert.Equal(t, test.wantErr, err != nil, "Build() error = %v, wantErr %v", err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
