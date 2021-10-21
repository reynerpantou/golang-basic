package main

import "fmt"

func logging() {
	fmt.Println("Function called succesfully")
}

func runApplication(value int) {

	//if some line of code getting error, than the compiler wont run the code below & stop by force
	//we can use defer to force program to schedule and run our function eventhough theres something wrong with our code
	defer logging()
	fmt.Println("Run application")

	result := 10 / value //return error if divided by 0
	fmt.Println("Result", result)
	// logging()
}

func main() {

	runApplication(0)
}
