package requests

type CreateUser struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	RoleID   string `json:"roleId" binding:"required"`
	Status   string `json:"status" binding:"required"`
	Password string `json:"password" binding:"required"`
}
