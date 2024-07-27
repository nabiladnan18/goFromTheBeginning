package main

import "fmt"

func main() {
	
	printMessage := true
	donotPrintMessage := false
	
	// positive
	if printMessage {
		fmt.Println("Message")
	}
	
	// if not donotPrintMessage
	if !donotPrintMessage {
		fmt.Println("Do not print message")
	}
	
	
	hasGas := true
	hasKeyInIgnition := true
	
	hasBurger := true
	hasSandwich := false
	
	accountBalance := 100
	accountCredit := 200
	
	if accountBalance + accountCredit > 0 {
		fmt.Println("Overused credit limit")
	} else {
		fmt.Printf("You have %d$ left", accountBalance)
	}
	
	testScoreGrade5 := 80
	testScoreGrade4 := 60
	testScoreGrade3 := 50
	testScore := 49
	
	if testScore >= testScoreGrade5 {
		fmt.Println("Top mark")
	} else if testScore >= testScoreGrade4 {
		fmt.Println("Pass with distinction")
	} else if testScore >= testScoreGrade3 {
		fmt.Println("Pass with distinction")
	} else if testScore < 0 {
		fmt.Println("Test score cannot be negative")
	} else {
		fmt.Println("Failed")
	}
	
	if hasGas && hasKeyInIgnition {
		fmt.Println("Can drive car")
	}
	
	if hasBurger || hasSandwich {
		fmt.Println("Can eat")
	}
	
	if !hasSandwich {
		fmt.Println("I'mma starve")
	} 
}

