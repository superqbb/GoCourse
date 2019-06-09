package main

import (
	"fmt"
	"sync"
	"time"
)

/*
为了方便演示，所以demo的main包用了多个main函数。
若此处SetNonBuffChan,SetBuffChan 不重命名，会提示undefined
因为mian包里，使用go run main.go，编译器只会加载main.go这个文件，不会加载main包里的其他文件，只有非main包里的文件才会通过依赖去自动加载。

golang推荐项目结构：
.
├── .gitignore
├── README.md
├── main.go
└── src
    ├── pkg1
    │   └── a.go
    ├── pkg2
    │   └── b.go
    └── pkg3
        └── c.go

*/

func SetNonBuffChan1(x chan int){
	time.Sleep(5*time.Second)
	x <- 1
}

func SetBuffChan1(y chan int){
	for i:=0;i<10;i++{
		time.Sleep(1*time.Second)
		y<-i
	}
	close(y)
}

func main() {
	wg := &sync.WaitGroup{}

	x := make(chan int)
	defer close(x)

	y := make(chan int,10)

	wg.Add(2)
	go SetNonBuffChan1(x)
	go SetBuffChan1(y)

	//开启两个goroutine同时读取channel x,y
	go func(){
		//range channel 可以直接取到 channel 中的值。一旦 channel 关闭，循环自动结束
		for ch := range y {
			fmt.Printf("y=%v\n",ch)
		}
		wg.Done()
	}()

	go func() {
		println("x=",<-x) //此处读取x时会阻塞
		wg.Done()
	}()

	wg.Wait()
}
