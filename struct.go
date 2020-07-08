package main

import "fmt"

func main() {
	fmt.Println(&person{name: "Ann", age: 40})//& 前缀生成一个结构体指针。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp:=&s
	fmt.Println(sp)



}

type person struct {
	name string
	age int
}