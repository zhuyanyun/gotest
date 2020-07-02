package main

import (
	"crypto/tls"
	"io"
	"log"
	"time"
)

//时间数据做了一点改动，因为毫秒总会有误差
var stimes = [25]int{122, 163, 204, 205, 245, 246, 247, 248, 249, 250, 251, 252,
	285, 286, 287, 288, 289, 290, 291, 325, 326, 327, 328, 759, 844,
}

//var stimes = [25] int {122, 163, 204, 205, 245, 246, 247, 248, 249, 250, 251, 252,
//						 278, 279, 280, 281, 282, 283, 284, 315, 316, 317, 318, 746, 844,
//					}

//注意有重发的情况，只需要发一半
var packageCount = [25]int{4, 2, 4, 4, 8, 4, 2, 4, 8, 2, 2, 6,
	4, 6, 4, 4, 8, 6, 6, 2, 4, 4, 2, 6, 2,
}
var packageLength = [25]int{26586, 52548, 49896, 55416, 44592, 11256, 11148, 22296, 38466, 8388, 5628, 16884,
	16776, 36204, 38856, 22296, 55632, 41724, 44484, 8388, 22296, 56822, 11358, 8376, 108,
}

//var slice = make([]string, 5)
var slice = make([]string, 0)

func init() {
	//330
	//a := "中新网6月15日电 综合报道，美国一名现年70岁的新冠患者在医院度过62天后终于治愈，可当他出院后，却发现自己账单总额超过了110万美元。由于拥有医疗保险，他应该不必由于拥有医疗保险，他应该不必应该不必由于拥有医疗保险，他应该不必他的a"
	a := "中新网6月15日电 综合报道，美国一名现年70岁的新冠患"
	//var count int = 0
	for i := 0; i < len(packageCount); i++ {
		var m = packageCount[i] / 2
		var plength = packageLength[i] / packageCount[i]

		for k := 0; k < m; k++ {
			var c = plength / len(a)
			str := ""
			for j := 0; j < c; j++ {
				str = str + a
			}
			slice = append(slice, str)
			//count ++
		}
	}
}

func main() {
	//注意这里要使用证书中包含的主机名称
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", "zyy.centos:443", tlsConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//defer conn.Close()

	//log.Println("Client Connect To ", conn.RemoteAddr())
	//status := conn.ConnectionState()
	//fmt.Printf("%#v\n", status)

	var count int = 0
	var j int = 0
	//第一个 1 毫秒需要睡
	for i := 0; i < len(stimes); i++ {
		if i == 0 {
			//time.Sleep(time.Millisecond * time.Duration(stimes[i] - 12))
			time.Sleep(time.Millisecond * time.Duration(stimes[i]-16))
		} else {
			time.Sleep(time.Millisecond * time.Duration(stimes[i]-stimes[i-1]))
		}
		for ; j < len(packageCount); j++ {
			var pkCount int = packageCount[j] / 2
			for {
				str := slice[count]
				_, err = io.WriteString(conn, str)
				pkCount--
				count++
				if pkCount <= 0 || count >= len(slice) {
					break
				}
			}
			if pkCount <= 0 {
				j++
				break
			}
		}
	}
}
