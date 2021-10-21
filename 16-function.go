package main

import "fmt"

func display(year int) string {

	var campaign string
	if year == 2017 {
		campaign = "Mulai Aja Dulu"
	} else {
		campaign = "Selalu Ada Selalu Bisa"
	}
	return campaign
}

func getData() (string, int) {
	return "Reyner", 22
}

//Named return function
func namedReturn() (_company, _founder string) {
	_company = "Tokopedia"
	_founder = "William Tanuwijaya"

	//Best way
	return

	//The other way
	//return _company, _founder
}

func main() {

	fmt.Println(display(2017))

	//Return multiple value
	name, age := getData()
	fmt.Println(name, ":", age)

	//Return specific value on multiple values function
	//using "_" to bypass the parameter which is required at multiple values function
	firstName, _ := getData()
	fmt.Println(firstName)

	//It works! Even the variables name below is totally different to the named functionn
	company, founder := namedReturn()
	fmt.Println("Company :", company, ", Founder", founder)
}
