package mail

import (
	"bytes"
	"github.com/TheFlies/ofriends/internal/app/types"
	envconfig "github.com/TheFlies/ofriends/internal/pkg/config/env"
	"github.com/go-gomail/gomail"
	"github.com/pkg/errors"
	"html/template"
)

type EmailConfig struct {
	OutGoingMailServer string `envconfig:"OUT_GOING_MAIL_SERVER" default:"smtp.tma.com.vn"`
	MailServerPort     int    `envconfig:"OUT_GOING_MAIL_PORT" default:"25"`
	Subject            string `envconfig:"EMAIL_SUBJECT" default:"Prepare for customer visit"`
	MailFrom           string `envconfig:"MAIL_FORM" default:"ofriends.oficial@tma.com.vn"`
	MailTo             string `envconfig:"MAIL_TO" default:"pdkhang@tma.com.vn"` // for test, to avoiding  send real to internal communication department
}

func SendEmail(data []types.NotificationData, htmlTemplate *template.Template) error {
	var conf EmailConfig
	envconfig.Load(&conf)

	dataForMail := types.VisitsData{
		Data: data,
	}
	htmlData := new(bytes.Buffer)
	err := htmlTemplate.Execute(htmlData, dataForMail)
	if err != nil {
		return errors.Wrap(err, "can't generate html data")
	}
	message := gomail.NewMessage()
	message.SetHeader("From", conf.MailFrom)
	message.SetHeader("To", conf.MailTo)
	message.SetHeader("Subject", conf.Subject)
	message.SetBody("text/html", htmlData.String())

	mailDialer := gomail.Dialer{
		Host: conf.OutGoingMailServer,
		Port: conf.MailServerPort,
	}
	if err := mailDialer.DialAndSend(message); err != nil {
		return errors.Wrap(err, "sending email is fail")
	}
	return nil

}
