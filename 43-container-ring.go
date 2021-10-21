package main

import (
	"fmt"
	"container/ring"
	"strconv"
)

func main(){
	data := ring.New(5)
	
	for i := 0; i <= data.Len(); i++ {
		data.Value = "Data " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}