package main

import (
	"fmt"
	"net/http"
)

// http server
func sayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你好，ares！")
}
func main() {
	http.HandleFunc("/", sayHi)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("Http 服务建立失败，err：", err)
		return
	}
}
