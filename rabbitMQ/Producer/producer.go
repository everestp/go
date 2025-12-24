package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main(){
	
	conn , err :=amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to Connect rabbitmq %v",err)
	}
	defer conn.Close()

	// create connection
	ch , err :=conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open the channel or connection")
	}
	defer ch.Close()
	 queue , err :=ch.QueueDeclare("email_queue",
	  true,  // Durable-> survive when server re-start
	   false,   // Auto Delete -> when server re-start 
	   false,  //Exclusive ->accepts to multiple connections
	    true,   //NoWaits ->wait for server response
		 nil,  //Arguments  
		)
		if err != nil {
		log.Fatalf("Failed to  declare queue")
	}

	emailBody := `
	{
	"to":"3verestp@gmail.com",
	"subject":"Welcome to golang and Devops Channel",
	"body":"Subscribe please and give motivation"
	}
	


	`
 ch.Publish(
	"", //Exchange name if not use default exchange
 queue.Name , //Queue name
  false, //Mandoratory -. do you want to drop message if queue not found
   false,  // if you want to find consumer immediately
  amqp.Publishing{
	ContentType: "application/json",  // tell consumer data type
	Body: []byte(emailBody),  // actual data in byte
  })
if err != nil {
	log.Fatalf("Failed to published messages %v", err)
}
fmt.Println("Email request is sent to queue",emailBody)


}