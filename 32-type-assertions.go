package main

import "fmt"

func random() interface{} {
	return 100
}

func main() {
	//set Empty interface to variable
	result := random()

	//Convert empty interface to string data type
	// resultString := result.(string)
	// fmt.Println(resultString)

	//Panic will be triggered because random() only return string, not int
	//resultInt := result.(int)

	//if panic showing up and theres no recover written, then the program will end immediately
	//in secure way, we can use switch expression for type assertion
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown Type")
	}
}
