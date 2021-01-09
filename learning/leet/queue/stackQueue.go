package queue

import "fmt"

func TestQueue() {
	qu := Constructor()
	qu.AppendTail(1)
	qu.AppendTail(2)
	fmt.Println(qu.DeleteHead())
	qu.AppendTail(3)
	fmt.Println(qu.DeleteHead())
	fmt.Println(qu.DeleteHead())
	fmt.Println(qu.DeleteHead())
}

type CQueue struct {
	Input  []int
	Output []int
}

func Constructor() CQueue {
	//instance := new(CQueue)
	instance := CQueue{
		Input:  make([]int, 0),
		Output: make([]int, 0),
	}
	return instance
}

func (this *CQueue) AppendTail(value int) {
	this.Input = append(this.Input, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.Output) == 0 {
		for index := len(this.Input) - 1; index >= 0; index-- {
			this.Output = append(this.Output, this.Input[index])
		}
		this.Input = this.Input[:0]
	}

	size := len(this.Output)
	if size == 0 {
		return -1
	}

	result := this.Output[size-1]
	this.Output = this.Output[:size-1]
	return result
}
