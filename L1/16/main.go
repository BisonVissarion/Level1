package main

import "fmt"

//Реализовать быструю сортировку массива встроенными методами

func quicksort(arr []int, left int, right int) []int {

	//Создаю копии пришедших переменных
	l := left
	r := right

	//Вычисляю "центр"
	center := arr[(left+right)/2]

	fmt.Println(l, r, (left+right)/2)

	//Цикл, который начинает сортировку
	for l <= r {

		//Ищет значения больше "центра"
		for arr[r] > center {
			r--
		}

		for arr[l] < center {
			l++
		}

		//После прохода циклов проверяет счетчики циклов
		if l <= r {
			//Если условие true, меняет друг с другом
			arr[r], arr[l] = arr[l], arr[r]
			l++
			r--
		}
	}

	if r > left {
		quicksort(arr, left, r)
	}

	if l > right {
		quicksort(arr, l, right)
	}

	return arr
}

func main() {
	arr := []int{2, 7, 67, 11, 999, 89, 00, 32, 12, 454, 4545}
	arr = quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
