package main

import (
	"fmt"
	"time"
)

type Logger struct{}

func (l Logger) Log(message string) {
	fmt.Println(time.Now().Format("15:04:05"), "LOG:", message)
}

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct {
	Logger
	EmailNotifier string
}

func (e EmailNotifier) Send(message string) {
	e.Log("Sending email to " + e.EmailNotifier)
	fmt.Printf("Email sent to %s: %s\n", e.EmailNotifier, message)
}

type SMSNotifier struct {
	Logger
	PhoneNumber string
}

func (s SMSNotifier) Send(message string) {
	s.Log("Sending SMS to  " + s.PhoneNumber)
	fmt.Printf("SMS sent to %s: %s\n ", s.PhoneNumber, message)
}

func SendNotification(n Notifier, message string) {
	n.Send(message)
}

func main() {
	email1 := EmailNotifier{EmailNotifier: "userExample@uss.edu.pe"}
	sms := SMSNotifier{PhoneNumber: "+51 999999999"}

	SendNotification(email1, "Tu pedido esta en camino ")
	SendNotification(sms, "Tu codigo es 01020")

	fmt.Scanln()
}
