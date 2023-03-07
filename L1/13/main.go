package main

//Поменять местами два числа без создания временной переменной.

import "fmt"

func main() {
	a := 1
	b := 2
	a, b = b, a
	fmt.Println(a, b)
}
