package main

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

//func processNum(numChan chan int) {
//	for num := range numChan {
//		fmt.Println("Processing number", num)
//		time.Sleep(time.Second)
//	}
//}

//func sum(result chan int, num1 int, num2 int) {
//	ans := num1 + num2
//	result <- ans
//}

//func task(done chan bool) {
//
//	defer func() { done <- true }()
//
//	fmt.Println("Processing.....")
//}

func main() {

	//done := make(chan bool)
	//go task(done)
	//
	//<-done

	//numChan := make(chan int)
	//
	//go sum(numChan, 1, 23)
	//
	//res := <-numChan //blocking
	//
	//fmt.Println(res)
	//
	//go processNum(numChan)
	//
	//for {
	//	numChan <- rand.Intn(100)
	//}
	//messageChan := make(chan string)
	//
	//messageChan <- "Ping" //channels are blocking
	//msg := <-messageChan
	//
	//fmt.Println(msg)

}
