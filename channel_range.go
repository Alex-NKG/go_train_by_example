package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}
//range 迭代从 queue 中得到每个值。
//因为我们在前面 close 了这个通道，
//所以，这个迭代会在接收完 2 个值之后结束。