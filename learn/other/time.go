package main

import (
	"time"
	"fmt"
)

const DATE_FORMAT = "2006-01-02 10:10:10"

func main() {
	t := time.Now()
	fmt.Println(t.YearDay())
	fmt.Println(t.Year())
	fmt.Println(t.Day())
	fmt.Println(t.Month().String())

	year := time.Now().Year()
	month := time.Now().Month() //time.Now().Month().String()
	day := time.Now().Day()

	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)

	timeStr := time.Now().Format("2006-01-02")
	fmt.Println("timeStr:", timeStr)
	t1, _ := time.Parse("2006-01-02", timeStr)
	timeNumber := t1.Unix()
	fmt.Println("timeNumber:", timeNumber)

}
