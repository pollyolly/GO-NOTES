package main

import (
	"fmt"
)

func main(){
	names := []string{"John Cena", "Katy Perry", "William Roberts", "Jullian"}
	for i:=0; i < len(names); i++ {
		if names[i] == "Katy Perry" {
			continue //Use to Skip one or more iteration of the Loop. "Katy Perry" will no be printed)
		}
		if names[i] == "William Roberts" {
			break   //Use to Break/Terminate the Loop 
		}
		fmt.Println(names[i])
	}
	numbers := []int{56 ,78, 90}
	for index, val := range numbers {
		fmt.Printf("%v\t%v\n", index, val)
	}
	//Output:
	/* John Cena
	0	56
	1	78
	2	90 */
	for {
       fmt.Println("Hello")
	}
}
