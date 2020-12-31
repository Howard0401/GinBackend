package utils

import (
	settings "VueGin/config/settingModels"
	"VueGin/global"
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type Email struct {
	info *settings.SMTP
}

// info *settings.SMTPInfo
func NewEmail() *Email {
	return &Email{
		info: &settings.SMTP{
			Host:     global.Global_Config.SMTP.Host,
			Port:     global.Global_Config.SMTP.Port,
			SSL:      global.Global_Config.SMTP.SSL,
			UserName: global.Global_Config.SMTP.UserName,
			Password: global.Global_Config.SMTP.Password,
			From:     global.Global_Config.SMTP.From,
			To:       global.Global_Config.SMTP.To,
		},
	}
}

func (e *Email) SendEmail(to string, subject, body string) error {
	//引入gomail包
	//若報錯 Invalid login: Application-specific password required，可參考以下：
	//Gmail驗證:https://stackoverflow.com/questions/60701936/error-invalid-login-application-specific-password-required
	//範例 https://github.com/go-gomail/gomail/issues/82
	m := gomail.NewMessage()
	m.SetHeader("From", e.info.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(e.info.Host, e.info.Port, e.info.UserName, e.info.Password)
	// Insecure僅為測試用
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.info.SSL}
	err := dialer.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
