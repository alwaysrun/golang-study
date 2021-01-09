package main

import (
	"fmt"
	"learning/leet/queue"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//fmt.Println("hello")
	//cgo.CgoTest()
	//cgo.PrintOS()
	//cgo.FunTest()
	//cgo.SortTest()
	//optStr.Test()

	//optStr.TestString()
	//.TestList()
	//optList.TestList()
	queue.TestQueue()
	//number.TestNumber()

	//cmdLine()
	//learn.DoWork()
	//learn.ChTest()
	//learn.ProduceConsume()
	//learn.PubSub()
	//learn.GenPrime()
	//for i:=0; i<10; i++ {
	//	learn.GenRandom(10)
	//}
	//learn.CancelThread()
	//learn.ContextTest()
	//err := learn.RecoverTest()
	//fmt.Println(err)

	// Ctrl+C
	fmt.Println("Print Ctrl+C to quit:")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%V)\n", <-sig)
}

func cmdLine() {
	fmt.Println(os.Args, len(os.Args))

	var all []string
	for _, v := range os.Args {
		all = append(all, v)
	}

	fmt.Println("All", all)
}
