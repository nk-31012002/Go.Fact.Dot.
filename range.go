package main

import "fmt"

/*

	range is used for iterating over data structures in Go

*/

func main() {
	fmt.Println("Go Start...")

	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//
	//for index, num := range arr {
	//	fmt.Println(index, "->", num)
	//}

	mp := map[string]int{"even": 1, "odd": 2}

	for k, v := range mp {
		fmt.Println(k, v)
	}

	fmt.Println("Go End.......")
}
