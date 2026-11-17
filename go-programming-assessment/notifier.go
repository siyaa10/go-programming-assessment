package main

import "fmt"

// 1. Define the Notifier interface
type Notifier interface {
	SendNotification(message string) error
}

// 2. Create EmailNotifier struct
type EmailNotifier struct{}

// 3. Create SMSNotifier struct
type SMSNotifier struct{}

// 4. Implement the interface for EmailNotifier
func (e EmailNotifier) SendNotification(message string) error {
	fmt.Println("Email Notification:", message)
	return nil
}

// 5. Implement the interface for SMSNotifier
func (s SMSNotifier) SendNotification(message string) error {
	fmt.Println("SMS Notification:", message)
	return nil
}
