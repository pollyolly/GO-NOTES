package main

import "fmt"

func main(){
	value := 42
	pointer := &value

	fmt.Printf("Value at address: %v\n", *pointer)

	*pointer = 100

	fmt.Printf("New value:\t%v\n", value)
}
