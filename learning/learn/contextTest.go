package learn

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func contextWorker(index int, ctx context.Context, wg *sync.WaitGroup)  {
	defer wg.Done()
	
	for{
		select {
		default:
			fmt.Println("worker:", index)
		case <-ctx.Done():
			return
		}
		
		time.Sleep(100*time.Millisecond)
	}
}

func ContextTest()  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup
	for i:=0; i<10; i++{
		wg.Add(1)
		go contextWorker(i+1, ctx, &wg)
	}

	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}
