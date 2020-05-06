package requests

type CreateImage struct {
	Name        string `json:"name" binding:"required"`
	Path        string `json:"path" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Thumbnail   string `json:"thumbnail" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Alt         string `json:"alt" binding:"required"`
	Description string `json:"description" binding:"required"`
	Info        string `json:"info" binding:"required"`
	Cdn         bool   `json:"cdn" binding:"required"`
}

type UpdateImage struct {
	Name        string `json:"name" binding:"required"`
	Path        string `json:"path" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Thumbnail   string `json:"thumbnail" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Alt         string `json:"alt" binding:"required"`
	Description string `json:"description" binding:"required"`
	Info        string `json:"info" binding:"required"`
	Cdn         bool   `json:"cdn" binding:"required"`
}
