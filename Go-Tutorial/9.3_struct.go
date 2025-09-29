package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person {name: name}
	p.age = 42
	return &p
}

func main(){
	fmt.Println(person{"Bob", 34})
	fmt.Println(person{name: "Thamuz"})
	fmt.Println(newPerson("Hilda"))

	s := person{name:"leomord", age: 40}
	fmt.Println(s.name)

	sp := &s 
	fmt.Println(sp.age)
}
