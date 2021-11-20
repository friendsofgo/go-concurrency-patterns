package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	workers := 1000
	ch := make(chan string, workers)
	for i := 0; i < workers; i++ {
		go dataSource(ch, i)
	}

	for {
		if workers <= 0 {
			break
		}
		select {
		case v := <-ch:
			workers--
			fmt.Println(v)
		}
	}
}

func dataSource(ch chan string, worker int) {
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			ch <- fmt.Sprintf("<worker %d> processing data", worker)
		}
	}()
}
