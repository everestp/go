package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp" // RabbitMQ client
	"gopkg.in/gomail.v2"        // Email sending library
)

// EmailData represents the structure of email to send
type EmailData struct {
	To      string `json:"to"`      // Recipient email address
	Subject string `json:"subject"` // Email subject
	Body    string `json:"body"`    // Email content (HTML allowed)
}

func main() {

	// ----------------------
	// 1. Connect to RabbitMQ server
	// ----------------------
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // default RabbitMQ credentials
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close() // ensure connection is closed when main exits

	// ----------------------
	// 2. Open a channel
	// ----------------------
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}
	defer ch.Close() // close channel when done

	// ----------------------
	// 3. Declare a queue
	// ----------------------
	// Queue ensures that messages are stored until consumed
	queue, err := ch.QueueDeclare(
		"email_queue", // queue name
		true,          // durable -> survives server restart
		false,         // autoDelete -> do not delete when unused
		false,         // exclusive -> multiple connections can use
		true,          // noWait -> wait for server response
		nil,           // arguments -> optional
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	// ----------------------
	// 4. Register a consumer to listen to messages
	// ----------------------
	msgs, err := ch.Consume(
		queue.Name, // queue name
		"",         // consumer name, empty = auto generated
		true,       // autoAck -> messages are acknowledged automatically
		false,      // exclusive -> false means multiple consumers allowed
		false,      // noLocal -> not supported by RabbitMQ
		false,      // noWait -> wait for server response
		nil,        // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	// ----------------------
	// 5. Listen for messages forever
	// ----------------------
	forever := make(chan bool) // block main so program doesn't exit
	fmt.Println("Waiting for email messages...")

	// Launch goroutine to process incoming messages
	go func() {
		for msg := range msgs {
			fmt.Printf("\nReceived Message: %s\n", msg.Body)

			// Parse JSON message into EmailData struct
			var email EmailData
			if err := json.Unmarshal(msg.Body, &email); err != nil {
				log.Printf("Invalid email data: %v", err)
				continue // skip this message if parsing fails
			}

			// Call function to send email
			sendEmail(email)
		}
	}()

	<-forever // keep main running
}

// ----------------------
// 6. Function to send email using gomail
// ----------------------
func sendEmail(email EmailData) {
	m := gomail.NewMessage()                   // create new email message
	m.SetHeader("From", "3verestp@gmail.com") // sender
	m.SetHeader("To", email.To)               // recipient
	m.SetHeader("Subject", email.Subject)     // subject
	m.SetBody("text/html", email.Body)        // email body, HTML format supported

	// Create SMTP dialer (Gmail in this case)
	d := gomail.NewDialer("smtp.gmail.com", 587, "3verestp@gmail.com", "your_app_password_here")

	// Send email
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send email to %s: %v", email.To, err)
		return
	}

	log.Printf("Email sent successfully to: %s", email.To)
}

/* ----------------------
Explanation of how this works:

1. RabbitMQ Connection:
   - Connects to RabbitMQ server using default credentials.
   - A channel is required for declaring queues and consuming messages.

2. Queue Declaration:
   - Ensures queue exists before sending or receiving messages.
   - Durable queue survives server restarts.

3. Consumer Registration:
   - `Consume` returns a channel (`msgs`) which receives messages from the queue.

4. Goroutine for Processing:
   - Processes messages concurrently.
   - Parse
