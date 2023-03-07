package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "главрыба"
	//добавляем символ строки в слайс
	s := strings.Split(str, "")
	//проходимся по слайсу s
	for i := range s {
		//stdout
		fmt.Print(s[len(s)-1-i])
	}
}
