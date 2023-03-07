package main

import (
	"fmt"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
//квадратов(4+16+36….) с использованием конкурентных вычислений.

func Summ(arr []int) (sum int) {
	ch := make(chan int)
	go func() {
		for _, a := range arr {
			go func(a int) {
				ch <- (a * a) //раздельно запускаю горутины, производящие вычисление
			}(a)
		}
	}()
	for i := 0; i < len(arr); i++ {
		sum += <-ch //получаю значение от горутин в количестве len(arr) раз
	}
	close(ch)
	return sum
}

func main() {
	fmt.Println(Summ([]int{2, 4, 6, 8, 10}))
}
