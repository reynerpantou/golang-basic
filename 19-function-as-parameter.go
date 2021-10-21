package main

import "fmt"

//best way to implement function as parameter if we have alot of parameters
type Filter func(string) string //its type declaration
func sayHelloWithFilter(name string, filter Filter) string {

	//we can use this pattern if the parameter is not much < 5 maybe(?)
	//func sayHelloWithFilter(name string, filter func(string) string) string {
	filteredName := filter(name)
	return "Hello, " + filteredName
}

func spamFilter(name string) string {
	if name == "Nakama" {
		return "..."
	} else {
		return name
	}
}
func main() {

	//We can save function as a value, but not only that
	//We can use function as parameter too

	greeting := sayHelloWithFilter("Reyner", spamFilter)
	fmt.Println(greeting)

	greeting = sayHelloWithFilter("Nakama", spamFilter)
	fmt.Println(greeting)
}
