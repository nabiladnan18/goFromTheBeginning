package main

import "fmt"

func loopThisManyTimes(times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("Printing: %d\n", i)
	}
}

func while(value int) {
	fmt.Printf("Print %d times\n", value)
	fmt.Printf("But will exit on the 10th time\n")
	
	// Go does not have the while keyword
	i := 1
	for i < value {
		if i == 5 {
			fmt.Printf("Skip 5\n")
			i ++
			continue
		} else if i == 10 {
			fmt.Printf("Printed 10 times. Exiting...\n")
			break
		} else {
			fmt.Printf("Printing: %d\n", i)
			i ++
		}
	}
}

func main() {
	
	// Classic for-loop
	times := 2
	fmt.Printf("Looping %d times\n", times)
	loopThisManyTimes(times)
	
	// "while-loop". Note that Go does not have the "while" keyword
	value := 11
	while(value)
	
	// idk types of data yet; i cannot strictly define the data type "array" for a func 🤷‍♂️
	// Anyhoo
	// Looping over an iterable like an Array
	arr := []int{-1, 2, 3}
	for i := 0; i < 2; i++ {
		fmt.Println(arr[i])
		if arr[i] < 0 {
			break
		}
	}
	
	var keepGoing = true
	answer := ""
	for keepGoing {
		fmt.Println("Type command: ")
		fmt.Scan(&answer)
		if answer == "quit" {
			keepGoing = false
		} else if answer == "print" {
			fmt.Println("Printing file.")
			break
		}
	}
	fmt.Println("program exit")
	
}