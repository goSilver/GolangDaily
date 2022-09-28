package main

import "fmt"

func main() {

	var n News = fakeNews{
		Name: "BBC",
	}
	n.Report()
}

type News struct {
	Name string
}

func (n News) Report() {
	fmt.Println("I am news:" + n.Name)
}

type fakeNews = News
