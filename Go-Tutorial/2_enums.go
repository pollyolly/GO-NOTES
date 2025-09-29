package main

import "fmt"

type Status int

const (
	StatusActive Status = iota
	StatusInPending
	StatusInactive
)

func (s Status) String() string {
	switch s {
		case StatusActive:
			return "Active"
		case StatusInPending:
			return "Pending"
		case StatusInactive:
			return "Inactive"
		default:
			return "Unknown"
	}
}
func main(){
	taskStatus := StatusInPending
	fmt.Printf("Task Status: %v\n", taskStatus) //Output:Task Status: Pending
	if taskStatus == StatusInactive {
		fmt.Println("Task is inactive")
	} else {
		fmt.Println("Task is active")
	}
}
