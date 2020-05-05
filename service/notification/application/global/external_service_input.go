package global

type SendEmailInput struct {
	To               EmailAddress
	From             EmailAddress
	Subject          string
	PlainTextContent string
	HTML             string
}

type EmailAddress struct {
	Name  string
	Email string
}
