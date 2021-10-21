package main

import "fmt"

func main() {
	counter := 0
	name := "Reyner"
	increment := func() {
		name = "Nakama" //Reyner changed to Nakama -> its using existing variable
		name := "Admin" //Wont replace Reyner, because its create a new variable instead replacing the existing variable
		fmt.Println("Counter :", counter, ";", name)
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)

	//Its changed to Nakama according to the 9th line of code
	fmt.Println(name)
}
