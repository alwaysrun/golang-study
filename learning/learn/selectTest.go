package learn

import (
	"fmt"
	"strconv"
)

func GenRandom(count int)  {
	ch := make(chan int)
	go func(count int) {
		for i:=0; i<count; i++{
			select {
			case ch<-0:
			case ch<-1:
			case ch<-2:
			case ch<-3:
			case ch<-4:
			case ch<-5:
			case ch<-6:
			case ch<-7:
			case ch<-8:
			case ch<-9:
			}
		}

		close(ch)
	}(count)

	r := ""
	for v:= range ch{
		//fmt.Println(v)
		r += strconv.Itoa(v)
	}

	fmt.Println(r)
}
