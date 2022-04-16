package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	fmt.Println("main spinning another goroutine")

	go func() {
		work()
		ch <- struct{}{}
	}() //fork point

	<-ch //join point
	fmt.Println("main done waiting")
}

func work() {
	fmt.Println("Goroutine started work")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Goroutine completed; time to join")
}
