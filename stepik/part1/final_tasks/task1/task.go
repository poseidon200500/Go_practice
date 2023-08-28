package main

import "fmt"

func main() {
	var ch int
	fmt.Scan(&ch)
	sum := ch/100 + ch/10%10 + ch%10
	fmt.Println(sum)
}
