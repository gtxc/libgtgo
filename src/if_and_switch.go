/*
 * Created by gt on 1/21/22 - 4:03 AM.
 * Copyright (c) 2022 GTXC. All rights reserved.
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	if true {
		fmt.Println("The test is true")
	}

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862496,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	// after if until ; we have initializer which allows us to run a statement to generate
	// comparable variable
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	// standard && and || operators apply
	// NOT operator is !

	// floating point numbers are the approximation of the decimal numbers
	// following may act different in each run
	//myNum := 0.1
	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}

	// to overcome make if condition like
	//math.Abs(myNum / math.Pow(math.Sqrt(myNum), 2) - 1) < 0.001
	// by doing so, you can add as many decimal places you want and tune the threshold accordingly

	// switch statement in C applies as the same
	//switch tag { case 1: ...}
	// like in if block, you can use initializer syntax to evaluate a value before switch it
	// different ways to create switch blocks
	// case statements are compares
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // use to execute next case even if test case false
	case i >= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// type switch
	var ii interface{} = [3]int{}
	switch ii.(type) {
	case int:
		fmt.Println("ii is an int")
	case float64:
		fmt.Println("ii is an float64")
	case string:
		fmt.Println("ii is an string")
	case [2]int:
		fmt.Println("ii is a [2]int")
	case [3]int:
		fmt.Println("ii is a [3]int")
		//break // to exit switch early
		fmt.Println("This will print too")
	default:
		fmt.Println("ii is another type")
	}
}
