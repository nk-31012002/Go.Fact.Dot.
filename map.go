package main

import (
	"fmt"
)

/*
	maps -> hash, object, dict
*/

func main() {
	fmt.Println("Go Start...")

	////creating map
	//m := make(map[string]string)
	//
	////setting an element in map m
	//m["Lang"] = "Go"
	//m["Lang1"] = "CPP"
	//
	////getting an element
	//fmt.Println(m["Lang"], m["Lang1"])

	/*
		Note : If key does not exist in the map then it return zero value --------------> important
		string --> nil
		int --> 0
		bool --> false
		so on....
	*/

	//mp := make(map[int]bool)
	//mp[1] = true
	//
	//fmt.Println(len(mp))
	////delete(mp, 1)
	//clear(mp)
	//fmt.Println(len(mp))

	//other way of creating and adding element in the map
	//mp := map[string]int{"one": 1, "two": 2, "three": 3}
	//
	//fmt.Println(mp)
	//
	//k, ok := mp["two"]
	//fmt.Println(k)
	//if ok {
	//	fmt.Println("all okay")
	//} else {
	//	fmt.Println("not okay")
	//}

	//maps packages

	//m1 := map[int]bool{1: true, 2: false, 3: true}
	//m2 := map[int]bool{1: true, 2: false, 3: false}
	//
	//fmt.Println(maps.Equal(m1, m2))
	//
	fmt.Println("Go End.......")
}
