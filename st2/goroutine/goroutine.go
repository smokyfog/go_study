package main

import (
	"fmt"
	"runtime"
	"time"
)

func basic()  {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("hello from " + "goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}

func gosched()  {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched()	// 主动交出控制权
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func main() {
	gosched()
}


