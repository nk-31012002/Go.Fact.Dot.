package main

import "fmt"

/*
	Enumerated types
*/

type OrderStatus int

const (
	Recieved OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	fmt.Println("Go Starts...")

	fmt.Println(OrderStatus(Recieved))

	fmt.Println("Go Ends....")
}
