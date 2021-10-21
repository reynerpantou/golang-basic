package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("in Method :", man.Name)
}

func main() {
	person := Man{"Reyner"}
	person.Married()
	fmt.Println(person)
}
