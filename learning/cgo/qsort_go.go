package cgo

// #include <stdlib.h>
// typedef int (*qsort_cmp_func_t)(const void *a, const void *b);
// extern int _cgo_qsort_compare(void *a, void *b);
// extern int _slice_qsort_compare(void *a, void *b);
import "C"

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

var go_qsort_compare_info struct {
	fun func(a, b unsafe.Pointer) int
	sync.Mutex
}

//export _cgo_qsort_compare
func _cgo_qsort_compare(a, b unsafe.Pointer) C.int {
	return C.int(go_qsort_compare_info.fun(a, b))
}

func GoQSort(base unsafe.Pointer, count, size int, cmp func(a, b unsafe.Pointer) int) {
	go_qsort_compare_info.Lock()
	defer go_qsort_compare_info.Unlock()

	go_qsort_compare_info.fun = cmp
	C.qsort(base, C.size_t(count), C.size_t(size),
		C.qsort_cmp_func_t(C._cgo_qsort_compare))
}

var slice_compare_info struct {
	base  unsafe.Pointer
	count int
	size  int
	less  func(a, b int) bool
	sync.Mutex
}

//export _slice_qsort_compare
func _slice_qsort_compare(a, b unsafe.Pointer) C.int {
	var (
		base = uintptr(slice_compare_info.base)
		size = uintptr(slice_compare_info.size)
	)

	i := int((uintptr(a) - base) / size)
	j := int((uintptr(b) - base) / size)

	switch {
	case slice_compare_info.less(i, j):
		return -1
	case slice_compare_info.less(j, i):
		return +1
	default:
		return 0
	}
}

func SliceQSort(slice interface{}, less func(i, j int) bool) {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("SliceQsort called with none-slice value"))
	}

	if sv.Len() == 0 {
		return
	}

	slice_compare_info.Lock()
	defer slice_compare_info.Unlock()

	defer func() {
		slice_compare_info.base = nil
		slice_compare_info.count = 0
		slice_compare_info.size = 0
		slice_compare_info.less = nil
	}()

	slice_compare_info.base = unsafe.Pointer(sv.Index(0).Addr().Pointer())
	slice_compare_info.count = sv.Len()
	slice_compare_info.size = int(sv.Type().Elem().Size())
	slice_compare_info.less = less

	C.qsort(
		slice_compare_info.base,
		C.size_t(slice_compare_info.count),
		C.size_t(slice_compare_info.size),
		C.qsort_cmp_func_t(C._slice_qsort_compare),
	)
}
