package main

import "fmt"

func main() {
	var a, b, otv int
	flag := false
	fmt.Scan(&a, &b)
	for otv = b; otv >= a; otv-- {
		if otv%7 == 0 {
			flag = true
			break
		}
	}
	if flag {
		fmt.Println(otv)
	} else {
		fmt.Println("NO")
	}

}
