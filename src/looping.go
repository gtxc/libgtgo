/*
 * Created by gt on 1/21/22 - 4:33 AM.
 * Copyright (c) 2022 GTXC. All rights reserved.
 */

package main

import (
	"fmt"
)

func main() {
	// simple loops (++i does not work)
	//for i := 0, j := 0; i < 5; i++, j++ { // there is no comma operator in go
	// increment operator is a statement itself
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 { // initialize multiple variables
		fmt.Println(i, j)
	}
	fmt.Println()
	// manipulating the counter
	for a := 0; a < 5; a++ {
		fmt.Println(a)
		if a%2 == 0 {
			a /= 2
		} else {
			a = 2*a + a
		}
	}
	fmt.Println()
	// following works too
	b := 0 // changing scope
	for ; b < 5; b++ {
		fmt.Println(b)
	}
	// do not need incrementer value
	c := 0
	for c < 5 {
		fmt.Println(c)
		c++
	}
	fmt.Println()
	// infinite for loop
	for {
		fmt.Println(c)
		if c == 5 {
			break // the same continue keyword exist
		}
	}
	// to break nested loops
Loop:
	for {
		for {
			if true {
				break Loop
			}
		}
	}

	// looping through collections
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v) // k's are the indexes, v's are the values
	}

	// looping over maps
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862496,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	// string as a source for for range loop
	for k, v := range "this is a string" {
		fmt.Println(k, string(v))
	}

	// for loops with channels
}
