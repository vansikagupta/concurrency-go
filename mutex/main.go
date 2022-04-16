package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	m  sync.Mutex
	rw sync.RWMutex
	x  int
}

func read(id int, d *data, wg *sync.WaitGroup) {
	defer wg.Done()
	d.m.Lock()
	defer d.m.Unlock()
	time.Sleep(time.Second * 2)
	fmt.Printf("Goroutine %d reads value of x %d\n", id, d.x)
}

func write(id int, d *data, wg *sync.WaitGroup) {
	defer wg.Done()
	d.m.Lock()
	defer d.m.Unlock()
	time.Sleep(time.Second * 2)
	d.x = d.x + 1
	fmt.Printf("Goroutine %d writes value of x %d\n", id, d.x)
}

func rwread(id int, d *data, wg *sync.WaitGroup) {
	defer wg.Done()
	d.rw.RLock()
	defer d.rw.RUnlock()
	time.Sleep(time.Second * 2)
	fmt.Printf("Goroutine %d reads value of x %d\n", id, d.x)
}

func rwwrite(id int, d *data, wg *sync.WaitGroup) {
	defer wg.Done()
	d.rw.Lock()
	defer d.rw.Unlock()
	time.Sleep(time.Second * 2)
	d.x = d.x + 1
	fmt.Printf("Goroutine %d writes value of x %d\n", id, d.x)
}

func main() {
	d := &data{x: 1}
	var wg sync.WaitGroup

	wg.Add(4)

	/*
		go read(1, d, &wg)
		go write(2, d, &wg)
		go read(3, d, &wg)
		go write(4, d, &wg)
	*/

	go rwread(1, d, &wg)
	go rwwrite(2, d, &wg)
	go rwread(3, d, &wg)
	go rwwrite(4, d, &wg)

	/*
		Multiple readers can hold the read lock, but only one writer.
		When writer is trying to lock and lock is not available, it will block until lock is available
		Use RWMutex for read extensive operations
	*/

	wg.Wait()
}
