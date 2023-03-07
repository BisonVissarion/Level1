package main

import (
	"fmt"
)

type Counter struct {
	count int
}

func increaseCounter(counter *Counter, wait chan bool) {
	counter.count++
	wait <- true
}
func main() {
	counter := Counter{count: 0}
	wait := make(chan bool)
	for i := 0; i < 4000; i++ {
		go increaseCounter(&counter, wait)
		<-wait
	}
	fmt.Println(counter)
}
