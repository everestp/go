package main

import (
	"fmt"
	"slices"
)
 func main(){
	var nums []int
	var nums1 = make([]int, 3 ,8)
	nums1 = append(nums1, 1)
	nums1 = append(nums1, 1)
	nums1 = append(nums1, 1)
	nums1 = append(nums1, 1)
	nums1 = append(nums1, 1)
	fmt.Println(nums)
	fmt.Println(nums1)
	fmt.Println(cap(nums1))
	var num2 = make([]int, len(nums))
	num2 = append(num2, 2)
	fmt.Println(nums1,num2)

var nums3 = []int{1,2,5}
var nums4 = []int{1,2,5}
fmt.Println(slices.Equal(nums3, nums4))
 }