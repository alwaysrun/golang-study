package learn

import (
	"fmt"
	"time"
)

func GoOutput(chDone chan<-bool)  {
	fmt.Println("GoOutput start")
	//chDone <- true
	time.Sleep(time.Second)
	close(chDone)

	time.Sleep(time.Second)
	fmt.Println("GoOutput over")
}

func ChTest()  {
	ch := make(chan bool)

	go GoOutput(ch)
	fmt.Println(<-ch)

	fmt.Println("Test over")
}

//////////////////////////////////////////
// producer-customer 
func customer(in <-chan int)  {
	for v:=range in{
		fmt.Println("Consume:", v)
	}

	fmt.Println("customer over")
}

func producer(fac int, out chan<- int)  {
	defer func() {
		if r:=recover(); r!=nil{
			fmt.Println(r)
		}
	}()

	for i:=0; ; i++{
		out<-i*fac
	}

	fmt.Println("producer over")
}

func ProduceConsume()  {
	ch := make(chan int, 16)
	go producer(3, ch)
	go producer(5, ch)
	go customer(ch)

	time.Sleep(3*time.Second)
	close(ch)
}