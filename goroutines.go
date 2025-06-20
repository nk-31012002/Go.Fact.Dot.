package main

import "fmt"

/*

A goroutine is a lightweight thread managed by the Go runtime.
It lets you run a function at the same time as the rest of your program.

Think of it like starting a new task without waiting for the previous one to finish.

if i want to run things concurrently or want to use multithreading we use ---> goroutines.

*/

func solve1(t int) {
	fmt.Println("Goroutines", t)
}

func main() {
	fmt.Println("Go Starts")

	for i := 0; i < 10; i++ {
		solve1(i) // ---> what this will do, this will print each line sequentially ---> like first, then second and so on
	}

	fmt.Println("Go Ends")
}
