package mutex

type List[T any] interface {
	Get(index int) T
	Set(index int, t T)
	DeleteAt(index int) T
	Append(t T)
}

// ArrayList 一个普通的、不安全的List
type ArrayList[T any] struct {
	values []T
}

func (a *ArrayList[T]) Get(index int) T {
	return a.values[index]
}

func (a *ArrayList[T]) Set(index int, t T) {
	if index >= len(a.values) || index < 0 {
		panic("index out of range")
	}
	a.values[index] = t
}
