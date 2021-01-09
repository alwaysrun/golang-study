package learn

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var totalValue struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10000; i++ {
		totalValue.Lock()
		totalValue.value += i
		totalValue.Unlock()
	}
}

var totalN uint64

func atomWorker(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i < 10000; i++ {
		atomic.AddUint64(&totalN, i)
	}
}

func DoWork() {
	var wg sync.WaitGroup
	wg.Add(4)

	go worker(&wg)
	go worker(&wg)
	//go worker(&wg)
	//go worker(&wg)

	go atomWorker(&wg)
	go atomWorker(&wg)

	wg.Wait()

	fmt.Println("TotalValue:", totalValue.value)
	fmt.Println("TotalN:", totalN)
}
