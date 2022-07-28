package auth

type LoginForm struct {
	Email    string `json:"email" xml:"name" form:"name"`
	Password string `json:"password" xml:"password" form:"password"`
}
