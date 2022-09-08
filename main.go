package main

import (
	"ReadFile/example/ReadCSVFile"
	"github.com/go-co-op/gocron"
	_ "github.com/go-co-op/gocron"
	_ "ReadFile/example/ReadCSVFile"

//	"gopkg.in/robfig/cron.v2"
	_ "gopkg.in/robfig/cron.v2"

	"time"
)

func main(){
	runCronJobs()


}

func runCronJobs() {

	s := gocron.NewScheduler(time.Local)
	s.Every(20).Seconds().Do(func(){ReadCSVFile.ReadStudent()})
	s.StartBlocking()
}






