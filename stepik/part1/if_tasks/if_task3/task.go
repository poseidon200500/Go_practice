package main

import "fmt"

func main() {
	var ch int
	fmt.Scan(&ch)
	switch {
	case ch == 10000:
		fmt.Println(1)
	case ch > 1000:
		fmt.Println(ch / 1000)
	case ch > 100:
		fmt.Println(ch / 100)
	case ch > 10:
		fmt.Println(ch / 10)
	default:
		fmt.Println(ch)
	}
}
