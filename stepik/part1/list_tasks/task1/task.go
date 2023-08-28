package main

import "fmt"

func main() {
	var (
		workArray     [10]uint8
		first, second uint8
	)
	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Scan(&first, &second)
		workArray[first], workArray[second] = workArray[second], workArray[first]
	}
	for _, elem := range workArray {
		fmt.Printf("%v ", elem)
	}

}
