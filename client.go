package gomail

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Client struct {
	Key string
	ctx *sendgrid.Client
}

func NewClient(key string) *Client {
	client := &Client{
		Key: key,
		ctx: sendgrid.NewSendClient(key),
	}

	return client
}

func (c *Client) Send(e Email) (*Response, error) {
	from, to := mail.NewEmail(e.From.Name, e.From.Address), mail.NewEmail(e.To.Name, e.To.Address)

	email := mail.NewSingleEmail(from, e.Subject, to, e.PlainTextContent, e.HtmlContent)

	r, err := c.ctx.Send(email)

	response := &Response{
		r.StatusCode,
		r.Body,
		r.Headers,
	}

	return response, err
}
