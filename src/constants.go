/*
 * Created by gt on 1/21/22 - 12:21 AM.
 * Copyright (c) 2022 GTXC. All rights reserved.
 */

package main

import (
	"fmt"
	"unsafe"
)

// shadowed by inner declaration in main
const a int16 = 27

// enumerated constants
// can assign just x to iota only
const (
	x = iota
	y = iota
	z = iota
)

// resets counter to 0, x2 : 0
const (
	// use first value as error checking
	// since un-initialize variables assigned to 0
	// can also use _ (underscore) for the first element
	// can also add some number to give an offset
	_          = iota + 5
	errorValue = iota
	x2
)

const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)
	// cannot assign a value is determined in runtime
	//const mConst float64 = math.Sin(1.57)
	//fmt.Printf("%v, %T\n", mConst, mConst)
	const a = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Println(unsafe.Sizeof(a))

	fmt.Printf("%v, %T\n", x, x) // 0
	fmt.Printf("%v, %T\n", y, y) // 1
	fmt.Printf("%v, %T\n", z, z) // 2

	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	fmt.Printf("%v, %T\n", TB, TB)
}
