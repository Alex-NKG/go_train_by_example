package main

import "fmt"

func main() {
	m:=make (map[string]int)
	m["K1"]=7
	m["K2"]=13
	v1 :=m["K1"]
	fmt.Println(v1)
	/*当从一个 map 中取值时，
	还有可以选择是否接收的第二个返回值，
	该值表明了 map 中是否存在这个键。
	这可以用来消除 键不存在 和 键的值为零值 产生的歧义，
	例如 0 和 ""。这里我们不需要值，
	所以用 空白标识符(blank identifier) _ 将其忽略。
	 */
	_,prs :=m["K2"]
	fmt.Println(prs)
	n:=map[string]int{"x":1}
	fmt.Println(n)

}
