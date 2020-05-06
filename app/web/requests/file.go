package requests

type CreateDocument struct {
	Name        string `json:"name" binding:"required"`
	Path        string `json:"path" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
	Info        string `json:"info" binding:"required"`
	Cdn         bool   `json:"cdn" binding:"required"`
}

type UpdateDocument struct {
	Name        string `json:"name" binding:"required"`
	Path        string `json:"path" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
	Info        string `json:"info" binding:"required"`
	Cdn         bool   `json:"cdn" binding:"required"`
}
