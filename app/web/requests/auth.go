package requests

type SignIn struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSignIn struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
	Token    string `json:"token" binding:"required"`
}

type SignUp struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
}
