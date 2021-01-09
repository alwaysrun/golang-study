package list

import (
	"fmt"
)

func TestList() {
	//reversePrint(nil)
	ary := []int{1, 2, 3}
	fmt.Println(ary)
	Reverse(ary)
	fmt.Println(ary)
}

func Reverse(in interface{}) {
	ary := in.([]int)
	size := len(ary)-1
	for index := 0; index < (size+1)/2; index++ {
		tmp := ary[index]
		//ary[index], ary[size-index] = ary[size-index], ary[index]
		ary[index] = ary[size-index]
		ary[size-index] = tmp
	}
}

func reversePrint(head *ListNode) []int {
	result := make([]int, 0)
	for ; head != nil; head = head.Next {
		result = append(result, head.Val)
	}

	//Reverse(result)
	for index,last := 0, len(result)-1; index<last; index++{
		result[index],result[last] = result[last], result[index]
		last--
	}
	return result
}
