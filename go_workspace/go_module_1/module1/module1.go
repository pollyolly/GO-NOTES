package main

import (
	"fmt"
	"example.com/go_module_2/module2"
)
func Module1(){
	fmt.Println("Module 1")
}
func main(){
	module2.Module2()
	Module1()
}
