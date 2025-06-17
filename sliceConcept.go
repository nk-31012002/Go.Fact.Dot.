package main

import "fmt"

/*
	slices -> dynamic array
	most used construct in Go
	most useful methods

	almost same as vector in cpp
*/

func main() {
	fmt.Println("Go Start")

	//uninitialized slice is nil
	//var nums []int
	//fmt.Println(len(nums))

	//var nums = make([]int, 2)
	//fmt.Println("inititally =", nums)
	//fmt.Println(cap(nums))
	//
	//for i := 0; i < 6; i++ {
	//	nums = append(nums, i)
	//}
	//fmt.Println(cap(nums))
	//fmt.Println(nums)

	////copy-slice
	//var nums = make([]int, 5)
	//for i := 0; i < 5; i++ {
	//	nums[i] = i
	//}
	//fmt.Println(nums)
	//var nums2 = make([]int, len(nums))
	//copy(nums2, nums)
	//fmt.Println(nums2)
}
