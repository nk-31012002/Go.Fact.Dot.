package main

import "fmt"

func updated(num *int) {
	*num = 1

	fmt.Println(num)

}

func main() {

	fmt.Println("Go Start...")

	num := 2

	updated(&num)
	fmt.Println(num)

	fmt.Println("Go End...")
}
