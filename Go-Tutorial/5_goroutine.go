package main

import ("fmt")


func main(){
   go display("Process 1")
   display("Process 2")
}

func display(s string){
	fmt.Println(s)
}
