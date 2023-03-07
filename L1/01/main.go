package main

import "fmt"

//объявление структуры Human
type Human struct {
	name   string
	action string
}

//объявление структуры Action
type Action struct{ Human } //поле является структурой типа Human

//создаем метод связанный с типом Human
func (h *Human) Do() string {
	return h.name + h.action
}

func main() {
	//создаем экземпляр типа Human
	var Somebody Action
	Somebody.name = "hel"
	Somebody.action = "lo"
	//Все методы типа Human автоматически делаются доступными через тип Action
	Somebody.Do()
	fmt.Println("Action: ", Somebody.Do())
	fmt.Println("name:", Somebody.name)
}
