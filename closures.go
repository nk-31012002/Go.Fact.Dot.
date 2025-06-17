package main

import "fmt"

/*

	A closure is a function that can remember and use values from the place where it was created — even after that place is gone.


*/

//	func counter() func() int {
//		count := 0
//		return func() int {
//			count++
//			return count
//		}
//	}
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {

	c := counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

}
