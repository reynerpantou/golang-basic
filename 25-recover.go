package main

import "fmt"

func endApp() {
	message := recover()

	if message != nil {
		fmt.Println("Error message :", message)
	}

	fmt.Println("Application closed")

}

func runApp(error bool) {
	defer endApp()
	if error == true {

		panic("Application error")
	}

	fmt.Println("Application running")

}

func main() {
	runApp(false)

	//This code below wont run if error / panic triggered ("Reyner")
	//We can use recover to get the panic message & it will not stop the program
	//with recover, the code below will not be skipped
	fmt.Println("Reyner")

	//comment 6th - 10th line of code to show the different between used / without recover
	//and set 28th code to true, does `Reyner` keep showing up?

}
