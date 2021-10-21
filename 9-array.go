package main

import "fmt"

func main(){
	
	var names [3]string
	names[0] = "Reyner"
	names[1] = "Anton"
	names[2] = "Edward"
	
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//Out of bounds
	//Error Message : invalid array index 3 (out of bounds for 3-element array)
	//fmt.Println(names[3]) 

	var values = [5]int{
		90,
		90,
		80,
		60,
		100,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println("Values Count = ", len(values))
	
	//The output should be 10 eventhough array still empty
	//because len basically count the maximum length of bound
	var list [10]string
	fmt.Println("List Count =",len(list))
}