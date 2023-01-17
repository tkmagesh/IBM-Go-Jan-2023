package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()
	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 200
	}()

	go func() {
		time.Sleep(3 * time.Second)
		d3 := <-ch3
		fmt.Println("data from ch3 :", d3)
	}()

	/*
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			fmt.Println(<-ch1)
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			fmt.Println(<-ch2)
			wg.Done()
		}()
		wg.Wait()
	*/
	for i := 0; i < 3; i++ {
		select {
		case d1 := <-ch1:
			fmt.Println(d1)
		case d2 := <-ch2:
			fmt.Println(d2)
		case ch3 <- 300:
			fmt.Println("data sent to ch3")
			/* default:
			fmt.Println("No channel operations were successful!") */
		}
	}
}
