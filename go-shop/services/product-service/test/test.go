package main

import "fmt"

type student struct {
	Name string
	Age  int
}

type parent struct {
	Name string
	son  student
}

func Newpaarent(name string, son *student) *parent {
	return &parent{
		Name: name,
		son:  *son,
	}
}

func main() {
	s := student{
		Name: "张三",
		Age:  18,
	}
	p := Newpaarent("王五", &s)
	println(p.Name)
	println(s.Name)
	fmt.Println(&s.Age)
}
