package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("time module of golang")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// if we want to create our own datetime, we do the following:
	createdDate := time.Date(2025, time.April, 15, 22, 22, 0, 0, time.UTC)  // here only the month is of type time.Month, rest everything is of type int
    // time.UTC is for location which is of format time.Location
    // location basically tells us in what timezone we are
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}