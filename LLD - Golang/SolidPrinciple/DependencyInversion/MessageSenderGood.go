package dependencyinversion

type MessageSender interface {
	Send(message string)
}

type EmailMessageSender struct{}

func (EmailMessageSender) Send(message string) {
	_ = message
}

type NotificationServiceGood struct {
	sender MessageSender
}

func NewNotificationServiceGood(sender MessageSender) *NotificationServiceGood {
	return &NotificationServiceGood{sender: sender}
}

func (s *NotificationServiceGood) Notify(message string) {
	s.sender.Send(message)
}
