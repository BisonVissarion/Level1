/**
Удалить i-ый элемент из слайса.
*/
package main

import "fmt"

func main() {
	b := []string{"G", "H", "J", "K", "L"}
	i := 2
	copy(b[i:], b[i+1:]) //Совершить сдвиг a[i+1:] влево на один индекс

	b[len(b)-1] = "" //Удалить последний элемент

	b = b[:len(b)-1] //Усечь срез

	fmt.Println(b)
}
