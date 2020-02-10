package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500))  * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

type worker struct {
	in chan int
}

func createWorker(id int) chan <-int {
	c := make(chan int)
	go worker(id, c)
	return c
}


func main() {
	var c1, c2 = generator(), generator()
	w := c
	for {
		select {
		case n := <-c1:
			fmt.Println("Received from c1:", n)
		case n := <-c2:
			fmt.Println("Received from c2:", n)
		}
	}

}
