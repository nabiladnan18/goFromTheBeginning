package main

import (
	"fmt"
)

func log(message string) {
	fmt.Println(message)
}

func calc(num1 int, num2 int) (sum int, product int) {
	sum = num1 + num2
	product = num1 * num2
	return
}

func main() {
	log("hello there")
	log("general kenobi")
	num1 := 2
	num2 := 3
	sum, product := calc(num1, num2)
	fmt.Printf("Sum of %d and %d = %d, Product of %d and %d = %d", num1, num2, sum, num1, num2, product)
}
