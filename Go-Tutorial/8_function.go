package main

import "fmt"

 func Operate(a, b int, AnotherOperation func(int, int) int) int {
      return AnotherOperation(a, b)
 }

func main(){
	add := func(x, y int) int { return x + y }
	result := Operate(3, 4, add) // 7
	fmt.Println(result)
}
