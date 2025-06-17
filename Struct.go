package main

import "fmt"

/*

 struct --> same as in C only the syntax changes

*/

//type Person struct {
//	name string
//	age  int
//}
//
//func (p Person) Greet() {
//	fmt.Printf("Hello, my name is %s.\n", p.name)
//}
//
//// Constructor function
//func NewPerson(name string, age int) Person {
//	return Person{name: name, age: age}
//}
//
//// Pointer receiver (modifies the struct)
//func (p *Person) Birthday() {
//	p.age++
//}

func main() {
	fmt.Println("Go Start...")
	//
	////creating an instances
	//var p Person
	//p.name = "Jack"
	//p.age = 23
	//
	//fmt.Println(p)
	//
	////or
	//
	////var p1 Person
	//p1 := Person{"Jack", 23}
	//fmt.Println(p1)
	//p3 := &Person{name: "Charlie", age: 40}
	//fmt.Println(p3.name) // "Charlie"
	//p3.age = 41
	//
	//p := NewPerson("Alice", 30)
	//fmt.Println(p)
	//
	//language :=
	//	struct {
	//		name string
	//		age  int
	//	}{"Amnic", 4}
	//
	//fmt.Println(language)
	//fmt.Println(language.name)
	//fmt.Println(language.age)
	fmt.Println("Go End...")
}
