/*
 * Created by gt on 1/22/22 - 5:11 AM.
 * Copyright (c) 2022 GTXC. All rights reserved.
 *
 * goroutines are racing each other
 */

// go run - race src/main.go

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} // to prevent simultaneous reading writing

func main() {
	fmt.Printf("Available by default Threads: %v\n", runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	//m.RLock()
	fmt.Printf("hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	//m.RLock()
	counter++
	m.Unlock()
	wg.Done()
}
