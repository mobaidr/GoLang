package main

import (
	"bytes"
	"fmt"
	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"
	"html/template"
	"sync"
	"time"
)

type Mail struct {
	Domain         string
	Host           string
	Port           int
	Username       string
	Password       string
	Encryption     string
	FromAddress    string
	FromName       string
	Wait           *sync.WaitGroup
	MailerChan     chan Message
	ErrorChan      chan error
	MailerDoneChan chan bool
}

type Message struct {
	From          string
	FromName      string
	To            string
	Subject       string
	Attachments   []string
	AttachmentMap map[string]string
	Data          interface{}
	DataMap       map[string]interface{}
	Template      string
}

//Function to listen for messages on the MailerChan

func (app *Config) listenForMail() {
	for {
		select {
		case msg := <-app.Mailer.MailerChan:
			{
				go app.Mailer.sendMail(msg, app.Mailer.ErrorChan)
			}
		case err := <-app.Mailer.ErrorChan:
			{
				app.ErrorLog.Println(err)
			}
		case <-app.Mailer.MailerDoneChan:
			{
				return
			}
		}
	}
}

func (m *Mail) sendMail(msg Message, errorChan chan error) {

	defer m.Wait.Done()

	if msg.Template == "" {
		msg.Template = "mail"
	}

	if msg.From == "" {
		msg.From = m.FromAddress
	}

	if msg.FromName == "" {
		msg.FromName = m.FromName
	}

	if msg.AttachmentMap == nil {
		msg.AttachmentMap = make(map[string]string)
	}

	if len(msg.DataMap) == 0 {
		msg.DataMap = make(map[string]interface{})
	}

	msg.DataMap["message"] = msg.Data

	//Build HTML mail
	formattedMessage, err := m.BuildHTMLMessage(msg)
	if err != nil {
		errorChan <- err
	}

	//BUIl plain mail
	plainMessage, err := m.BuildPlainTextMessage(msg)
	if err != nil {
		errorChan <- err
	}

	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.getEncryption(m.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()

	if err != nil {
		errorChan <- err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)
	email.SetBody(mail.TextPlain, plainMessage)
	email.AddAlternative(mail.TextHTML, formattedMessage)

	if len(msg.Attachments) > 0 {
		for _, x := range msg.Attachments {
			email.AddAttachment(x)
		}
	}

	if len(msg.AttachmentMap) > 0 {
		for k, v := range msg.AttachmentMap {
			email.AddAttachment(v, k)
		}
	}

	err = email.Send(smtpClient)
	if err != nil {
		errorChan <- err
	}
}

func (m *Mail) BuildHTMLMessage(msg Message) (string, error) {
	templateToRender := fmt.Sprintf("./cmd/web/templates/%s.html.gohtml", msg.Template)

	t, err := template.New("emai-html").ParseFiles(templateToRender)

	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	plainMessage := tpl.String()

	return plainMessage, nil
}

func (m *Mail) BuildPlainTextMessage(msg Message) (string, error) {
	templateToRender := fmt.Sprintf("./cmd/web/templates/%s.plain.gohtml", msg.Template)

	t, err := template.New("email-plain").ParseFiles(templateToRender)

	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	formattedMessage := tpl.String()

	if err != nil {
		return "", err
	}

	return formattedMessage, nil
}

func (m *Mail) inlineCSS(s string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(s, &options)

	if err != nil {
		return "", err
	}

	html, err := prem.Transform()

	if err != nil {
		return "", err
	}

	return html, nil
}

func (m *Mail) getEncryption(e string) mail.Encryption {
	switch e {
	case "tls":
		{
			return mail.EncryptionSTARTTLS
		}
	case "ssl":
		{
			return mail.EncryptionSSLTLS
		}
	case "none":
		{
			return mail.EncryptionNone
		}
	default:
		{
			return mail.EncryptionSTARTTLS
		}
	}
}
