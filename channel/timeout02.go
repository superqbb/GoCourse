package main

import "time"
import "fmt"
func main() {
	c1 := make(chan string, 1)
	timer1 := time.NewTimer(time.Second * 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-timer1.C:
		fmt.Println("timeout 1")
	}
}