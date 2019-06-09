package main

import (
	"fmt"
	"time"
)

func main(){
	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(4*time.Second)
		close(c)
	}()

	go func() {
		time.Sleep(3*time.Second)
		ch1 <- 3
	}()

	go func() {
		time.Sleep(3*time.Second)
		ch2 <- 5
	}()

	fmt.Println("Blocking on read...")
	select {
		//若将ch的gorountine休眠时间改为5s，将会走case <- c
		case <- c:
			fmt.Printf("Unblocked %v later.\n", time.Since(start))
		case <- ch1:
			fmt.Printf("ch1 case...")
		case <- ch2:
			fmt.Printf("ch2 case...")
		default:
			//由于当前时间未到3秒，因此走default
			fmt.Printf("default go...")
	}
}
