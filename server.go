package main
import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/harlov/letitbe/forms"
    "github.com/harlov/letitbe/db"
    "github.com/gin-gonic/gin/binding"
    "github.com/harlov/letitbe/security"
)



func main() {
    var err error
    log.Print("letitbe init...")
    err  = db.Init()
    if err != nil {
        log.Fatal("db connect error. break")
        return
    }
    log.Println("db inited...")

    createServer()
}

func createServer() {
    r := gin.Default()
    router := r.Group("/v1")

    r.GET("/", rootAction)

    setupRoutes(router)
    r.Run(":8080")
}


func setupRoutes(router *gin.RouterGroup) {
    router.POST("/login", loginAction)
    router.POST("/register", registerAction)
}



func loginAction(c *gin.Context) {
    var err error
    var  login_form forms.LoginForm
    c.BindWith(&login_form, binding.Form)

    if login_form.Username == "" || login_form.Password == "" {
        c.JSON(400, gin.H{"error": 1, "error_message": "required username and password",  "request": login_form })
        return
    }

    err, user_session := security.AuthUser(login_form)

    if err == nil {
        c.JSON(200, gin.H{"error" : 0, "request": login_form, "session": user_session})
    } else {
        c.JSON(403, gin.H{"error" : 100, "request": login_form, "session": user_session})
    }
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
    new_user.PassHash = register_form.Password
    new_user.Email = register_form.Email

    err, new_user := db.AddUser(new_user)
    if  err == nil {
        c.JSON(200, gin.H{"error" : 0, "request": register_form})
    }  else {
        c.JSON(400, gin.H{"error" : 1, "request": register_form, "error_sys": err})
    }
}

func rootAction(c *gin.Context) {
    c.String(200, "Auth&permissions microservice")
}