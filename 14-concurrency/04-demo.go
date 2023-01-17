package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {

	rand.Seed(20)
	wg := &sync.WaitGroup{}
	count := flag.Int("count", 0, "Number of goroutines to execute")
	proc_count := flag.Int("pcount", runtime.NumCPU(), "Number of processors to use")
	flag.Parse()
	runtime.GOMAXPROCS(*proc_count)
	fmt.Printf("main starting %d goroutines using %d processors.. Hit ENTER to continue....\n", *count, *proc_count)
	fmt.Scanln()
	start := time.Now()
	for i := 1; i <= *count; i++ {
		wg.Add(1)    //increment the counter by 1
		go fn(i, wg) //schdeuling this function execution through the scheduler
	}
	wg.Wait() // wait for the counter to become 0
	elapsed := time.Since(start)
	fmt.Printf("main completed in %v time.. Hit ENTER to shutdown\n", elapsed)
	fmt.Scanln()
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(7)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
