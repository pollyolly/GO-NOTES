package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main(){
	fmt.Println("Start CronJob")
	//Create new cronjob
	c := cron.New()
	//Add cron schedule and task to cronjob
	c.AddFunc("@every 00h00m01s", printThis)
	//Start cron job
	c.Start()

	//Keep main program running
	select {}
}

func printThis(){
	fmt.Println("Message after 1 second")
}
