package gomail_test

import (
	"os"
	"testing"

	"github.com/avila-r/gomail"
	"github.com/joho/godotenv"
)

func Test_Send(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Errorf(err.Error())
	}

	client := gomail.NewClient(os.Getenv("SENDGRID_KEY"))

	email := &gomail.Email{
		From: &gomail.User{
			Name:    "Sender",
			Address: "avila.dev@outlook.com", // From's email must match w/ sendgrid's sender
		},

		To: &gomail.User{
			Name:    "Remittee",
			Address: "r.souza.avila@outlook.com",
		},

		Subject:          "Assunto do Email",
		PlainTextContent: "Este é um e-mail enviado com SendGrid!",
		HtmlContent:      "<strong>Este é um e-mail enviado com SendGrid!</strong>",
	}

	_, err := client.Send(*email)

	if err != nil {
		t.Errorf(err.Error())
	}
}
