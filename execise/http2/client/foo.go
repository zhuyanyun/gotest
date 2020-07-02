package main

/*
#cgo LDFLAGS: -L./ -lfoo
#include <stdio.h>
#include <stdlib.h>
#include "foo.h"
*/
import "C"

func main() {
	C.foo()
}
