package main

import "fmt"

/*
	In system design -> our function should be ***closed for modification but open for extensions***

	below is the prob
*/

type payments struct {
}

func (p payments) makePayments(amount float64) {
	razorpayPaymentGw := razorpay{}
	razorpayPaymentGw.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float64) {
	//logic
	fmt.Println("Making payment through razorpay", amount)
}

func main() {
	fmt.Println("Go Start...")
	newPayments := payments{}
	newPayments.makePayments(100)
	fmt.Println("Go End...")
}
