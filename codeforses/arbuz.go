package main

import "fmt"

func main() {
	var height int
	fmt.Scan(&height)
	if height%2 == 0 && height > 2 {
		fmt.Printf("YES")
	} else {
		fmt.Printf("NO")
	}
}
