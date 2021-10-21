package helper

import "fmt"

//Access Modifier
//Uppcase at the first letter : can be called outside the `helper` package
//Lowercase at the first letter : can't be called outside package

var version = "1.0.0"      //Private Variable
var Application = "Golang" //Public Variable

//This is Public Function
func SayHello(name string) {
	fmt.Println("Hello", name)
}

//This is Private Function
func sayGoodBye(name string) {
	fmt.Println("Goodbye", name)
}
