package main

import (
	"bytes"
	"encoding/binary"
	. "fmt"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

type Handle uintptr

type TCPHeader struct {
	SrcPort   uint16
	DstPort   uint16
	SeqNum    uint32
	AckNum    uint32
	Offset    uint8
	Flag      uint8
	Window    uint16
	Checksum  uint16
	UrgentPtr uint16
}

type PsdHeader struct {
	SrcAddr   uint32
	DstAddr   uint32
	Zero      uint8
	ProtoType uint8
	TcpLength uint16
}

func inet_addr(ipaddr string) uint32 {
	var (
		segments []string = strings.Split(ipaddr, ".")
		ip       [4]uint64
		ret      uint64
	)
	for i := 0; i < 4; i++ {
		ip[i], _ = strconv.ParseUint(segments[i], 10, 64)
	}
	ret = ip[3]<<24 + ip[2]<<16 + ip[1]<<8 + ip[0]
	return uint32(ret)
}

func htons(port uint16) uint16 {
	var (
		high uint16 = port >> 8
		ret  uint16 = port<<8 + high
	)
	return ret
}

func CheckSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)

	return uint16(^sum)
}

func main() {
	var (
		msg       string
		psdheader PsdHeader
		tcpheader TCPHeader
	)

	//Printf("Input the content: ")
	//Scanf("%s", &msg)
	msg = "美国大型科技股多数上涨，苹果涨2.3%，亚马逊跌0.46%，奈飞涨0.87%，谷歌涨2.54%，Facebook涨2.11%，微软涨1.07%。\n\n全球多国新冠病毒疫情反弹，但受惠波音737 Max进行关键认证测试，加上成屋签约销售指数急弹，带动美股上涨。航空股和汽车股普涨。道指成分股波音收盘大涨逾14%。\n\n美国全国房地产经纪人协会（NAR）报告称，美国5月二手房销售签约指数环比增长44.3%，创历史上（2001年有数据记录以来）最大单月增幅。该指数上扬至99.6%，但仍然低于2月份111.4的水平。\n\n美联储主席鲍威尔表示，数据出现部分积极迹象，但美国失去了数以百万计的工作岗位；重申美联储将致力于使用各种货币工具。\n\n鲍威尔还称，未来的路径极端不确定，取决于疫情发展；将公布的数据开始表明经济活动恢复；除非人们认为情况安全，否则不可能全面复苏。\n\n欧洲主要股指也全线收涨，德股英股涨幅均超1%。截止收盘，德国法兰克福股市DAX指数涨幅为1.18%，英国伦敦股市《金融时报》100种股票平均价格指数涨幅为1.08%，法国巴黎股市CAC40指数涨幅为0.73%。\n\n国际原油再度大涨 美页岩油巨头申请破产\n\n美东时间周一，国际油价再度大涨。截止收盘，纽约8月原油期货收涨1.21美元，涨幅3.14%，报39.70美元/桶。布伦特8月原油期货收涨0.69美元，涨幅1.68%，报41.71美元/桶。"

	/*填充TCP伪首部*/
	psdheader.SrcAddr = inet_addr("10.38.4.191")
	psdheader.DstAddr = inet_addr("10.38.4.171")
	psdheader.Zero = 0
	psdheader.ProtoType = syscall.IPPROTO_TCP
	psdheader.TcpLength = uint16(unsafe.Sizeof(TCPHeader{})) + uint16(len(msg))
	Println(len(msg))

	/*填充TCP首部*/
	tcpheader.SrcPort = htons(7000)
	tcpheader.DstPort = htons(80)
	tcpheader.SeqNum = 0
	tcpheader.AckNum = 0
	tcpheader.Offset = uint8(uint16(unsafe.Sizeof(TCPHeader{}))/4) << 4
	tcpheader.Flag = 2 //SYN
	tcpheader.Window = 60000
	tcpheader.Checksum = 0

	/*buffer用来写入两种首部来求得校验和*/
	var (
		buffer bytes.Buffer
	)
	binary.Write(&buffer, binary.BigEndian, psdheader)
	binary.Write(&buffer, binary.BigEndian, tcpheader)
	tcpheader.Checksum = CheckSum(buffer.Bytes())

	/*接下来清空buffer，填充实际要发送的部分*/
	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, tcpheader)
	binary.Write(&buffer, binary.BigEndian, msg)

	/*下面的操作都是raw socket操作，大家都看得懂*/
	var (
		sockfd int
		addr   syscall.SockaddrInet4
		err    error
	)
	sockfd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_TCP)
	//sockfd2, err = syscall.Socket(2, 3, 6)

	if err != nil {
		Println("Socket() error: ", err.Error())
		return
	}
	defer syscall.Shutdown(sockfd, syscall.SHUT_RDWR)
	addr.Addr[0], addr.Addr[1], addr.Addr[2], addr.Addr[3] = 10, 38, 4, 191
	addr.Port = 80
	if err = syscall.Sendto(sockfd, buffer.Bytes(), 0, &addr); err != nil {
		Println("Sendto() error: ", err.Error())
		return
	}
	Println("Send success!")
}
