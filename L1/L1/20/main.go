package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "snow dog sun"
	//добавляем каждое слово в слайс
	w := strings.Split(s, " ")
	//проходимся по слайсу words
	for i := range w {
		//stdout
		fmt.Print(w[len(w)-i-1], " ")
	}
}
