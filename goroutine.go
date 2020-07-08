package main

import (
	"fmt"
	"time"
)

/*
协程(goroutine) 是轻量级的执行线程。
当我们运行这个程序时，首先会看到阻塞式调用的输出，
然后是两个协程的交替输出。
这种交替的情况表示 Go runtime 是以并发的方式运行协程的。
(不是很明白）
*/
func main() {
	f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

