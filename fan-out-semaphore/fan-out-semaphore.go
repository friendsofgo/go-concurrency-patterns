package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	workers := 1000
	ch := make(chan string, workers)

	concurrency := 100
	sem := make(chan bool, concurrency)

	for i := 0; i < workers; i++ {
		go dataSource(ch, sem, i)
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

func dataSource(ch chan string, sem chan bool, worker int) {
	for {
		// acquire
		sem <- true

		// do work
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		ch <- fmt.Sprintf("<worker %d> processing data", worker)

		// release
		<-sem
	}
}
