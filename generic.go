package main

import "fmt"

/*

generic :- Generics allow you to write functions that can work with any data types.


we can use -> any or interface{}

if i want to scope the genetic say---> only 2 data types should be used for that we can use ---> data type1 | data type2 | data type3.....so on


*/

//func printSlices[T any](item []T) {
//	for _, item := range item {
//		fmt.Println(item)
//	}
//}

//func printSlices[T int | string](item []T) {
//	for _, item := range item {
//		fmt.Println(item)
//	}
//}

//func printSlices[T interface{}](item []T) {
//	for _, item := range item {
//		fmt.Println(item)
//	}
//}

//func printSlicesString(item []string) {
//	for _, item := range item {
//		fmt.Println(item)
//	}
//}

func main() {
	fmt.Println("Go Starts")
	//printSlices([]int{1, 2, 3, 4})
	//printSlices([]string{"one", "two", "three"})
	fmt.Println("Go Ends")
}
