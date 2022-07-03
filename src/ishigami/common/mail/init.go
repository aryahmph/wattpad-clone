package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

type MailManager struct {
	Email    string
	Password string
	SmtpHost string
	SmtpPort int
}

func NewMailManager(email string, password string, host string, port int) *MailManager {
	return &MailManager{Email: email, Password: password, SmtpHost: host, SmtpPort: port}
}

func (m MailManager) SendMail(receiver Receiver) (err error) {
	// Set up authentication information.
	//smtp.PlainAuth("", m.Email, m.Password, m.SmtpHost)
	auth := smtp.CRAMMD5Auth(m.Email, m.Password)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	bodyMessage := "To: " + strings.Join(receiver.To, ",") + "\n" +
		"Cc: " + strings.Join(receiver.Cc, ",") + "\n" +
		"Subject: " + receiver.Subject + "\n\n" + receiver.Message
	smtpAddr := fmt.Sprintf("%s:%d", m.SmtpHost, m.SmtpPort)

	err = smtp.SendMail(smtpAddr, auth, m.Email, append(receiver.To, receiver.Cc...), []byte(bodyMessage))
	return err
}
