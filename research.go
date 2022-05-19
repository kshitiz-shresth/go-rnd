package main

import (
	"fmt"
	"time"
)

func main() {
	var currentTime = time.Now()

	//basic time rules
	message("Basic Today's Date Formating")
	fmt.Println(currentTime.Format("2006 Jan 02 15:04:05")) //got surprised because it gave output of today's date

	lineBreak()
	lineBreak()
	//for custom suffix
	message("Custom Function for suffix in day of Today's Date")
	fmt.Println(formatData(currentTime))

	lineBreak()
	lineBreak()

	message("High Order Functions")
	newArr := []int{10, 20, 30, 40, 50}
	fmt.Println(newArr)
	fmt.Println("To")
	fmt.Println(reverse(newArr))
}

func reverse[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)

	for i, ele := range s {
		r[l-i-1] = ele
	}
	return r
}

func message(messageAttr string) {
	fmt.Println("*********" + messageAttr + "*********")
}

func lineBreak() {
	fmt.Println("")
}

func formatData(t time.Time) string {
	suffix := "th"
	switch t.Day() {
	case 1, 21, 31:
		suffix = "st"
	case 2, 22:
		suffix = "nd"
	case 3, 23:
		suffix = "rd"
	}
	return t.Format("Monday, 2" + suffix + " January, 2006")
}
