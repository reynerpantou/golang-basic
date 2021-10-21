package main

import "fmt"

func getName(name string) string {
	return "Hi, my name is " + name
}

func main() {
	account := getName

	result := account("Reyner")
	fmt.Println(result)

}
