package main

import "fmt"

type Notification interface {
	Send(message string)
}

type EmailNotification struct{}

func (EmailNotification) Send(message string) {
	// Email sending logic goes here
	fmt.Printf("From email end : %s", message)
}

type SMSNotification struct{}

func (SMSNotification) Send(message string) {
	fmt.Printf("From SMS end : %s", message)
}

func CreateNotification(notificationType string) (Notification, error) {
	switch notificationType {
	case "email":
		return EmailNotification{}, nil
	case "sms":
		return SMSNotification{}, nil
	default:
		return nil, fmt.Errorf("unknown type : %s", notificationType)
	}
}

// Usage:
// notif, _ := CreateNotification("email")
// notif.Send("Hello")

// func main() {
// 	notif, _ := CreateNotification("email")
// 	notif.Send("Hello")
// }
