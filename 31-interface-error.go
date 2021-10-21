package main

import (
	"errors"
	"fmt"
)

func Divide(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Divider is zero!")
	} else {
		return value / divider, nil
	}
}
func main() {
	result, err := Divide(90, 0)
	if err == nil {
		fmt.Println("Result :", result)
	} else {
		fmt.Println("Error :", err.Error())
	}
}
