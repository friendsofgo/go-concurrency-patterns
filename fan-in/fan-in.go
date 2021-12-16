package main

import (
	"math/rand"
	"time"
)

func main() {
	ch1 := dataSource("Hello")
	ch2 := dataSource("Gophers!")

	ch := fanin(ch1, ch2)
	go func() {
		for v := range ch {
			println(v)
		}
	}()
	time.Sleep(2 * time.Second)
}

func fanin(chs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ch := range chs {
		go func(ch <-chan string) {
			for {
				c <- <-ch
			}
		}(ch)
	}
	return c
}

func dataSource(data string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			ch <- data
		}
	}()
	return ch
}
