package pattern

import (
	"fmt"
)

type EmailSender struct{}

func (e *EmailSender) SendEmail(to, message string) {
    fmt.Printf("Sending email to %s: %s\n", to, message)
}

type SMSSender struct{}

func (s *SMSSender) SendSMS(to, message string) {
    fmt.Printf("Sending SMS to %s: %s\n", to, message)
}

type PushSender struct{}

func (p *PushSender) SendPushNotification(to, message string) {
    fmt.Printf("Sending push to %s: %s\n", to, message)
}

type NotificationService struct {
    emailSender        *EmailSender
    smsSender          *SMSSender
    pushNotificationSender *PushSender
}

func NewNotificationService() *NotificationService {
    return &NotificationService{
        emailSender:        &EmailSender{},
        smsSender:          &SMSSender{},
        pushNotificationSender: &PushSender{},
    }
}

func (ns *NotificationService) SendNotification(to, message string) {
    ns.emailSender.SendEmail(to, message)
    ns.smsSender.SendSMS(to, message)
    ns.pushNotificationSender.SendPushNotification(to, message)
}



//Паттерн "Фасад" предоставляет простой интерфейс, но  если подсистема становится слишком сложной, фасад также может стать сложным и трудно поддерживаемым.
//Клиенты взаимодействуют только с фасадом, что уменьшает связность между клиентами и подсистемой. Но при этом клиентам может потребоваться более низкоуровневый