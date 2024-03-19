package request

type UserRequest struct { // != gin -> validator
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=32,containsany=@#!$_"`
	Nickname string `json:"nickname" binding:"min=1,max=32"`
	Username string `json:"username" binding:"required,min=3,max=32"`
}
