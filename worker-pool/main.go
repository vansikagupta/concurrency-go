package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id  int
	num int
}

type Result struct {
	jobId int
	val   int
}

//worker should listen on the jobs pool, pick a job, perform and write the result and continue
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		//perform on the job
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d performing on Job %d\n", id, j.id)
		//write the result
		results <- Result{j.id, j.num * j.num}
	}
}

func main() {

	//create pool
	jobs := make(chan Job)
	results := make(chan Result, 20)

	var wg sync.WaitGroup

	//create fixed number of workers
	for i := 0; i < 7; i++ {
		wg.Add(1)
		go worker(i+1, jobs, results, &wg)
	}

	timeNow := time.Now()
	go func(jobs chan Job) {
		//add tasks to job pool
		for i := 0; i < 20; i++ {
			random := rand.Intn(99)
			jobs <- Job{i + 1, random}
		}
		//closing the jobs channel is important otherwise workers will keep waiting on it forever
		close(jobs)
	}(jobs)

	//waiting on all worker pools to finish
	wg.Wait()

	//see this change by changing number of workers
	fmt.Println("All jobs done in ", time.Since(timeNow))

	close(results)

	//read the results
	for r := range results {
		fmt.Println(r)
	}
}
