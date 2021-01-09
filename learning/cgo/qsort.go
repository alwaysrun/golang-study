package cgo

// #include <stdlib.h>
// typedef int (*qsort_cmp_func_t)(const void *a, const void *b);
// extern int go_qsort_compare(void *a, void *b);
import "C"
import (
	"fmt"
	"unsafe"
)

//export go_qsort_compare
func go_qsort_compare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}

type CompareFunc C.qsort_cmp_func_t

func QSort(base unsafe.Pointer, num, size int,
	cmp CompareFunc) {
	C.qsort(base, C.size_t(num), C.size_t(size), C.qsort_cmp_func_t(cmp))
}

func SortTest() {
	values := []int32{23, 43, 45, 21, 90, 76, 65, 54}

	//QSort(unsafe.Pointer(&values[0]),
	//	len(values), int(unsafe.Sizeof(values[0])),
	//	CompareFunc(C.go_qsort_compare))
	//GoQSort(unsafe.Pointer(&values[0]),
	//	len(values), int(unsafe.Sizeof(values[0])),
	//	func(a, b unsafe.Pointer) int {
	//		pa, pb := (*int32)(a), (*int32)(b)
	//		return int(*pa - *pb)
	//	})
	SliceQSort(values, func(i, j int) bool {
		return  values[i]<values[j]
	})

	fmt.Println(values)
}
