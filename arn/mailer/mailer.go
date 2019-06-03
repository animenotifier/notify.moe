package mailer

import (
	"github.com/animenotifier/notify.moe/arn"
	gomail "gopkg.in/gomail.v2"
)

// SendEmailNotification sends an e-mail notification.
func SendEmailNotification(email string, notification *arn.PushNotification) error {
	m := gomail.NewMessage()
	m.SetHeader("From", arn.APIKeys.SMTP.Address)
	m.SetHeader("To", email)
	m.SetHeader("Subject", notification.Title)
	m.SetBody("text/html", "<h2>"+notification.Message+"</h2><p><a href='"+notification.Link+"' target='_blank'><img src='"+notification.Icon+"' alt='Anime cover image' style='width:125px;height:181px;'></a></p>")

	d := gomail.NewDialer(arn.APIKeys.SMTP.Server, 587, arn.APIKeys.SMTP.Address, arn.APIKeys.SMTP.Password)

	// Send the email
	return d.DialAndSend(m)
}
