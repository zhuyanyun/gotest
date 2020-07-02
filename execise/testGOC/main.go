package main

/*
#include "savepacket.c"
#include "function_send.c"

*/
import "C"

import "fmt"

func main() {
	C.save_packet()
	fmt.Println("======")

}
