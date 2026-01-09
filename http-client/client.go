package main

import (
	"fmt"
	"io"
	"net/http"
)



func main(){
	//Crearte a new Http Client
	client :=&http.Client{}

	resp , err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil{
		fmt.Println("Error making GET Request",err)
		return
	}
	defer resp.Body.Close()

	//Read and print the response body'
	body , err := io.ReadAll(resp.Body)
	if err!= nil{
		fmt.Println("Error  reading the reposne body",err)
		return
	}
	fmt.Print(body)
}