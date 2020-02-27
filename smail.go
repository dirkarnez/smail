package smail

import (
	"errors"

	"gopkg.in/gomail.v2"
)

type smail struct {
	email    string
	password string
	smtp     string
	port     int
}

var s *smail = nil

// Dial dials and authenticates to an SMTP server. The returned gomail.SendCloser should be closed when done using it.
func Dial(email, password, smtp string, port int) (gomail.SendCloser, error) {
	s = &smail{email, password, smtp, port}
	return gomail.NewDialer(s.smtp, s.port, s.email, s.password).Dial()
}

// Send sends emails using the given Sender.
func Send(sender gomail.Sender, subject, body string) error {
	if sender == nil {
		return errors.New("sender is nil")
	}

	message := gomail.NewMessage()
	message.SetHeader("From", s.email)
	message.SetHeader("To", s.email)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)
	return gomail.Send(sender, message)
}
