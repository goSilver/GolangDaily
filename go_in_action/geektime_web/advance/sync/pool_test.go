package sync

import (
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	p := sync.Pool{
		New: func() interface{} {
			// 创建函数，sync.Pool会回调
			return nil
		},
	}

	obj := p.Get()
	// do something:在这里用取出来到对象

	// 用完再还回去
	p.Put(obj)
}
