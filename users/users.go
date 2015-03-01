package users
import (
    "github.com/gin-gonic/gin"
    "github.com/harlov/letitbe/forms"
    "github.com/gin-gonic/gin/binding"
    "github.com/harlov/letitbe/db"
    "github.com/harlov/letitbe/security"
)

func SetupRoutes(router *gin.RouterGroup) {
    router.POST("/register", registerAction)
}


func registerAction(c *gin.Context) {
    var register_form forms.RegisterForm
    c.BindWith(&register_form, binding.Form)
    if  register_form.Username == "" ||
    register_form.Password == "" ||
    register_form.Email == "" {
        c.JSON(400, gin.H{"error": 1, "error_message": "required username,password and email"})
    }

    var new_user db.UserEntity
    new_user.Username = register_form.Username
    new_user.PassHash = security.HashPass(register_form.Password)
    new_user.Email = register_form.Email

    err, new_user := db.AddUser(new_user)
    if  err == nil {
        c.JSON(200, gin.H{"error" : 0, "request": register_form})
    }  else {
        c.JSON(400, gin.H{"error" : 1, "request": register_form, "error_sys": err})
    }
}