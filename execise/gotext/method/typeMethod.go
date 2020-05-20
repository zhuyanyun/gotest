package main

import "fmt"

type person struct {
	name string
}

func (p person) string() string {
	return "the person name is:" + p.name
}

func (p person) modify() {
	p.name = "李四"
}

func (p *person) modifyutpr() {
	p.name = "李四"
}

func main() {
	p := person{name: "张三"}
	p.modify()
	fmt.Println(p.string())

	p.modifyutpr()    //go编辑器会自动帮取指针
	(&p).modifyutpr() //都可以
	fmt.Println((&p).string())

}
