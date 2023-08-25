package main

import "fmt"

func main() {
	var lim, c, d int
	fmt.Scan(&lim, &c, &d)
	for i := 1; i < lim+1; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}
}
