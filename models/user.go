package models

type User struct {
	Base

	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

type RegisterUser struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Password string `json:"password"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}
