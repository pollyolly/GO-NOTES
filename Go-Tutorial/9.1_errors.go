package main

import (
	"fmt"
	"errors"
)

func testError(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Cant work with 42")
	}
	return arg + 3, nil
}

func main(){
	for _, i := range []int{47, 42} {
		if val, err := testError(i); err != nil {
			fmt.Println("failed:", err)
		} else {
			fmt.Println("worked:", val)
		}
	}
}
