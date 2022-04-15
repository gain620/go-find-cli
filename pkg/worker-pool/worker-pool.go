package worker_pool

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	"sync/atomic"
	"time"
)

func worker2(done chan struct{}) {
	fmt.Println("working...")
	time.Sleep(time.Second * 3)
	fmt.Println("done")

	done <- struct{}{}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		res := j*j + id
		results <- res
		log.Printf("Worker Id[%v]: Job [%v], Result [%v]\n", id, j, res)
		time.Sleep(time.Millisecond * 300)
	}
}

func TestWorkerPool() {
	done := make(chan struct{}, 1)
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker2(done)
	// init 3 workers
	for id := 0; id < 3; id++ {
		go worker(id, jobs, results)
	}

	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 10; i++ {
		<-results
	}

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 4)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("WAITING!", i)
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	jobs2 := make(chan int, 5)
	done2 := make(chan struct{})

	go func() {
		for {
			j, more := <-jobs2
			if more {
				fmt.Println("received job", j, more)
			} else {
				fmt.Println("received all jobs", j, more)
				done2 <- struct{}{}
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs2 <- j
		fmt.Println("sent job", j)
	}
	close(jobs2)
	fmt.Println("sent all jobs")

	<-done2

	var ops uint64 = 0
	var mutex = &sync.Mutex{}
	for i := 0; i < 50; i++ {
		go func() {
			for {
				mutex.Lock()
				atomic.AddUint64(&ops, 1)
				mutex.Unlock()
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
