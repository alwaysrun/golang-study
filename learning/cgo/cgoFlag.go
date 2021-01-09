package cgo

/*
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
#cgo linux CFLAGS: -DCGO_OS_LINUX=1

//char *os="1111";
#if defined(CGO_OS_WINDOWS)
	 char *os="windows";
#elif defined(CGO_OS_LINUX)
	 char *os="linux";
#else
	//char *os="unknow";
	#error(unknown os)
#endif
*/
import "C"

func PrintOS()  {
	print("to print C.OS: ")
	println(C.GoString(C.os))
}
