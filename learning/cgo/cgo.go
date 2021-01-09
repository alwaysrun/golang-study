package cgo
import "C"

/*
#include <stdio.h>
#include <stdlib.h>

static void SayHello(const char *str){
	printf(str);
}
*/
import "C"

import "unsafe"

func CgoTest()  {
	s := "Hello CGO\n"
	cs := C.CString(s)
	C.SayHello(cs)
	C.free(unsafe.Pointer(cs))
}
