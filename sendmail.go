package main

import (
	"bytes"
	"os/exec"
)

func FindSendmail() (path string, err error) {
	for _, option := range []string{"/usr/sbin/sendmail", "sendmail"} {
		path, err = exec.LookPath(option)
		if err == nil {
			break
		}
	}
	return
}

type SendMailMailer struct {
	path string
}

func (m *SendMailMailer) Send(to []string, from string, body []byte) error {
	sendmail := exec.Command(m.path, to...)
	sendmail.Stdin = bytes.NewReader(body)
	return sendmail.Run()
}
