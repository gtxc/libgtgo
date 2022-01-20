/*
 * Created by gt on 1/20/22 - 12:48 PM.
 * Copyright (c) 2022 GTXC. All rights reserved.
 */

package main

import (
	"fmt"
)

// all un-initialize variables in go assigned to 0
func main() {
	// boolean type
	var n bool = true
	n = 1 == 1
	m := 2 == 2
	fmt.Printf("var: %v, type: %T\n", n, n)
	fmt.Printf("var: %v, type: %T\n", m, m)
	// numeric types
	// int8, int16, int32, int64
	// uint8, uint16, uint32
	// byte
	// arithmetic operations on different type not allowed (e.g. int8 + int 32)
	// to make it work, use type-casting
	// bit shifting applies as the same as C
	a := 8              // 2^3
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0
	// floating point numbers
	f := 3.14
	f = 13.7e72
	f = 2.1e14
	fmt.Printf("%v, %T\n", f, f)
	// complex numbers - standard arithmetic operations apply
	var c complex64 = 1 + 2i
	var cc complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", real(c), real(c))
	fmt.Printf("%v, %T\n", imag(c), imag(c))
	fmt.Printf("%v, %T\n", imag(cc), imag(cc))
	// text type - string concatenation with + applies - utf8 only
	s := "this is a string"
	fmt.Printf("%v, %T\n", s[2], s[2])
	// string to collection of bytes
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
	// rune - utf32
	//var r rune = 'a' // same as below
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)
}
