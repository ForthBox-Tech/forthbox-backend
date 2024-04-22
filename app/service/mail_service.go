// Mail helpers wrap the provider client used for account verification delivery.
package service

import (
	"context"
	"fmt"
	"forthboxbe/pkg/setting"
	"time"

	"github.com/mailgun/mailgun-go/v3"
)

func getMailGun() *mailgun.MailgunImpl {
	domain := setting.MailSetting.Domain
	apiKey := setting.MailSetting.MGkey
	mg := mailgun.NewMailgun(domain, apiKey)

	return mg
}

// placeholder
func SendMail(to string, subject string, text string) (string, error) {
	mg := getMailGun()

	m := mg.NewMessage(
		"ForthBox<support@forthbox.io>",
		subject, text, to,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}

func SendVerifyEmail(to string, code string, mtype string) error {
	tpl := "Welcome to ForthBox! Your email verify code is: %s"
	subject := "Please verify your email address"
	text := fmt.Sprintf(tpl, code)

	_, err := SendMail(to, subject, text)
	if err != nil {
		// TODO: log the error
		return err
	} else {
		return nil
	}
}


