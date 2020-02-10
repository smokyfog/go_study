package main

import (
	"fmt"
	"sync"
)

func doWork(id int, c chan int, w worker)  {
	for n := range w.in {
		fmt.Printf("worker %d received %d\n", id, n)
		w.done()
	}
}

type worker struct {
	in chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan  int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w.in, w)
	return w
}

func chanDemo()  {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i,&wg)
	}

	wg.Add(20)
	for i, worker := range workers{
		worker.in <- 'a' + i
	}
	for i, worker := range workers{
		worker.in <- 'A' + i
	}
	wg.Wait()

	// wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}
}
func main() {
	chanDemo()
}
