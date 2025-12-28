package main

import (
	"fmt"

	"time"
)
type customer struct{
	name string
}
// if you don't set any field , default value is zero value
type order  struct{
	id string
	amount float32
	status string
	createdAt time.Time   // nanosecond  precision
	customer customer
}

//constructor like 
//intial setup goes here...
func newOrder(id string , amount float32 ,status string) *order{
	myOrder := order{
		id: id,
		amount: amount,
		status: status,
		customer: customer{
			name: "Everest",
		},
	}
	return &myOrder
}

//reveiver method
func(o *order) changeStatus(status string){
        o.status=status
		 fmt.Println(o.status)
}

func main(){

	langugage :=struct{
  name string
	}{"evert"}
myOrder :=newOrder("1", 30, "newOrder")
 
 myOrder.changeStatus("changr")

 fmt.Println(myOrder.status)
  fmt.Println("New Order",*myOrder)
 fmt.Println(langugage)
}
