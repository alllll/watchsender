package sender

import (
	"fmt"
	"net/mail"
	"net/smtp"

	"github.com/alllll/watchsender/models"
	"github.com/scorredoira/email"
)

func Sender(c chan models.Email, smtphost string, smtpport string, smtplogin string, smtppassword string, from string) {

	for s := range c {
		//горутина отправки
		go Send(s, smtphost, smtpport, smtplogin, smtppassword, from)
	}
}

func Send(message models.Email, smtphost string, smtpport string, smtplogin string, smtppassword string, from string) {
	m := email.NewMessage(message.Subject, message.Body)
	m.From = mail.Address{Address: from}
	m.To = []string{message.To}
	for _, attach := range message.Attaches {
		m.AttachBuffer(attach.Name, attach.File, false)
	}
	err := email.Send(smtphost+":"+smtpport, smtp.PlainAuth("", smtplogin, smtppassword, smtphost), m)
	if err != nil {
		fmt.Println(err)
	}

}
