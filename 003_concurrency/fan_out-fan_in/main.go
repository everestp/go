package main

import (
	"fmt"
	"sync"
	"time"
)
 type Result struct{
	Value string
	Err error
 }

func worker(url string, wg *sync.WaitGroup , resultChan chan Result) {
	defer wg.Done()
time.Sleep(time.Millisecond * 50)
fmt.Printf("imaged processed : %s\n",url)
  resultChan <- Result{
	Value: url,
	Err: nil,
  }
// return nil
}




func main(){
	var wg sync.WaitGroup
	resultChain := make(chan Result , 5)
	startTime := time.Now()
	fmt.Println("Welcome to Go Concurrency")
	wg.Add(4)
	 go worker("Image_1.png", &wg ,resultChain)
	 go worker("Image_2.png" ,&wg, resultChain)
	  go worker("Image_3.png", &wg ,resultChain)
	   go worker("Image_4.png", &wg ,resultChain)
	
	wg.Wait()
   close(resultChain)
	for result  := range resultChain{
		fmt.Printf("Received : %v\n", result)
	}
	// fmt.Println("result1:=",result1)
	// fmt.Println("result:=",result2)
	fmt.Printf("it took %s ms\n", time.Since(startTime))

}
