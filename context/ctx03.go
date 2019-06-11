package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	//http://localhost:8989/debug/pprof/goroutine?debug=1
	go http.ListenAndServe(":8989", nil)
	//这里调用了context.WithCancel()，我们也可以使用context.WithTimeout()和context.WithDeadline()来设置goroutine的超时时间和最终的运行时间
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * time.Second)
		//调用cancel后，context会广播Done消息，A,B,C都会收到done信号，因此打印顺序随机
		//A,B,C捕获Done信号后，分别关闭了goroutine
		cancel()
	}()
	log.Println(A(ctx))
	select {}
}

func C(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return "C Done"
	}
	return ""
}

func B(ctx context.Context) string {
	go log.Println(C(ctx))
	select {
	case <-ctx.Done():
		return "B Done"
	}
	return ""
}

func A(ctx context.Context) string {
	go log.Println(B(ctx))
	select {
	case <-ctx.Done():
		return "A Done"
	}
	return ""
}