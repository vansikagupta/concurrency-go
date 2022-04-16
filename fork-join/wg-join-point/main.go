package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("main spinning another goroutine")
	wg.Add(1)
	go func() {
		work()
		wg.Done()
	}() //fork point

	wg.Wait() //join point
	fmt.Println("main done waiting")
}

func work() {
	fmt.Println("Goroutine started work")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Goroutine completed; time to join")
}
