package requests

type CreateEmailTemplate struct {
	Title     string `json:"title" binding:"required"`
	Subject   string `json:"subject" binding:"required"`
	HTML      string `json:"html" binding:"required"`
	PlainText string `json:"plainText" binding:"required"`
	Language  string `json:"language" binding:"required"`
	FromName  string `json:"fromName" binding:"required"`
	FromEmail string `json:"fromEmail" binding:"required"`
}

type UpdateEmailTemplate struct {
	Title     string `json:"title" binding:"required"`
	Subject   string `json:"subject" binding:"required"`
	HTML      string `json:"html" binding:"required"`
	PlainText string `json:"plainText" binding:"required"`
	Language  string `json:"language" binding:"required"`
	FromName  string `json:"fromName" binding:"required"`
	FromEmail string `json:"fromEmail" binding:"required"`
}

type SendEmail struct {
	TemplateTitle string   `json:"templateTitle" binding:"required"`
	ToName        string   `json:"toName" binding:"required"`
	ToEmail       string   `json:"toEmail" binding:"required"`
	Data          []string `json:"data" binding:"required"`
}
