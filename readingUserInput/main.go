package main

import (
	"fmt"
	"strings"
)

func noSpace() {
	// Single input - no space
	var response string
	fmt.Print("Type something without spaces: ")
	fmt.Scan(&response)
	fmt.Printf("User typed: %s\n", response)
	
}

func multiInput() {
	// multiple inputs
	var r1, r2 string
	fmt.Print("Type something with one space or press enter in between\n: ")
	fmt.Scan(&r1, &r2)
	fmt.Printf("User typed r1: %s, r2: %s\n", r1, r2)
	
	// Say we have invoices in the following pattern: INV9999.
	// The first three chars are explaining what it is i.e. the prefix.
	// The rest describes the invoice number.
	var prefix string
	var invNumber int
	
	fmt.Print("What is your invoice number?:\n")
	fmt.Scanf("%3s%d", &prefix, &invNumber)
	fmt.Printf("prefix: %s, inv: %d\n", prefix, invNumber)
}

func challenge() {
	// Challenge
	fmt.Print("How many players do you want?: ")
	var numPlayers int
	fmt.Scan(&numPlayers)
	
	var player []string
	for i := 1; i <= numPlayers; i++ {
		var playerName string
		fmt.Printf("Enter Player %d name:\n ", i)
		fmt.Scan(&playerName)
		player = append(player, playerName)
	}
	fmt.Printf("Players are:\n%s", strings.Join(player, ", \n"))
	fmt.Print("\n")
}


func main() {
	noSpace()
	multiInput()
	challenge()
	
}
