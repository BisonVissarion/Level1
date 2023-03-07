package main

import "fmt"

type Human struct {
	name   string
	action string
}
type Action struct{ Human }

func (h *Human) Do() string {
	return h.name + h.action
}

func main() {
	Somebody := Human{action: "do"}
	Bob := Human{name: "Bob"}

	Somebody.Do()
	fmt.Println("Action: ", Somebody.Do())
	fmt.Println("name:", Bob.name)
}
