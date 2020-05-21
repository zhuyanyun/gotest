package main

var x = 100

func init() {
	println("init:", x)
	x++
}

func main2() {
	println("main :", x)

}
