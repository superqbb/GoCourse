package main

import "fmt"

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	/*
	当select与for一起使用时，break可能会跳不出循环
	需要配合标签或goto跳出循环
	*/
	//LOOP:
	for {
		select {
		case c <- x:
			x, y = y, x+y
			if x>=8 {
				//break LOOP
				//goto GOTOLOOP
				break
			}
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
	//GOTOLOOP:

}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}