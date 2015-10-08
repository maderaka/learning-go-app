package email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
	"text/template"
)

type EmailUser struct {
	Username string
	Password string
	Host     string
	Port     int
}

type Engine struct {
	User *EmailUser
}

func NewEmailEngine(username, password, host string, port int) *Engine {
	return &Engine{&EmailUser{
		username,
		password,
		host,
		port,
	}}
}

func (this *Engine) SendEmail(receivers []string, subject, content string) (err error) {

	params := struct {
		From    string
		To      string
		Subject string
		Message string
	}{
		this.User.Username,
		strings.Join([]string(receivers), ","),
		subject,
		content,
	}

	buffer := new(bytes.Buffer)

	t := template.Must(template.New("emailTemplate").Parse(emailScript()))
	t.Execute(buffer, &params)

	auth := smtp.PlainAuth("", this.User.Username, this.User.Password, this.User.Host)
	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", this.User.Host, this.User.Port),
		auth,
		this.User.Username,
		receivers,
		buffer.Bytes(),
	)
	return err
}

func emailScript() (script string) {
	return `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}
MIME-version: 1.0
Content-Type: text/html; charset="UTF-8"

{{.Message}}`
}
