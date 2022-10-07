package mutex

import "sync"

// PublicResource 你永远不知道你到用户拿它会干啥
// 用户用不用PublicResourceLock，完全无法确定
// 如果用这个resource，一定要用锁
var PublicResource interface{}
var PublicResourceLock sync.Mutex

// privateResource 要好一点，不过依赖用户来看注释，知道要用锁
// 很多库都是这么写的
var privateResource interface{}
var privateResourceLock sync.Mutex

// safeResource 很棒。所有期望对resource对操作都只能通过定义在safeResource上的方法来进行
type safeResource struct {
	resource interface{}
	lock     sync.Mutex
}

func (s *safeResource) DoSomething() {
	s.lock.Lock()
	defer s.lock.Unlock()
	// do something
}

/* 锁的伪代码
type Lock struct {
	state int
}

func (l *Lock) Lock() {
	i = 0
	for locked = CAS(UN_LOCK, LOCKED); !locked && i < 10 {
		i ++
	}

	if locked {
		return
	}

	// 将自己的线程加入阻塞队列
	enqueue()
}
*/
