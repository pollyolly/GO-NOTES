package main
import "fmt"

func main(){
	var x int = 10
	y := 5
		if x > 5 {
			fmt.Println("Value is greater: ", x)
			goto Test
		}
		Test:
			if y < 9 {
				fmt.Println("Value is less: ", y)
			}
}
