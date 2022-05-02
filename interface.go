package main

type tester interface {
	test()
	string() string
}

type data struct{}

func (*data) test() {}

func (data) string() string { return "" }

func main() {
	var d data

	// 错误：data dose not implement tester
	//var t tester = d

	var t tester = &d
	t.test()
	println(t.string())
}
