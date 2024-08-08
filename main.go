package main

import (
	"time"
	"timesheet-app/app"
)

func main() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	time.Local = loc
	app.RunService()
}
