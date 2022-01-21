/*
 * Created by gt on 1/21/22 - 5:02 AM.
 * Copyright (c) 2022 GTXC. All rights reserved.
 */

package main

import (
	"fmt"
	"log"
)

func main() {
	//// DEFER - closing resources
	//// defers in LIFO order
	//fmt.Println("start")
	//defer fmt.Println("middle") // defers the line until the main function returns
	//fmt.Println("end")
	//// more realistic example to defer
	//res, err := http.Get("http://www.google.com/robots.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer res.Body.Close() // to not forget to close resource defer it and close after open it
	//robots, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%s", robots)
	//// another defer example
	//a := "start"
	//defer fmt.Println(a) // "start" will printed
	//a = "end"
	//
	//// PANIC
	////q, w := 1, 0
	////ans := q / w // compiler generates panic
	////fmt.Println(ans)
	//// create panic manually
	//fmt.Println("start")
	////panic("something bad happened")
	//fmt.Println("end")
	//// web handler example for panic
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello GT!"))
	//})
	//errr := http.ListenAndServe(":8080", nil)
	//if errr != nil {
	//	panic(errr.Error())
	//}

	// RECOVER
	// panics happen after defer statements are executed
	fmt.Println("start")
	// anonymous function
	//defer func() {
	//	if err := recover(); err != nil {
	//		log.Println("Error:", err)
	//	}
	//}()
	//panic("something bad happened")
	//fmt.Println("end")
	panicker()
	fmt.Println("end") // re-panicing prevents to execute this line

}
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic(err) // re-panic the application
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
