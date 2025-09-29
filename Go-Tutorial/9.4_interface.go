package main

import (
	"fmt"
	"math"
)

type operation interface {
	addition() float64
	subtraction() float64
}
type dataType1 struct {
	num1, num2 float64
}
type dataType2 struct {
	num3 float64
}
//======= Required to implement the struct to the two function in interface
func (p dataType1) addition() float64 {
	return p.num1 + p.num2
}
func (p dataType1) subtraction() float64 {
	return math.Pi * p.num1 - p.num2
}
//===== Required to implement the struct to the two function in interface
func (s dataType2) subtraction() float64 {
	return math.Pi - s.num3
}
func (s dataType2) addition() float64 {
	return s.num3 - 20
}
//=====
func process(o operation){
	fmt.Println(o)
	fmt.Println(o.addition())
	fmt.Println(o.subtraction())
}
func main(){
	p := dataType1{num1: 4, num2: 6}
	s := dataType2{num3: 45.2}
	process(p)
	process(s)
}
