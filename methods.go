package main

import "fmt"

/*
Go 支持为结构体类型定义方法(methods) 。
area 是一个拥有 *rect 类型接收器(receiver)的方法。
方法就是一个包含了接受者（receiver）的函数，
receiver可以是内置类型或者结构体类型的一个值或者是一个指针。
所有给定类型的方法属于该类型的方法集。
可以为值类型或者指针类型的接收者定义方法。
这是一个值类型接收者的例子。
func (r ReceiverType) funcName(parameters) (results)
当调用method时，会将receiver作为函数的第一个参数：
funcName(r, parameters);
*/
func main() {
	r:=rect{width: 10,height: 22}
	fmt.Println(r.area())
	fmt.Println(r.perim())
	rp:=&r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}

type rect struct {//定义新类
	width int
	height int
}
//method 语法 func (r ReceiverType) funcName(parameters) results type{}
func (r* rect)area()int{
	return r.width*r.height
}
func (r rect) perim() int {
	return 2*(r.width+r.height)

}