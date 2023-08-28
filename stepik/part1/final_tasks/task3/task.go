package main

import "fmt"

func main() {
	var sec, minutes, hours int
	fmt.Scan(&sec)
	hours = sec / 3600
	minutes = sec % 3600 / 60
	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
