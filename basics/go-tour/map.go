package main

import "fmt"

func constructor() {
	fmt.Println("*******************")
	// nil map
	var map1 map[string]string
	fmt.Println("Is map1 Nil Map?", map1 == nil)

	// empty map
	var map2 = make(map[string]string)
	fmt.Println("Is map2 Nil Map?", map2 == nil)
	// map2["A"] = "a"

	map3 := map[string]string{
		"A": "a",
		"B": "b",
	}
	fmt.Println(map3)
}

func mutation() {
	fmt.Println("*******************")
	// insert or update a key in map
	// var mymap map[string]string // nil map
	// mymap["A"] = "a"         //  we can insert into nil map

	map1 := map[string]string{
		"A": "a",
		"Z": "p",
	}
	map1["C"] = "c" // insert
	fmt.Println("After inserting a key map is", map1)
	map1["Z"] = "z" // update
	fmt.Println("After updating a key map is", map1)

	map1["FF"] = "ff"
	delete(map1, "FF")
	fmt.Println("After deleting a key map is ", map1)

	// any element exist in map
	val, ok := map1["B"]
	fmt.Println(val, ok)
}

func iteration() {
	fmt.Println("*******************")
	mymap := map[string]string{
		"A": "a",
		"B": "b",
		"C": "c",
	}
	for key, v := range mymap {
		fmt.Println(key, v)
	}
}

func main() {
	constructor()
	mutation()
	iteration()
}
