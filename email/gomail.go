// simple mail transfer protocol demo based on "gomail.v2"
package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"instance.golang.com/utils"
)

func Gomail() {
	// authentication config
	ec := NewEmailConfig(
		"smtp.163.com",    // email server host
		25,                // email server port
		"example@163.com", // your email account
		"******",          // authorization code not password
	)

	// email
	subject := utils.Now() + " [Hello]"
	body := "Verification code: " + utils.RandNumMath() + "\n\nThanks!"
	//link := "Hello <a href = \"http://www.google.com\">google</a>"

	fmt.Println("#______________start sending test mail...")
	if err := sendEmail(ec, subject, body); err != nil {
		fmt.Printf("#______________send failed: %v\n", err)
	} else {
		fmt.Println("#______________send success")
	}
}

func sendEmail(ec *EmailConfig, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", ec.Username)
	m.SetHeader("To", "3038777418@qq.com")
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	//m.SetBody("text/html", body)
	return gomail.NewDialer(ec.Host, ec.Port, ec.Username, ec.Password).DialAndSend(m)
}
