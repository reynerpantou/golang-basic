package main

import "fmt"

func main(){

	//For with while formula
	counter := 1
	for counter <= 10{
		fmt.Println("Iteration -", counter)
		counter++
	}
	fmt.Println()

	for i := 1; i <= 10; i++{
		for j :=1; j <= i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()


	slice := []string{"Reyner", "Anton", "Billy", "Edward"}
	
	//Implement slice with for loop
	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}
	fmt.Println()

	//Range Loop. Looks like ForEach in PHP and Kotlin
	for i, value := range slice{
		fmt.Println(i, "=", value)
	}
	fmt.Println()

	//Range loop without index
	//use `_` to tell the Golang that we dont need temp variable / statement to store index
	for _, value := range slice{
		fmt.Println(value)
	}
	fmt.Println()

	campaign := map[int]string{
		2016 : "Ciptakan Peluangmu",
		2017 : "Semua Dimulai dari Tokopedia",
		2018 : "Mulai Aja Dulu",
		2020 : "Selalu Ada, Selalu Bisa", 
	}

	for i, value := range campaign{
		fmt.Println(i, value)
	}
	fmt.Println()

	for _, value := range campaign{
		fmt.Println(value)
	}
}