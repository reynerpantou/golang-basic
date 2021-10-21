package main

import (
	"fmt"
	"strings"
)

func main(){
	
	fmt.Println(strings.Contains("Reyner Pantou", "Rey"))
	fmt.Println(strings.Contains("Reyner Pantou", "Edward"))

	fmt.Println(strings.ToLower("Reyner Pantou"))
	fmt.Println(strings.ToUpper("Reyner Pantou"))
	fmt.Println(strings.ToTitle("Reyner Pantou"))

	fmt.Println(strings.Trim("               Reyner Pantou        ", " "))
	fmt.Println(strings.ReplaceAll("Pantou Pan Pantou", "Pantou", "Rey"))
} 