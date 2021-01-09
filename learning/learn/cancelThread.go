package learn

import (
	"fmt"
	"sync"
	"time"
)

func cancelWorker(index int, wg *sync.WaitGroup, cancel chan bool)  {
	defer wg.Done()

	for{
		select {
		default:
			fmt.Println("worker:", index)
		case <-cancel:
			return
		}

		time.Sleep(100*time.Millisecond)
	}
}

func CancelThread()  {
	cancel := make(chan bool)
	var wg sync.WaitGroup

	for i:=0; i<10; i++{
		wg.Add(1)
		go cancelWorker(i+1, &wg, cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}