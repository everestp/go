package main

import (
	"fmt"
	"log"
	"net/http"
)


func main(){
	http.HandleFunc("/",  func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w,"Hello server")
	})
	const port string = ":3000"
	fmt.Println("Server is listening om port :",port)
	err :=http.ListenAndServe(port, nil)
	if err != nil{
		log.Fatalln("error starting the server",err)
	}
	
}