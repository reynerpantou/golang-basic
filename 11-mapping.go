package main

import "fmt"

func main(){

	person := map[string]string{
		"name" : "Reyner",
		"age" : "22",
		"status" : "Single Mingle :D",
	}

	fmt.Println(person)

	person["passion"] = "Tech"
	fmt.Println(person["passion"])
	fmt.Println(person)

	company := make(map[string]string)
	company["name"] = "GoTo"
	company["gojek"] = "Nadiem Makarim"
	company["tokopedia"] = "William Tanuwijaya"
	company["category"] = "Technology Company"

	fmt.Println(company)

	delete(company, "category")
	fmt.Println(company)
}