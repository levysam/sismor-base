package auth

import "github.com/golang-jwt/jwt"

type LoginForm struct {
	Email    string `json:"email" xml:"email" form:"email"`
	Password string `json:"password" xml:"password" form:"password"`
}

type UserClaims struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	jwt.StandardClaims
}
