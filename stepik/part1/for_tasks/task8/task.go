package main

import "fmt"

func main() {
	var num1, num2 int
	otv := 1
	fmt.Scan(&num1, &num2)
	for i := 10; i < num1*10; i *= 10 {
		first := num1 % i / (i / 10)
		for j := 10; j < num2*10; j *= 10 {
			second := num2 % j / (j / 10)
			if first == second {
				if otv == 1 {
					otv *= first
				} else {
					otv = otv*10 + first
				}
				break
			}
		}
	}
	for i := 10; i < otv*10; i *= 10 {
		num := otv % i / (i / 10)
		fmt.Print(num, " ")
	}

}
