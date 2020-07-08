package main

import "fmt"

/*
通道(channels) 是连接多个协程的管道。
你可以从一个协程将值发送到通道，
然后在另一个协程中接收。
默认发送和接收操作是阻塞的，直到发送方和接收方都就绪。
这个特性允许我们，不使用任何其它的同步操作

*/
func main() {
	message :=make(chan string,2)//创建
	go func() {message<-"sdd"}()//发送
	message<-"sddsss"
	msg:=<-message
	msg2:=<-message
	fmt.Println(msg)
	fmt.Println(msg2)
}
