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
	//
	//data, err := os.ReadFile("example.txt")
	//if err != nil {
	//	panic(err)
	//}

	////read folders
	//dir, err := os.Open(".")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer dir.Close()
	//
	//fileInfo, err := dir.ReadDir(-1)
	//
	//for _, fi := range fileInfo {
	//	fmt.Println(fi.Name())
	//}

	//create a file
	//f, err := os.Create("example2.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//f.WriteString("Hello World")
	//f.WriteString("Hello World2")
	//
	//bytes := []byte("hello world")
	//f.Write(bytes)

	//read and write to another file(streaming fashion)

	//sourceFile, err := os.Open("example.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer sourceFile.Close()
	//
	//destFile, err := os.Create("example2.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer destFile.Close()
	//
	//reader := bufio.NewReader(sourceFile)
	//writer := bufio.NewWriter(destFile)
	//
	//for {
	//	b, err := reader.ReadByte()
	//	if err != nil {
	//		if err.Error() == "EOF" {
	//			panic(err)
	//		}
	//		break
	//	}
	//	e := writer.WriteByte(b)
	//	if e != nil {
	//		panic(e)
	//	}
	//}
	//
	//writer.Flush()
	//
	//fmt.Println("written to new file successfully")
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Removed file example2.txt")
}
