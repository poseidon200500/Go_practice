package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%10 == 1 && n%100 != 11 {
		fmt.Println(n, "korova")
	} else if n%10 >= 2 && n%10 <= 4 && n%100 != 12 && n%100 != 13 && n%100 != 14 {
		fmt.Println(n, "korovy")
	} else {
		fmt.Println(n, "korov")
	}
}
