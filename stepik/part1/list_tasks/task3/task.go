package main

import "fmt"

func main() {
	var n, tmp int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		if i%2 == 0 {
			fmt.Printf("%v ", tmp)
		}
	}
}
