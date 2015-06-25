package main

import "net/smtp"

type SMTPMailer struct{}

func (m *SMTPMailer) Send(to []string, from string, body []byte) error {
	return smtp.SendMail("localhost:25", nil, from, to, body)
}
