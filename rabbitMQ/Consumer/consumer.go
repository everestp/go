package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"subject"`
}

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to Connect rabbitmq %v", err)
	}
	defer conn.Close()

	// create connection
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open the channel or connection")
	}
	defer ch.Close()
	queue, err := ch.QueueDeclare("email_queue",
		true,  // Durable-> survive when server re-start
		false, // Auto Delete -> when server re-start
		false, //Exclusive ->accepts to multiple connections
		true,  //NoWaits ->wait for server response
		nil,   //Arguments
	)
	if err != nil {
		log.Fatalf("Failed to  declare queue")
	}

	msgs, err := ch.Consume(
		queue.Name,
		"", // Leave blank  for auto generated
		true,
		false, //exclusive
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to  Register the consumer")
	}

	forever := make(chan bool)
	fmt.Println("Waiting for  email messages")
	go func() {
		for msg := range msgs {

			fmt.Printf("\n  Recieved Message %s \n", msg.Body)
			var email EmailData
			if err := json.Unmarshal(msg.Body, &email); err != nil{
				log.Printf("Invalid email data %v", err)
				continue
			}
			sendEmail(email)
		}

	}()
<-forever
}

func sendEmail(email EmailData){
	m := gomail.NewMessage()
	m.SetHeader("From", "3verestp@gmail.com")
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", email.Body)
	d := gomail.NewDialer("smtp.gmail.com", 587,"3verestp@gmail.com", "uagi geva qnhs yrgh")
	if err := d.DialAndSend(m) ; err != nil{
		log.Printf("FAiled to send email")
		return
	}
	log.Printf("Email sent sucessfully to = %s",email.To)
}