package main

/*
#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int id;
}ctx;

ctx *createCtx(int id) {
    ctx *obj = (ctx *)malloc(sizeof(ctx));
    obj->id = id;
    return obj;
}
*/
import "C"
import (
	"fmt"
	//"sync"
)

func main() {
	var ctx *C.ctx = C.createCtx(1000)
	fmt.Printf("id : %d\n", ctx.id)
}
