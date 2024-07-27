package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func add(first int, second int) int {
	return first + second
}

func addWithErrHandling(first int, second int) int {
	// Converting two strings into ints and adding them
	first, err1 := strconv.Atoi(os.Args[1])
	second, err2 := strconv.Atoi(os.Args[2])
	
	// nil is None
	if err1 != nil {
		fmt.Println(err1)
		fmt.Println("Could not convert: " + os.Args[1])
	}
	
	if err2 != nil {
		fmt.Println(err2)
		fmt.Println("Could not convert: " + os.Args[2])
	}
	
	// still returns something
	return first + second 
	
}

func intToString(value int) string {
	return strconv.Itoa(value)
}


func assignment(value1 int, value2 int) int {	
	return value1 + value2
}

func panic(err1 error, err2 error) string {
	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		fmt.Println(err2)
	}
	return "Something went wrong!"
}

func main() {
	// fmt.Printf("%T", os.Args[1])
	arg := 1234
	fmt.Println(reflect.TypeOf(arg))
	fmt.Printf("%T\n", 1)
	fmt.Printf("%T\n", "string")
	
	no1, _ := strconv.Atoi(os.Args[1])
	no2, _ := strconv.Atoi(os.Args[2])
	
	// var sum int = add(no1, no2)
	// fmt.Println(sum)
	
	var handledSum = addWithErrHandling(no1, no2)
	fmt.Println(handledSum)
	
	var no int = 100
	fmt.Println(reflect.TypeOf(no))
	
	var intStr string = "100"
	// ParseInt(<s string>, <base int>, <bit int>) (int64, error)
	fourBaseEightBitInt, _ := strconv.ParseInt(intStr, 4, 8) // becomes no 16 and int64
	tenBaseSixteenBitInt, _ := strconv.ParseInt(intStr, 10, 16) // no 100, int64
	fmt.Println(fourBaseEightBitInt)
	fmt.Println(reflect.TypeOf(fourBaseEightBitInt))
	fmt.Println(tenBaseSixteenBitInt)
	fmt.Println(reflect.TypeOf(tenBaseSixteenBitInt))
	
	var intToStringValue = intToString(8)
	fmt.Println(intToStringValue)
	fmt.Println(reflect.TypeOf(intToStringValue))
	
	value1, err1 := strconv.Atoi(os.Args[1])
	value2, err2 := strconv.Atoi(os.Args[2])
	
	if err1 != nil || err2 != nil {
		fmt.Println(panic(err1, err2))
	} else {
		fmt.Println(assignment(value1, value2))
	}
}

