package main

import "fmt"


func main(){
    fmt.Println(add(5, 10))
	lang1 ,lang2,lang3 , number := getLangugages()
	fmt.Println("we can return multiple value=>",lang1,lang2,lang3,number)
}

func add(a int, b int ) int {
	return  a+b
}

func getLangugages()(string,string,string,int){
	return "golang","javascript","rust", 5
}
