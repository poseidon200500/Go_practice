package main

import "fmt"

func main() {
	var tmp int
	fmt.Scan(&tmp)
	for tmp <= 100 {
		if tmp >= 10 {
			fmt.Println(tmp)
		}
		fmt.Scan(&tmp)
	}

}
