package main

import "fmt"

func main() {
	var count, sum, tmp int
	sum = 0
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		fmt.Scan(&tmp)
		if 100 > tmp && tmp > 9 && tmp%8 == 0 {
			sum += tmp
		}
	}
	fmt.Println(sum)
}
