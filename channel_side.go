package main

import "fmt"

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
func ping(pings chan<- string, msg string) {//只写
	pings <- msg
}
//pings 仅用于接收数据（只读），pongs 仅用于发送数据（只写）
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}