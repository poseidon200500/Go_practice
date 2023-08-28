package main

import "fmt"

func main() {
	//var n int
	//fmt.Scan(&n)
	fmt.Println(sumInt(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
}

func sumInt(args ...int) (int, int) {
	var sum, count int = 0, 0
	for _, elem := range args {
		sum += elem
		count++
	}
	return count, sum
}
