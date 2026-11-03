package main

import (
	"fmt"
	"time"
)

func timeres() {

	timer := time.NewTimer(2 * time.Second)
	<-timer.C //blocking nature
	timer.Stop()
	fmt.Println("Timer expired")
}
