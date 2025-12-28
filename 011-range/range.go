package main

import (
	"fmt"

)

func main(){
	nums := []int{6,7,8}

	for i :=0; i<len(nums); i++{
		fmt.Println(nums[i])
	}

	for index ,num :=range nums{
		fmt.Printf("The index is %v => value is %v\n",index ,num)
	}

	m :=map[string]int{"name":3,"paudel":4}
	for k ,v := range m {
		fmt.Printf("key : %v ,value :%v\n", k ,v)
	}

	//c-> unicode code point rune
	// i-> starting byte of rune
	for i ,c := range "golang"{
		fmt.Println(i,c)
			fmt.Println(i ,string(c))
	}
}
