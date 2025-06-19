package main

import "fmt"

/*
	Problem Without interface -> Imagine you're building a program where different animals make different sounds. You want a function to make any animal speak. You need to write a new function every time you add a new animal. You can't easily generalize behavior like "make any animal speak".Your code is repetitive and tightly coupled to concrete types.
	Inshort: --> ***for efficient function, it should follow -> ```Open for extensions but, close for modification``` ***
										****code ->         package main

														import "fmt"

														type Dog struct{}
														type Cat struct{}

														func DogSpeak(d Dog) {
															fmt.Println("Woof!")
														}

														func CatSpeak(c Cat) {
															fmt.Println("Meow!")
														}

														func main() {
															d := Dog{}
															c := Cat{}

															DogSpeak(d)
															CatSpeak(c)
														}****


	Solution With interface -> We define an Animal interface with a Speak() method. Dog and Cat both implement that interface by defining Speak(). The MakeItSpeak() function now works with any type that satisfies the Animal interface.
									****code ->         package main

														import "fmt"

														// Define an interface
														type Animal interface {
															Speak() string
														}

														// Dog implements the interface
														type Dog struct{}
														func (d Dog) Speak() string {
															return "Woof!"
														}

														// Cat implements the interface
														type Cat struct{}
														func (c Cat) Speak() string {
															return "Meow!"
														}

														// General function using the interface
														func MakeItSpeak(a Animal) {
															fmt.Println(a.Speak())
														}

														func main() {
															d := Dog{}
															c := Cat{}

															MakeItSpeak(d)
															MakeItSpeak(c)
														}
														****
Interface Gives :-

Flexibility: Easily add new animal types without changing the function.

Reusability: One function handles all "speaking" behavior.

Loose coupling: Your code depends on behavior (interface), not concrete types.




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
