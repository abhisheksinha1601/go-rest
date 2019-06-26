package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Starting")

	c := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go r1(c, &wg)
	go r2(c, &wg)

	// fmt.Println(<-c)

	// go func() {
	// 	fmt.Println("Starting routine r")
	// 	time.Sleep(5 * time.Second)
	// 	fmt.Println("Done routine r")
	// 	c <- 15
	// }()

	// fmt.Println(<-c)
	// fmt.Println(<-c)

	go func() {
		fmt.Println("Starting routine r")
		fmt.Println("Len", len(c))
		for i := range c {
			fmt.Println("Got", i)
		}
		fmt.Println("Len2", len(c))   // This line doesn't print
		fmt.Println("Done routine r") // This 1 too
	}()

	wg.Wait()

	// fmt.Println("Len2", len(c))
	// for i := range c {
	// 	fmt.Println("Got here", i)
	// }

	fmt.Println("Done")
}

func r1(c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Starting routine r1")
	time.Sleep(5 * time.Second)
	fmt.Println("Done routine r1")
	c <- 10
}

func r2(c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Starting routine r2")
	time.Sleep(2 * time.Second)
	fmt.Println("Done routine r2")
	c <- 20
}
