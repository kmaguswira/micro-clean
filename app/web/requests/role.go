package requests

type CreateRole struct {
	Title string `json:"title" binding:"required"`
}

type UpdateRole struct {
	Title string `json:"title" binding:"required"`
}
