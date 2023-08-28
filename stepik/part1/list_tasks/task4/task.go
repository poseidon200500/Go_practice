package main

import "fmt"

func main() {
	var n, count, tmp int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		if tmp > 0 {
			count++
		}
	}
	fmt.Println(count)
}
