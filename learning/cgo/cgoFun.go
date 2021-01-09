package cgo

/*
#include <errno.h>

static int div(int a, int b){
	if(b==0){
		errno = EINVAL;
		return 0;
	}
	return a/b;
}
*/
import "C"

import "fmt"

func FunTest()  {
	v, err := C.div(2,1)
	fmt.Println(v, err)

	v, err = C.div(1,0)
	fmt.Println(v, err)
}