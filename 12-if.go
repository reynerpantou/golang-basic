package main

import "fmt"

func main(){

	nama := "Reyner Pantou"

	if nama == "Reyner"{
		fmt.Println("Hello Reyner")
	} else if nama == "Billy"{ 
		fmt.Println("Hello Billy")	
	}else{
		fmt.Println("Can we be friend?")
	}

	//If with short statement
	if length := len(nama); length < 5{
		fmt.Println("The length must greater than 5")
	}else{
		fmt.Println("Correct")
	}
}