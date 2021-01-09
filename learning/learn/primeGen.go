package learn

import (
	"context"
	"fmt"
)

func genNature(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("getNature done")
				return
			case ch <- i:
				//fmt.Println("genNature:", i)
			}
		}
	}()

	return ch
}

func filterPrime(ctx context.Context, in <-chan int, prime int) chan int{
	out := make(chan int)
	go func() {
		fmt.Println("filter", prime, "enter")
		for {
			i := <-in
			//fmt.Println("to filter", i, "by", prime)
			if i%prime != 0 {
				select {
				case <-ctx.Done():
					fmt.Println("filter", prime, "done")
					return
				case out <- i:
					//fmt.Println("filterPrime:", i)
				}
			}
		}
	}()

	return out
}

func GenPrime()  {
	ctx, cancel := context.WithCancel(context.Background())

	ch := genNature(ctx)
	count := 10
	for i:=0; i<count; i++{
		prime := <- ch
		fmt.Printf("%v: %v\n", i+1, prime)
		if i<count-1 {
			ch = filterPrime(ctx, ch, prime)
		}
	}

	cancel()
}