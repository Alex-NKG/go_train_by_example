package main

import "fmt"

/*
闭包：
让一个变量常驻内存
让变量不污染全局
闭包是指有权访问另一个函数作用域中的变量的函数
创建方法：在一个函数内创建另一个函数，
通个另一个函数来访问这个函数的局部变量
*/
func main() {
	nexInt:=intSeq()
	fmt.Println(nexInt())
	fmt.Println(nexInt())
	fmt.Println(nexInt())
	newInt:=intSeq()
	fmt.Println(newInt())
}
func intSeq() func() int{//返回类型function int
	i:=0
	return func() int{//匿名函数
		i++
		return i
	}
}