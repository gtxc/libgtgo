/*
 * Created by gt on 1/22/22 - 6:00 AM.
 * Copyright (c) 2022 GTXC. All rights reserved.
 */

// designed to synchronize data transfer between multiple goroutines
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // signal only channel, 0 memory allocation required

func main() {
	ch := make(chan int, 50) // create a channel can store 50 integers
	wg.Add(2)
	go func(ch <-chan int) { // receiving goroutine
		//for i := range ch {
		//	fmt.Println(i)
		//}
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // sending goroutine
		ch <- 42
		ch <- 12
		ch <- 14
		ch <- 16
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()

	//----
	go logger()
	//defer func() {
	//	close(logCh)
	//}()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}

func logger() {
	//for entry := range logCh {
	//	fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	//}
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}
