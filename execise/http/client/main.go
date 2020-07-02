package main

import (
	"net"
	"time"
)

func main() {

	client()

	time.Sleep(time.Second * 4)
}

//客户端实现
func client() {
	con, _ := net.Dial("tcp", "zyy.centos:443")
	con.Write([]byte("hello wrold"))
	con.Close()

	//r := bufio.NewReaderSize(con, 5000) //缓冲区大小大于数据包
	//data := make([]byte, 100)
	//for i := 0; i < 5; i++ {
	//	data = append(data, data...)
	//	copy(data, []byte("PING"))
	//	con.Write(data)
	//	con.Write([]byte{' '}) //以空格作为结束符
	//
	//	//接收以空格符作为结束的数据包
	//	s, _ := r.ReadSlice(' ')
	//	msg := string(s[:len(s)])
	//
	//	fmt.Printf("%s %v\n", msg, len(msg))
	//}
	//
	//fmt.Println(r)
	//fmt.Println(con)
	//con.Close()

	//跳过证书验证
	//tr := &http.Transport{
	//	TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	//}
	//client := &http.Client{Transport: tr}
	//resp, err := client.Get("https://zyy.centos")
	//
	//if err != nil {
	//	fmt.Println("error:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

}
