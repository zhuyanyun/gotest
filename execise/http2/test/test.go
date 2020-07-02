package main

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"fmt"
)

func main() {
	data := []byte("[这里才是一个完整的数据包]")
	l := len(data)
	fmt.Println(l)
	magicNum := make([]byte, 4)
	binary.BigEndian.PutUint32(magicNum, 0x123456)
	lenNum := make([]byte, 2)
	binary.BigEndian.PutUint16(lenNum, uint16(l))
	packetBuf := bytes.NewBuffer(magicNum)
	packetBuf.Write(lenNum)
	packetBuf.Write(data)

	//注意这里要使用证书中包含的主机名称
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", "zyy.centos:443", tlsConfig)
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	for i := 0; i < 1000; i++ {
		_, err = conn.Write(packetBuf.Bytes())
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
}
