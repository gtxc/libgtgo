/*
 * Created by gt on 1/21/22 - 6:23 PM.
 * Copyright (c) 2022 GTXC. All rights reserved.
 */

package main // entry point package for a go application
import "fmt"

func main() { // entry point function for a go application
	//for i := 0; i < 5; i++ {
	//	sayMessage("Hey, GT", i)
	//}
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)

	sum("The sum is", 1, 2, 3, 4, 5)
	fmt.Println(returnSomething(3))

	// multi-return example
	d, err := multipleReturn(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// anonymous function
	// can be useful in a loop
	// or isolating
	func() {
		fmt.Println("anonymous")
	}()

	f := func() {
		fmt.Println("anonymous2")
	}
	f()
	// following is the same as above
	var f1 func(msg string) = func(msg string) {
		fmt.Println(msg)
	}
	f1("anonymous3")

	g := greeter{
		greeting: "Hi",
		name:     "GT",
	}
	g.greet()
	fmt.Println("The new name is:", g.name)
	g.greet()

	gg := greeter{
		greeting: "Hello",
		name:     "gt",
	}
	gg.greet1()
	fmt.Println("The new name is:", gg.name)
	gg.greet1()
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

func sayGreeting(greeting, name *string) { // same types can be declared as this
	*name = "GT"
	fmt.Println(*greeting, *name)
}

func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

func returnSomething(num int) int {
	num = num * 2
	return num
}

func returnPointer(num int) *int {
	num = num + 1
	return &num
}

func namedReturn(value int) (result int) {
	value++
	return
}

func multipleReturn(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

// add methods onto objects
func (g greeter) greet() { // value receiver
	fmt.Println(g.greeting, g.name)
	g.name = "" // do not affect the name property
}

func (g *greeter) greet1() { // value receiver
	fmt.Println(g.greeting, g.name)
	g.name = "new name" // do not affect the name property
}
