package requests

type CreateACL struct {
	Handler   string `json:"handler" binding:"required"`
	IsPublic  bool   `json:"isPublic"`
	Title     string `json:"title" binding:"required"`
	Permitted string `json:"permitted" binding:"required"`
}

type UpdateACL struct {
	Handler   string `json:"handler" binding:"required"`
	IsPublic  bool   `json:"isPublic" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Permitted string `json:"permitted" binding:"required"`
}
