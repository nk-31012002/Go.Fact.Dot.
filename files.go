package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("file name:", fileInfo.Name())
}
