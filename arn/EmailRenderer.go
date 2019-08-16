package arn

import (
	"context"
	"fmt"
	"time"

	"github.com/mailgun/mailgun-go/v3"
)

// HTMLEmailRenderer is the instance used for rendering emails.
var HTMLEmailRenderer EmailRenderer

// EmailRenderer is an interface for rendering HTML emails.
type EmailRenderer interface {
	Notification(notification *Notification) string
}

// SendEmail sends an e-mail.
func SendEmail(email string, subject string, html string) error {
	mg := mailgun.NewMailgun(APIKeys.Mailgun.Domain, APIKeys.Mailgun.PrivateKey)
	sender := fmt.Sprintf("Anime Notifier <notifications@%s>", APIKeys.Mailgun.Domain)
	message := mg.NewMessage(sender, subject, "", email)
	message.SetHtml(html)

	// Allow a 10-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message
	_, _, err := mg.Send(ctx, message)
	return err
}
