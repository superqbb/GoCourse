package main

import (
	"fmt"
	"time"
)

func DelayPrint(from string) {
	for i := 1; i <= 4; i++ {
		//此处的sleep不会阻塞 HelloWorld
		time.Sleep(250 * time.Millisecond)
		fmt.Println(from, ":", i)
	}
}

func HelloWorld() {
	fmt.Println("Hello world goroutine")
}

func main() {
	//使用go关键字，开启goroutine
	go DelayPrint("direct")

	go DelayPrint("goroutine")

	go HelloWorld()

	//使用匿名函数开启goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//这里使用Scanln()监听键盘输入，否则main执行完后进程退出，所创建的协程会销毁
	fmt.Scanln()
	fmt.Println("done")
}