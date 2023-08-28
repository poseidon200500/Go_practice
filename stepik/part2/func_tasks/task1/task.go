package main

import "fmt"

func main() {
	fmt.Println(minimumFromFour())
}

func minimumFromFour() int {
	var a1, a2, a3, a4 int
	fmt.Scan(&a1, &a2, &a3, &a4)
	return min(a1, a2, a3, a4)
}
