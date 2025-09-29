package main

import "fmt"

func main(){
	fmt.Println(plus(4)) //Output: 5
}

func plus(num int) (result int) {
	if num > 3 { //Check if higher than 3
		result = num  + plus(num - 3) //Re-check this if higher than 3 if not go else return 1
	} else {
		result = 1
	}
	return
}
