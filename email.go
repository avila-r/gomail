package gomail

type User struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"email,omitempty"`
}

type Email struct {
	From             *User
	Subject          string
	To               *User
	PlainTextContent string
	HtmlContent      string
}
