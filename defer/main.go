package main

import (
	"fmt"
	"io"
	"os"
	"time"
)



func readFile(filename string) error{
file , err := os.Open(filename)
if err != nil {
	return err
}
defer file.Close()
data , err := io.ReadAll(file)
if err != nil {
	return err
}
fmt.Println(string(data))
return nil


}

func safeOperation(){
	defer func(){
		if r := recover(); r != nil{
			fmt.Println("Recover from panic")
		}
	}()
	panic("Something went wrong")
	fmt.Println("Cannot reach this code")
}



func processData(data []int){
	start :=time.Now()

	defer func(){
		fmt.Println("Data process is completed in :=",time.Since(start))
	}()
	for _ ,d :=range data {
		fmt.Println(d)
		time.Sleep(time.Millisecond *10)
	}
}

func main(){
  err := readFile("output.txt")
  if err != nil{
	fmt.Println(err)
  }
data := []int{1,2,3,4,5,6}
processData(data)
safeOperation()

}