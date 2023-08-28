package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
	last, otv := 1, 1
	if n > 2 {
		for i := 3; i <= n; i++ {
			last, otv = otv, last+otv
		}
	}
	return otv
}
