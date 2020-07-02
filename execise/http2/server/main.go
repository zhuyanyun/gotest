package main

import (
	"fmt"
	"net/http"
)

func main() {
	//注册回调函数
	http.HandleFunc("/hello", handler)

	//绑定tcp监听地址，并开始接受请求，然后调用服务端处理程序来处理传入的连接请求。
	//params：第一个参数 addr 即监听地址；第二个参数表示服务端处理程序，通常为nil
	//当参2为nil时，服务端调用 http.DefaultServeMux 进行处理
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method = ", r.Method) //请求方法
	fmt.Println("URL = ", r.URL)       // 浏览器发送请求文件路径
	fmt.Println("header = ", r.Header) // 请求头
	fmt.Println("body = ", r.Body)     // 请求包体
	fmt.Println(r.RemoteAddr, "连接成功")  //客户端网络地址

	w.Write([]byte("hello http server"))
}
