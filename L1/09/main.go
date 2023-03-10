package main

import (
	"fmt"
	"time"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

// Получаем массив,записываем его в 1 канал
func write(arr []int, out chan<- int) {
	for _, val := range arr {
		out <- val
	}
	close(out)
}

// возводим в квадрат,записываем во 2 канал,закрываем 2 канал
func squarer(inp <-chan int, out chan<- int) {
	for val := range inp {
		out <- val * val
	}
	close(out)
}

// Читаем из 2 канала, записываем в stdout
func print(inp <-chan int) {
	for val := range inp {
		fmt.Println(val)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	go write(arr, ch1)
	go squarer(ch1, ch2)
	go print(ch2)
	time.Sleep(10 * time.Millisecond)
}
