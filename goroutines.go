package main

import (
	"fmt"
	"sync"
)

/*

A goroutine is a lightweight thread managed by the Go runtime.
It lets you run a function at the same time as the rest of your program.

Think of it like starting a new task without waiting for the previous one to finish.

if i want to run things concurrently or want to use multithreading we use ---> goroutines.

important --> so in go, if we use goroutines then it not run like traditional function. It first go to goroutine scheduler then from there it runs ->

						1. The Go runtime schedules the function to run in the background, in a new goroutine.

						2. The function runs concurrently, managed by the Go scheduler.

						3. The actual execution may happen on the same OS thread, a different thread, or even a different CPU core — depending on runtime needs.

*/

func solve1(t int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Goroutines", t)
}

func main() {
	fmt.Println("Go Starts")

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		//solve1(i) // ---> what this will do, this will print each line sequentially ---> like first, then second and so on
		wg.Add(1)
		go solve1(i, &wg)

		//go func(i int) {
		//	fmt.Println("Goroutines", i)
		//}(i)

	}

	wg.Wait()

	//time.Sleep(time.Second)
	fmt.Println("Go Ends")
}
