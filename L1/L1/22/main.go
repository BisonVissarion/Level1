package main

import (
	"fmt"
	"math/big"
)

func main() {
	//инициализируем *big.Int
	a, _ := new(big.Int).SetString("111111111111111111111111111111111", 10)
	b, _ := new(big.Int).SetString("999999999999999999999999999999999", 10)
	fmt.Printf("%v + %v = %v\n", a, b, new(big.Int).Add(a, b)) // сложение двух типов *big.Int
	fmt.Printf("%v - %v = %v\n", a, b, new(big.Int).Sub(a, b)) // вычетание двух типов *big.Int
	fmt.Printf("%v * %v = %v\n", a, b, new(big.Int).Mul(a, b)) // умножение двух типов *big.Int
	//преобразовываем *big.Int в *big.Float
	fmt.Printf("%v / %v = %v\n", a, b, new(big.Float).Quo(new(big.Float).SetInt(a), new(big.Float).SetInt(b))) // деление двух типов *big.Float

}
