package main

import "fmt"

func main() {
	fmt.Println(plus(1,9))
	a, b := vals()
	fmt.Println(a,b)
	sum(1,2,3)
	num:=[]int{1,2,3,4}
	sum(num...)
	//如果你有一个含有多个值的 slice，
	//想把它们作为参数使用， 你需要这样调用 func(slice...)。

}
func plus(a,b int)int {
	return a+b
}
func vals()(int,int){
	return 0, 1
}
func sum(nums...int){
	fmt.Println(nums)
	total:=0
	for _, num:= range nums{
		total +=num
	}
	fmt.Println(total)
}