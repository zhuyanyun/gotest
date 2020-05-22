package unittest

import (
	"os"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
}

func TestB(t *testing.T) {
	if os.Args[len(os.Args)-1] == "b" {
		//只有一个测试函数调用Parallel方法并没有效果，且go test 执行参数parallel必须大于1
		//  执行命令   go test -v -args "b"
		t.Parallel()
	}
	time.Sleep(time.Second * 2)
}
