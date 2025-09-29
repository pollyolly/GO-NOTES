package main

import (
	"fmt"
	"example.com/greetings_module/greetings"
	"example.com/greetings_module/welcome"
	"example.com/hello_module/testing"
)

func main(){
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	welcome.Welcome()
	testing.Testme()
}
