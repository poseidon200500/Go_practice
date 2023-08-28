package main

import "fmt"

func main() {
	var ch, ch1, ch2, ch3 int
	fmt.Scan(&ch)
	ch1 = ch / 100
	ch2 = ch % 100 / 10
	ch3 = ch % 10
	if ch1 != ch2 && ch1 != ch3 && ch2 != ch3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
