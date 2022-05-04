package main

type user struct{}

type manager struct {
	user
}

func (user) toString() string {
	return "user"
}

func (m manager) toString() string {
	return m.user.toString() + ";manager"
}

func main() {
	var m manager
	// 匿名字段引入但同名方法遮蔽实现类型覆盖效果
	println(m.toString())
	println(m.user.toString())
}
