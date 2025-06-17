package main

import (
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	fmt.Println("Go Start")

	//shorthand syntax
	//name := "string"
	//fmt.Println(name)
	//name1 := 1
	//fmt.Println(name1)
	//var name2 float32 = 50.61111
	//fmt.Println(name2)
	//
	//const name3 string = "Bruce wayne is batman"
	//fmt.Println(name3)

	//const (
	//	port = 8080
	//	host = "localhost"
	//)

	//fmt.Println("https://" + host + ":" + strconv.Itoa(port) + ".com")

	/*
		//while loop

		i := 1
		for i <= 3 {
			fmt.Println(i)
			i++
		}

	*/

	//i=1

	/*
		//infinite loop

		for {
			fmt.Println(i)
			i++
		}
	*/

	/*
		//classic for-loop

		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	*/

	/*

		//rang-key word

		for i := range 5 { //---> print from 0 to 4
			fmt.Println(i)
		}
	*/

	/*
			//if-else-condition

			i := 10

			if i == 10 {
				fmt.Println("Go End")
			} else {
				fmt.Println("Go continue")
			}

		//else-if->elif-> same as for cpp
		//logical operator-> same as for cpp

	*/

	/*
		//declaring variable in if condition itself

		if age := 10; age == 10 {
			fmt.Println("Go End")
		}
	*/

	//inside go there is no ternary operator
	//switch statememt->almost same as cpp

	//type switch
	/*
		ans := func(i interface{}) {
			switch t := i.(type) {
			case int:
				fmt.Println("integer", t)
			case bool:
				fmt.Println("boolean", t)
			case string:
				fmt.Println("string", t)
			}
		}
		ans("string")
	
	*/
}
