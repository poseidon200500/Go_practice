package main

import "fmt"

func main() {
	var ch int
	fmt.Scan(&ch)
	ch = ch%10*100 + ch/10%10*10 + ch/100
	fmt.Println(ch)
}
