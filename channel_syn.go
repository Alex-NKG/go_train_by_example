package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 1)
	go worker(done)
	//程序将一直阻塞，直至收到 worker 使用通道发送的通知。

	<-done
}
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}