/*
 * Created by gt on 1/21/22 - 2:59 AM.
 * Copyright (c) 2022 GTXC. All rights reserved.
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// MAPS
	// maps can store the same type of data
	// key values used in maps must be testable for equality
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862496,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)

	// make function can be used to create maps
	myMap := make(map[string]int)
	fmt.Println(myMap)

	// manipulate elements inside maps
	fmt.Println(statePopulations["Ohio"])
	statePopulations["Georgia"] = 10310371
	// order changes, return order is not guaranteed in maps
	fmt.Println(statePopulations)

	// delete entries in map
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	// querying Georgia again returns 0
	fmt.Println(statePopulations["Georgia"])

	// to make sure key is not represent in map
	pop, ok := statePopulations["Oho"] // misspelled knowing that
	fmt.Println(pop, ok)               // return still a number which is 0, but we can get also false
	// you can also use blank identifier to get just boolean returned from query
	_, ok = statePopulations["Ohio"]
	fmt.Println(ok)

	// get length of map
	fmt.Println(len(statePopulations))

	// maps are passed by reference just like slices
	// modifying map for one variable modifies the other one also
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)

	// STRUCTS
	type Doctor struct {
		number     int
		actorName  string
		companions []string
	}

	// you can initialize struct with positional syntax by removing field names - not recommended
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)

	// anonymous struct
	anEngineer := struct{ name string }{name: "G TXC"}
	fmt.Println(anEngineer)
	// struct are independent
	anotherEngineer := anEngineer
	anotherEngineer.name = "Elf G"
	fmt.Println(anEngineer)
	fmt.Println(anotherEngineer)

	// COMPOSITION - embedding
	// Bird is not a type of Animal like in traditional interface
	// Bird has Animal like characteristics
	type Animal struct {
		Name   string
		Origin string
	}

	type Bird struct {
		Animal   // embed Animal struct into Bird struct
		SpeedKPH float32
		CanFly   bool
	}

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Name)

	// tags - need reflect package
	// tags meaningless to go itself
	type AnotherAnimal struct {
		Name   string `required max: "100"`
		Origin string
	}

	t := reflect.TypeOf(AnotherAnimal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
