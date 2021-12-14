package main

import (
	"fmt"
	"time"
)

func main() {
	processor := make(chan string, 100)
	process(processor)

	tasks := 1000
	sent := 0
	dropped := 0
	for i := 0; i < tasks; i++ {
		select {
		case processor <- "some data":
			sent++
		default:
			time.Sleep(100 * time.Millisecond)
			dropped++
		}
	}
	close(processor)
	fmt.Println("sent:", sent, "dropped:", dropped)
}

func process(ch chan string) {
	go func() {
		for range ch {
		}
	}()
}
