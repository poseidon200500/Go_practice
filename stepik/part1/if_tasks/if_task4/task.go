package main

import "fmt"

func main() {
	var bilet int
	fmt.Scan(&bilet)
	part1 := bilet / 1000
	part2 := bilet % 1000
	sum1 := part1/100 + part1/10%10 + part1%10
	sum2 := part2/100 + part2/10%10 + part2%10
	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
