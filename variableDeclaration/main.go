package main

import "fmt"

var lastName string = "Topher"

var someThing string

var (
	firstName = "Chris"
	age = 31
)

func main() {

	bleh := 69

	someThing = "this is a variable that I declared before but assigned later"

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
	fmt.Println(bleh)
	fmt.Println(someThing)
	
	fmt.Printf("%s %s is %d years old", firstName, lastName, age)
	
}
