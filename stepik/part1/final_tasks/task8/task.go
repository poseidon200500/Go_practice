package main

import "fmt"

func main() {
	var n, tmp int
	min, count := 0, 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		if i == 0 || min > tmp {
			min = tmp
			count = 1
		} else if tmp == min {
			count++
		}
	}
	fmt.Println(count)
}
