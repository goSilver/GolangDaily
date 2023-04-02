package c_builder

import (
	"errors"
	"fmt"
)

const (
	defaultMaxTotal = 10
	defaultMaxIdle  = 9
	defaultMinIdle  = 1
)

type ResourcePoolConfig struct {
	name     string
	maxTotal int64
	maxIdle  int64
	minIdle  int64
}

type ResourcePoolConfigBuilder struct {
	name     string
	maxTotal int64
	maxIdle  int64
	minIdle  int64
}

func (r *ResourcePoolConfigBuilder) setName(name string) error {
	if len(name) == 0 {
		return errors.New("name为空")
	}
	r.name = name
	return nil
}

func (r *ResourcePoolConfigBuilder) setMaxTotal(maxTotal int64) error {
	if maxTotal < 0 {
		return errors.New("maxTotal小于零")
	}
	r.maxTotal = maxTotal
	return nil
}

func (r *ResourcePoolConfigBuilder) setMaxIdle(maxIdle int64) error {
	if maxIdle < 0 {
		return errors.New("maxIdle小于零")
	}
	r.maxIdle = maxIdle
	return nil
}

func (r *ResourcePoolConfigBuilder) setMinIdle(minIdle int64) error {
	if minIdle < 0 {
		return errors.New("minIdle小于零")
	}
	r.minIdle = minIdle
	return nil
}

// Build 建造方法
func (r *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if r.name == "" {
		return nil, fmt.Errorf("name can not be empty")
	}

	// 设置默认值
	if r.minIdle == 0 {
		r.minIdle = defaultMinIdle
	}

	if r.maxIdle == 0 {
		r.maxIdle = defaultMaxIdle
	}

	if r.maxTotal == 0 {
		r.maxTotal = defaultMaxTotal
	}

	if r.maxTotal < r.maxIdle {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", r.maxTotal, r.maxIdle)
	}

	if r.minIdle > r.maxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", r.maxIdle, r.minIdle)
	}

	return &ResourcePoolConfig{
		name:     r.name,
		maxTotal: r.maxTotal,
		maxIdle:  r.maxIdle,
		minIdle:  r.minIdle,
	}, nil
}
