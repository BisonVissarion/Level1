package main

import "fmt"

//Разработать программу, которая в рантайме способна определить тип
//переменной: int, string, bool, channel из переменной типа interface{}

//определение типа через декларацию типов и запись типа в string
func typeSpecifier(in interface{}) string {
	switch in.(type) {
	case nil:
		return "nil"
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int, chan float32, chan string, chan bool, chan struct{}:
		return "chan"
	default:
		return fmt.Sprintf("error %T: %v", in, in)
	}
}

func main() {
	arr := []interface{}{nil, 32, "some", true, make(chan int), map[int]int{}}
	for _, val := range arr {
		fmt.Println(typeSpecifier(val))
	}
}
