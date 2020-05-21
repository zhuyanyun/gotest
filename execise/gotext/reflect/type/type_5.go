package main

import (
	"fmt"
	"reflect"
	"time"
)

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func main5() {
	var m manager
	t := reflect.TypeOf(&m)

	if t.Kind() == reflect.Ptr { //获取指针的基类型
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type, f.Offset)
		if f.Anonymous { //输出匿名字段结构
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				fmt.Println(" ", af.Name, af.Type)
			}
		}
	}

	println("=======================", "对于匿名字段，可用多级索引（按定义顺序）直接访问")
	time.Sleep(time.Second)

	s := reflect.TypeOf(m)
	name, _ := s.FieldByName("name") //按名称查找
	fmt.Println(name.Name, name.Type)

	age := s.FieldByIndex([]int{0, 1}) //按多级索引查找
	fmt.Println(age.Name, age.Type)

}
