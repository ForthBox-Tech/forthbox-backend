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

