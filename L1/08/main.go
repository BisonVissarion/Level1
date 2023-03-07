package main

//Дана переменная int64. Разработать программу
// которая устанавливает i-й бит в 1 или 0.

import "fmt"

func MoveBit(digit, moveBit, value int64) int64 {

	if value == 1 {
		return digit | int64(1<<moveBit)
	} else {
		return digit &^ int64(1<<moveBit)
	}
}

func main() {
	var digit, moveBit, value int64
	fmt.Println("Введите число, смещение на бит(от 0 до 63) и ноль или единицу")
	fmt.Scan(&digit, &moveBit, &value)

	fmt.Printf("До: %064b\n", uint64(digit))
	newDigit := MoveBit(digit, moveBit, value)
	fmt.Printf("После: %064b\n", uint64(newDigit))

}
