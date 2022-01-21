/*
 * Created by gt on 1/21/22 - 3:51 PM.
 * Copyright (c) 2022 GTXC. All rights reserved.
 */

package main

import "fmt"

func main() {
	a := 42
	b := a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)

	var aa int = 42
	var bb *int = &aa
	fmt.Println(aa, bb)
	fmt.Println(aa, *bb)
	aa = 27
	fmt.Println(aa, bb)
	fmt.Println(aa, *bb)
	*bb = 14
	fmt.Println(aa, *bb)
	// pointer arithmetic is the same as in C except
	// cannot subtract int from *int e.g. bb - 4
	// if you need to do these type of operations use unsafe package

	type myStruct struct {
		foo int
	}
	var ms *myStruct
	ms = &myStruct{foo: 42}
	ms2 := new(myStruct)
	var ms3 *myStruct
	fmt.Println(ms)
	fmt.Println(ms2)
	fmt.Println(ms3) // nil
	(*ms).foo = 14
	fmt.Println((*ms).foo)
	ms.foo = 15
	fmt.Println(ms.foo) // pointer itself has a field foo

	fmt.Println()

	x := [3]int{1, 2, 3}
	y := x
	fmt.Println(x, y)
	x[1] = 42
	fmt.Println(x, y)

	fmt.Println()

	xx := []int{1, 2, 3}
	yy := xx
	fmt.Println(xx, yy)
	xx[1] = 42
	fmt.Println(xx, yy)
	// maps also have the same behavior as slices
}
