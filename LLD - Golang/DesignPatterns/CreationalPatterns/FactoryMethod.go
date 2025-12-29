package main

import "fmt"

type Notification interface {
	Send(message string)
}

type EmailNotification struct{}

func (EmailNotification) Send(message string) {
	// Email sending logic goes here
}

type SMSNotification struct{}

func (SMSNotification) Send(message string) {
	// SMS sending logic here
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
