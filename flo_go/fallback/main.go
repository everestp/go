package main

import (
	"cmp"
	"fmt"
	"os"
)


func main(){
port := cmp.Or(getPortFlag(),getPortFromEnv(),"")
fmt.Println("Starting server at port",port)
}

func getPortFlag() string{
	return ""
}

func getPortFromEnv() string {
	return  os.Getenv("PORT")
}
