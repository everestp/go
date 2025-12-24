package main

import (

	"fmt"
)



type LogLevel int

const (
	LevelTrace = 0 
	LevelDebug=1
	LevelWarining=3
	LevelError=4
)


var levelName =[]string{"Trace","Debug","Info","Waring","Error"}


func (l LogLevel) String() string {
	if l <LevelTrace || l > LevelError {
		return "Unknown"
	}
	return  levelName[l]
}


func printLogLevel(level LogLevel){
	fmt.Printf("Log Level: %d %s\n",level ,level.String())

}

func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelWarining)
	printLogLevel(LevelError)
}
