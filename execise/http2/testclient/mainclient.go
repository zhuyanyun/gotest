package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
)

func main() {
	//注意这里要使用证书中包含的主机名称
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", "zyy.centos:443", tlsConfig)
	_, err = io.WriteString(conn, "a")

	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()
	fmt.Println("====")

}
