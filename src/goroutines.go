/*
 * Created by gt on 1/22/22 - 4:55 AM.
 * Copyright (c) 2022 GTXC. All rights reserved.
 */

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	go sayHello() // prints Hello
	wg.Wait()

	var msg = "Hi"
	wg.Add(1)
	go func() { // prints Goodbye
		fmt.Println(msg)
		wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	msg = "Goodbye"       // before executing goroutine main func reaches this line and changes msg
	go func(msg string) { // prints Hi
		fmt.Println(msg)
		wg.Done()
	}(msg)
	wg.Wait()

	//time.Sleep(100 * time.Millisecond) // to give time to run goroutine
}

func sayHello() {
	fmt.Println("Hello")
	wg.Done()
}
