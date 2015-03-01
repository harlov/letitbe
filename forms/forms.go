package forms

type LoginForm struct{
    Username string `form:"username" binding:"required"`
    Password string `form:"password" binding:"required"`
}

type RegisterForm struct {
    Username string `form:"username" binding:"required"`
    Email string `form:"email" binding:"required"`
    Password string `form:"password" binding:"required"`
}


type CheckSession struct {
    SessionToken string `form:"session_token" binding:"required" `
}


