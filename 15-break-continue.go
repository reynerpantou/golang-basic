package main

import "fmt"

func main(){

	for i := 1; i <= 15; i++{
		if i % 10 == 0{
			break			
		}else{
			fmt.Println(i)
		}
	}
	fmt.Println()
	
	for i := 1; i <= 15; i++{
		if i % 2 == 0{
			continue
		}
		fmt.Println(i)
	}
}