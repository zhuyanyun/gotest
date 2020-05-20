package main

import "fmt"

type Handler interface {
	Do(k, v interface{})
}

type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandlerFunc(f))
}

//type welcome string

func selfInfo(k, v interface{}) {
	fmt.Printf("大家好, 我叫%s, 今年%d岁 \n", k, v)
}

func Do(k, v interface{}) {
	fmt.Printf("%s, 我叫%s, 今年%d岁 \n", k, v)
}

func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["zyy"] = 28
	persons["王五"] = 55

	//var w welcome = "大家好"

	//Each(persons, w)
	//Each(persons, HandlerFunc(w.selfInfo))
	EachFunc(persons, selfInfo)
}
