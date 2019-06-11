package main



import (
	"sync"
	"time"
)

var l sync.Mutex
var a string

func f() {
	time.Sleep(5*time.Second)
	a = "hello, world"
	print("21,")
	l.Unlock()
	print("22,")
}

func main() {
	l.Lock()
	print("1,")
	go f()
	print("2,")
	l.Lock()
	print("3,")
	print(a)
}