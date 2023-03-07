package main

import (
	"fmt"
)

func main() {
	showType(777.777)
	showType(make(chan int))
	showType("Hello World")
	showType(999)
}

// определяет типа через fmt.Printf()
func showType(d interface{}) {
	fmt.Printf("type  %T\n", d)
}
