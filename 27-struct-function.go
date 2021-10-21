package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, ", My name is", customer.Name)
}
func main() {
	reyner := Customer{
		Name:    "Reyner",
		Address: "Sidoarjo",
		Age:     22,
	}
	reyner.sayHello("Edward")
}
