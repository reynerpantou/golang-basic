package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	//Traditional Way
	//Eventhough we doesnt set all the parameters, the variable still works like charm
	var reyner Customer
	reyner.Name = "Reyner"
	reyner.Address = "Sidoarjo"
	reyner.Age = 22

	fmt.Println(reyner.Name)
	fmt.Println(reyner.Address)
	fmt.Println(reyner.Age)

	//Best Way - more clean
	//Eventhough we doesnt set all the parameters, the variable still works like charm
	anton := Customer{
		Name:    "Antonius",
		Address: "Tangerang",
		Age:     22,
	}
	fmt.Println(anton.Name)
	fmt.Println(anton.Address)
	fmt.Println(anton.Age)

	//Not Recommeded Way
	//Return error if doesnt fulfill all the parameters from struct
	edward := Customer{"Edward", "Surabaya", 22}
	fmt.Println(edward.Name)
	fmt.Println(edward.Address)
	fmt.Println(edward.Age)
}
