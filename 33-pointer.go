package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}
func main() {
	//var reyner Address
	reyner := Address{"Sidoarjo", "Jawa Timur", "Indonesia"}

	//Pass by Value
	//anton := reyner

	//Pass by reference
	//var antonius *Address = &reyner
	antonius := &reyner
	antonius.City = "Surabaya"

	//It will create new memory address and unlinked the old address, in this case &reyner
	antonius = &Address{"Tangerang", "Banten", "Indonesia"}

	//var edward *Address = &reyner
	edward := &reyner

	//change the value of parent variable
	//so all the data that reference to that address will change immediately

	*edward = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	//Create new pointer that reference to new memory address
	var billy *Address = &Address{"Manado", "Sulawesi", "Indonesia"}

	//Create new pointer that reference to new memory address but have no value (plain)
	fmt.Println(reyner)
	fmt.Println(antonius)
	fmt.Println(edward)
	fmt.Println(billy)

	var alamat Address = Address{
		City:     "Surabaya",
		Province: "Jawa Timur",
		Country:  "",
	}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
