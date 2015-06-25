package main

import (
	"bytes"
	"os"
	"os/exec"
	"os/user"
	"text/template"
	"time"
)

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

Cmd: {{.Result.Cmd.Args}}
Start: {{.Result.Start}}
End: {{.Result.End}}
Duration: {{.Result.Duration}} total {{.Result.Cmd.ProcessState.UserTime}} user {{.Result.Cmd.ProcessState.SystemTime}} system
ProcessState: {{.Result.Cmd.ProcessState}}
Error: {{.Result.Error}}
Stderr:{{if .Result.StderrExtra}}
... {{.Result.StderrExtra}} more bytes ...{{end}}
{{.Result.Stderr}}

Stdout:{{if .Result.StdoutExtra}}
... {{.Result.StdoutExtra}} more bytes ...{{end}}
{{.Result.Stdout}}
`

type Message struct {
	To      string
	From    string
	Subject string
	Result  *Result
}

func (m *Message) Bytes() []byte {
	var buf bytes.Buffer
	t := template.New("mail")
	t, _ = t.Parse(emailTemplate)
	t.Execute(&buf, m)
	return buf.Bytes()
}

func (m *Message) String() string {
	return string(m.Bytes())
}

type Result struct {
	Cmd         *exec.Cmd
	Error       error
	Start       time.Time
	End         time.Time
	Duration    time.Duration
	Stdout      string
	Stderr      string
	StdoutExtra int
	StderrExtra int
}

func identity() string {
	var username string
	user, err := user.Current()
	if err != nil {
		username = "root"
	} else {
		username = user.Username
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "localhost"
	}
	return username + "@" + hostname
}

type Mailer interface {
	Send([]string, string, []byte) error
}

func findMailer() Mailer {
	if _, err := exec.LookPath("sendmail"); err == nil {
		return &SendMailMailer{}
	}

	return &SMTPMailer{}
}
