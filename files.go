package main

import (
	"fmt"
	"os"
)

func main() {
	//
	//f, err := os.Open("example.txt")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fileInfo, err := f.Stat()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("file name:", fileInfo.Name())

	//read file

	//f, err := os.Open("example.txt")

	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//
	//buf := make([]byte, 10)
	//d, err := f.Read(buf)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for i := 0; i < len(buf); i++ {
	//	println("data", d, string(buf[i]))
	//}

	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
