package smail

import (
	"errors"
	"gopkg.in/gomail.v2"
)

type smail struct {
	email string
	password string
	smtp string
	port int
}

type Smail interface {
	Dial() (gomail.SendCloser, error)
	Send(gomail.Sender, string) error
}

func NewSmail(email, password, smtp string, port int) Smail {
    return &smail{email, password, smtp, port}
}

func (smail *smail) Dial() (gomail.SendCloser, error) {
	return gomail.NewDialer(smail.smtp, smail.port, smail.email, smail.password).Dial()
}

func (smail *smail) Send(sender gomail.Sender, subject string) error {
	if sender == nil {
		return errors.New("sender is nil")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", smail.email)
	m.SetHeader("To", smail.email)
	m.SetHeader("Subject", subject)
	return gomail.Send(sender, m)
}