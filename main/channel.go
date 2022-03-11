package main

import "fmt"

//关闭chan后,再从中读取消息
func main() {
	c := make(chan int, 2)
	c <- 1
	close(c)
	for {
		val, ok := <-c
		if !ok {
			fmt.Println("收到关闭消息1")
			break
		}
		val++
	}
	fmt.Println("正常执行:创建channel,发送消息,关闭,接收消息")

	c = make(chan int)
	close(c)
	for {
		val, ok := <-c
		if !ok {
			fmt.Println("收到关闭消息2")
			break
		}
		val++
	}
	for {
		val, ok := <-c
		if !ok {
			fmt.Println("收到关闭消息3")
			break
		}
		val++
	}
	fmt.Println("正常执行:创建channel,关闭,接收消息")
}
