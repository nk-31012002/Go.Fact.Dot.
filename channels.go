package main

import "fmt"

/*

Channels are a way for goroutines to communicate with each other and synchronize execution.
You can think of them like pipes:
➡️ One goroutine sends data into the channel
⬅️ Another goroutine receives it

Declaring a Channel
ch := make(chan int) // a channel of integers
chan int means "a channel that carries integers"
Use make() to create it

Sending to a Channel
ch <- 10 // send 10 into the channel


Receiving from a Channel
value := <-ch // receive value from the channel



*/
//
//func sendData(ch chan int) {
//	ch <- 10
//}

func main() {
	fmt.Println("Go Starts")
	//ch := make(chan int)
	//go sendData(ch)
	//val := <-ch
	//fmt.Println(val)

	messageChannel := make(chan string)

	messageChannel <- "Hello World"

	msg := <-messageChannel

	fmt.Println(msg)

	fmt.Println("Go Ends")
}
