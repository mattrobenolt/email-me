package main

import (
	"bytes"
	"os/exec"
)

type SendMailMailer struct{}

func (m *SendMailMailer) Send(to []string, from string, body []byte) error {
	sendmail := exec.Command("sendmail", to...)
	sendmail.Stdin = bytes.NewReader(body)
	return sendmail.Run()
}
