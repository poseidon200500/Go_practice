package main

import "fmt"

func main() {
	var start, persent, finish int
	fmt.Scan(&start, &persent, &finish)
	i := 1
	for {
		start += start * persent / 100
		if start >= finish {
			fmt.Println(i)
			break
		}
		i++
	}

}
