package requests

type CreateEmailTemplate struct {
	Title    string `json:"title" binding:"required"`
	Subject  string `json:"subject" binding:"required"`
	Body     string `json:"body" binding:"required"`
	Language string `json:"language" binding:"required"`
}

type UpdateEmailTemplate struct {
	Title    string `json:"title" binding:"required"`
	Subject  string `json:"subject" binding:"required"`
	Body     string `json:"body" binding:"required"`
	Language string `json:"language" binding:"required"`
}
