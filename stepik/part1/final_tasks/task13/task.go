package main

import "fmt"

func main() {
	var n int
	last, now, ind, flag := 1, 2, 3, false
	fmt.Scan(&n)
	for now = 2; now <= n; {
		if now == n {
			flag = true
			break
		}
		last, now = now, last+now
		ind++
	}
	if flag {
		fmt.Println(ind)
	} else {
		fmt.Println(-1)
	}
}
