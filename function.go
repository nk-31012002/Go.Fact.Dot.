package main

import "fmt"

/*

func function_name(argument_type value, ... ) return_type {
 	//body
}


-------> In Go, there can be multiple return values.

func function_name(argument_type, ... ) (return_type1,return_type2, ...) {
	//body
}

*/

func add(x int, y int) int {
	return x + y
}

func give() (string, string) {
	return "hello", "world"
}

func main() {

	//ans := add(1, 2)
	//fmt.Println(ans)

	x, y := give()

	fmt.Println(x)
	fmt.Println(y)

}
