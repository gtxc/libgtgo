/*
 * Created by gt on 1/22/22 - 4:17 AM.
 * Copyright (c) 2022 GTXC. All rights reserved.
 */

package main

import "fmt"

func main() {
	var w Writer = ConsoleWriter{}
	write, err := w.Write([]byte("Hello GT"))
	if err != nil {
		return
	}
	fmt.Println("number of bytes written:", write)
}

// Writer - interfaces do not describe data, just describe behaviors
// store method definitions here
type Writer interface {
	Write([]byte) (int, error)
}

// ConsoleWriter - implement here
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
