package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const requestIDKey = "rid"

//从request中获取requestid，然后写入上下文
func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
	reqID := req.Header.Get("Request-ID")
	if reqID == "" {
		reqID = strconv.FormatInt(time.Now().UnixNano(),10)
	}
	//使用WithValue函数，可以写入K-V到context
	return context.WithValue(ctx, requestIDKey, reqID)
}

func requestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}

//中间件接收一个http.Handler入参，返回一个http.Handler
func middleWare(next http.Handler) http.Handler {
	//实现了HandlerFunc接口的handler
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		//使用http.Request.Conext()新建一个context
		ctx := newContextWithRequestID(req.Context(), req)
		//使用http.Request.WithContext()，可以将context写入request实例中
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func h(w http.ResponseWriter, req *http.Request) {
	//从上下文获取requestID
	reqID := requestIDFromContext(req.Context())
	fmt.Fprintln(w, "Request ID: ", reqID)
	return
}

func main() {
	http.Handle("/", middleWare(http.HandlerFunc(h)))
	fmt.Println("Listing 0.0.0.0:9201")
	http.ListenAndServe(":9201", nil)
}