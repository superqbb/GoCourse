package main

import "time"

func SetNonBuffChan(x chan int){
	time.Sleep(5*time.Second)
	x <- 1
}

func SetBuffChan(y chan int){
	for i:=0;i<10;i++{
		time.Sleep(1*time.Second)
		y<-i
	}
	close(y)
}


func main() {
	x := make(chan int)
	defer close(x)

	y := make(chan int,10)

	go SetNonBuffChan(x)
	go SetBuffChan(y)

	//开启两个goroutine同时读取channel x,y
	go func(){
		//range channel 可以直接取到 channel 中的值。一旦 channel 关闭，循环自动结束
		for ch := range y {
			println("y=",ch)
		}
	}()

	go func() {
		println("x=",<-x) //此处读取x时会阻塞
	}()

	select{}
}
