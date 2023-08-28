package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	for !(n > 0 && n < 10) {

		tmp_sum := 0
		for i := 1; i <= n; i *= 10 {
			tmp_sum += (n / i) % 10
		}
		n = tmp_sum
	}
	fmt.Println(n)
}
