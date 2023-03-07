//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.

package main

import "fmt"

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	// возвращает обратный стринг
	return string(rns)
}

func main() {

	// обращает стринг
	str := "статуя"

	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)
}
