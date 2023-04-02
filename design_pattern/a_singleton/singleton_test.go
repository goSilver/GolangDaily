package a_singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstanceLazy(t *testing.T) {
	assert.Equal(t, GetInstanceLazy(), GetInstanceLazy())
}

func BenchmarkGetInstanceLazyParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstanceLazy() != GetInstanceLazy() {
				b.Errorf("test fail")
			}
		}
	})
}

func TestGetInstanceLazyWithLock(t *testing.T) {
	assert.Equal(t, GetInstanceLazyWithLock(), GetInstanceLazyWithLock())
}

func BenchmarkGetInstanceLazyWithLockParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstanceLazyWithLock() != GetInstanceLazyWithLock() {
				b.Errorf("test fail")
			}
		}
	})
}

func TestGetInstanceLazyWithDoubleCheck(t *testing.T) {
	assert.Equal(t, GetInstanceLazyWithDoubleCheck(), GetInstanceLazyWithDoubleCheck())
}

func BenchmarkGetInstanceLazyWithDoubleCheckParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstanceLazyWithDoubleCheck() != GetInstanceLazyWithDoubleCheck() {
				b.Errorf("test fail")
			}
		}
	})
}

func TestGetInstanceHungry(t *testing.T) {
	assert.Equal(t, GetInstanceHungry(), GetInstanceHungry())
}

func BenchmarkGetInstanceHungryParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstanceHungry() != GetInstanceHungry() {
				b.Errorf("test fail")
			}
		}
	})
}

func TestGetInstanceOnce(t *testing.T) {
	assert.Equal(t, GetInstanceOnce(), GetInstanceOnce())
}

func BenchmarkGetInstanceOnceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstanceOnce() != GetInstanceOnce() {
				b.Errorf("test fail")
			}
		}
	})
}
