package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"Name": name,
		}
	}
}

func main() {
	person := NewMap("Reyner")

	//nil is the convenient way.
	//it can be used for interface, function, map, slice, pointer, & channel
	//without nil, we must check is some value is empty of stirng
	//e.g : person["name"] == "" -> it takes a lot energy :)
	if person == nil {

		fmt.Println("Empty Data")
	} else {

		fmt.Println(person)
	}

}
