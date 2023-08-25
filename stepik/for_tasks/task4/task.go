package main

import "fmt"

func main() {
	var tmp, count, max int = 0, 0, 0

	for fmt.Scan(&tmp); tmp != 0; fmt.Scan(&tmp) {
		if tmp > max {
			count = 1
			max = tmp
		} else if tmp == max {
			count++
		}
	}
	fmt.Println(count)
}
