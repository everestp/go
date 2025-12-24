package main

import "fmt"


func main(){
	// for -- only way to loop

	//C-style loop
	for i :=1; i<=10; i++{
		fmt.Println(i)
	}

	// while -stlye
	k :=3
	 for k>0{
		fmt.Println(k)
		k--

	 }


	  fmt.Println("----------------------infinit loop---------")
	  counter := 0
	  for {
		fmt.Println("cunter",counter)
		counter ++
		if counter >=50{
			break
		}
	  }
	   fmt.Println("----------------------array---------")
	   items := [3]string{"eversrt"}
}
