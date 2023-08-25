package main

import "fmt"

func main() {
	var degrees, hours, minutes int
	fmt.Scan(&degrees)
	hours = degrees / 30
	minutes = degrees % 30 * 2
	fmt.Println("It is", hours, "hours", minutes, "minutes.")
}
