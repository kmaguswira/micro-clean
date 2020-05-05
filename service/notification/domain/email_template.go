package domain

type EmailTemplate struct {
	ID        string
	Title     string
	Subject   string
	HTML      string
	PlainText string
	Language  string
	FromName  string
	FromEmail string
}
