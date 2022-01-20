/*
 * Created by gt on 1/21/22 - 1:59 AM.
 * Copyright (c) 2022 GTXC. All rights reserved.
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ARRAY
	grades := [3]int{97, 85, 93}
	grades2 := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades: %v\n", grades2)
	fmt.Printf("Grades: %v\n", grades[1])

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	fmt.Printf("Students: %v\n", students)
	students[1] = "John"
	students[2] = "Arnold"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of the students: %v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix2)

	// copying an array creates a new copy, not the address of the first element
	a := [...]int{1, 2, 3}
	// but if you use b := &a, then b assigned as pointer points to a, and changing b changes also a
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// SLICE
	aa := []int{11, 22, 33}
	fmt.Println(aa)
	fmt.Println("Length aa: " + strconv.Itoa(len(aa)))
	fmt.Println("Capacity aa: " + strconv.Itoa(cap(aa)))
	// aa and bb points to the same array
	bb := aa
	bb[1] = 5
	fmt.Println(aa)
	fmt.Println(bb)
	// several ways to implement a slice
	// putting ... inside brackets turns aaa into array
	aaa := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bbb := aaa[:]   // slice of all elements
	ccc := aaa[3:]  // slice from 4th element to end
	ddd := aaa[:6]  // slice first 6 elements
	eee := aaa[3:6] // slice the 4th, 5th, 6th elements
	aaa[5] = 42
	fmt.Println(aaa)
	fmt.Println(bbb)
	fmt.Println(ccc)
	fmt.Println(ddd)
	fmt.Println(eee)

	// make function to create slice
	// takes 2 or 3 arguments, first type, second length, third capacity
	//aaaa := make([]int, 3)
	aaaa := make([]int, 3, 100)
	fmt.Println(aaaa)
	fmt.Printf("Length aaaa: %v\n", len(aaaa))
	fmt.Printf("Capacity aaaa: %v\n", cap(aaaa))

	z := []int{}
	fmt.Println(z)
	fmt.Printf("Length z: %v\n", len(z))
	fmt.Printf("Capacity z: %v\n", cap(z))
	z = append(z, 1)
	fmt.Println(z)
	fmt.Printf("Length z: %v\n", len(z))
	fmt.Printf("Capacity z: %v\n", cap(z))
	// append is a variadic function
	z = append(z, 2, 3, 4, 5)
	fmt.Println(z)
	fmt.Printf("Length z: %v\n", len(z))
	fmt.Printf("Capacity z: %v\n", cap(z))

	// spread operator
	// if you want to append another slice
	//z = append(z, []int{11, 22, 33}) // this does not work
	z = append(z, []int{11, 22, 33}...)
	fmt.Println(z)
	fmt.Printf("Length z: %v\n", len(z))
	fmt.Printf("Capacity z: %v\n", cap(z))

	// removing elements from the slice
	// you can use operations used to create slice
	v := z[:len(z)-1] // remove last element
	fmt.Println(v)
	u := z[1:] // remove first element
	fmt.Println(u)
	fmt.Printf("Length u: %v\n", len(u))
	fmt.Printf("Capacity u: %v\n", cap(u))
	// if you want to remove middle elements
	// by doing so original z is changed
	// make sure you do not have any other references to z
	fmt.Println(z)
	uu := append(z[:2], z[3:]...)
	fmt.Println(uu)
	fmt.Println(z)
}
