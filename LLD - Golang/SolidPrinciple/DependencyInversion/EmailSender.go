package dependencyinversion

type EmailSender struct{}

func (EmailSender) Send(message string) {
	_ = message
}

// High level module depends directly on low-level details
type NotificationServiceBad struct {
	emailSender EmailSender
}

func NewNotificationServiceBad() *NotificationServiceBad {
	return &NotificationServiceBad{emailSender: EmailSender{}}
}

func (s *NotificationServiceBad) Notify(message string) {
	s.emailSender.Send(message)
}
