/*
 * Created by gt on 1/20/22 - 12:08 PM.
 * Copyright (c) 2022 GTXC. All rights reserved.
 */

package main

import (
	"fmt"
	"strconv"
)

// outside the function, in package level, cannot use :=, must use full declaration syntax
var gt float32 = 123

// block of variable declaration
// by doing so, we say that those variables are related to each other
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

var (
	counter int  = 0
	flag    bool = false
	// we can declare i variable here also
	// i declared in main() shadows the package level declaration
	i int = 13
)

func main() {
	// type declared after variable name
	var i int = 42
	i = 47
	// by doing so, compiler decides the type
	j := 12
	fmt.Println(i)
	fmt.Println(j)
	k := 99.
	fmt.Printf("var: %v, type: %T\n", k, k)
	fmt.Println(gt)
	// type conversion
	k = float64(j)
	i = int(k)

	i = 42
	var str string
	// str takes the ascii value of i
	str = string(i)
	fmt.Printf("var: %v, type: %T\n", str, str)
	// import strconv package
	str = strconv.Itoa(i)
	fmt.Printf("var: %v, type: %T\n", str, str)
}
