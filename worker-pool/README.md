## Worker Pool - a concurrency pattern.

- There's a pool/queue of tasks/jobs and a fixed number of workers.
- The workers pick task from the pool, perform and complete the job and make themselves available and pick another task until the pool is empty.

In Go, we implement the workers with goroutines and the pool of jobs using buffered channel.

## Why we need Worker Pool?

- If I want to perform n number of jobs concurrently, I can spawn off n goroutines and get them done. Then why we need the worker pool pattern?
- The answer is that resources are limited. Each goroutine needs some memory, though very less. With this pattern we are sure that there would be limited goroutines working at a time and also the jobs will be performed concurrently.