package main

import "fmt"

func main() {
	var (
		n           int = 0
		chislo, len int
	)
	fmt.Scan(&n, &chislo)
	for len = 1; n >= len; len *= 10 {
	}
	len /= 10
	for ; len != 0; len /= 10 {
		otv := n / len % 10
		if otv != chislo {
			fmt.Printf("%d", otv)
		}
	}
}
