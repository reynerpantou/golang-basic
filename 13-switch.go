package main

import "fmt"

func main(){
	
	name := "Reyners"

	switch name {
	case "Reyner":
		fmt.Println("Hi Reyner")
	case "Pantou":
		fmt.Println("Hi Pantou")
	default:
		fmt.Println("Can we be friend?")
	}

	//Switch with short statement
	switch length := len(name); length > 5{
	case true:
		fmt.Println("Correct")
	case false:
		fmt.Print("Incorrect")
	}

	//Switch with no condition
	grade := 95
	switch{
	case grade > 85 && grade < 100:
		fmt.Println("A")
	case grade > 75:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}
}