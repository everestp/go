package main

import "fmt"

//map -> hash ,object , maps

func main(){
	//creating map

	m := make(map[string]string)
	// setting an element
	m["name"]="golang"
  //IMP-> if key  does  not exist in the map then it return zero value
  k:=make(map[string]int)
  k["everest"]=982827
  k["sirjana"]=982182
  fmt.Println("This the data",k["everest"])
	fmt.Println(m["name"])
	// delete(k,"sirjana")
	fmt.Println("This the data",k)
	m1:=map[string]int{
		"Price":9,
		"Everest":43,
		"Sirjana":143,
	}
	fmt.Println("This the data",m1)

// get Item from the map
el,ok:=m1["Sirjana"]
if ok{
	fmt.Println("All ok  The value is :=",el)
} else{
	fmt.Println("Not ok")
}


	clear(k)
}