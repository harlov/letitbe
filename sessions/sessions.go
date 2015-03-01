package sessions
import (
    "github.com/gin-gonic/gin"
    "github.com/harlov/letitbe/forms"
    "github.com/gin-gonic/gin/binding"
    "github.com/harlov/letitbe/security"
    "github.com/harlov/letitbe/db"
)

func SetupRoutes(router *gin.RouterGroup) {
    router.POST("/start", start_session)
    router.POST("/check", check_session)
}

func start_session(c *gin.Context) {
    var err error
    var  login_form forms.LoginForm
    c.BindWith(&login_form, binding.Form)

    if login_form.Username == "" || login_form.Password == "" {
        c.JSON(400, gin.H{"error": 1, "error_message": "required username and password",  "request": login_form })
        return
    }

    err, user_session := security.AuthUser(login_form)

    if err == nil {
        c.JSON(200, gin.H{"error" : 0, "session": user_session})
    } else {
        c.JSON(403, gin.H{"error" : 100, "session": user_session})
    }
}

func check_session(c *gin.Context) {
    var err error
    var checksession_form forms.CheckSession
    var session db.UserSession
    c.BindWith(&checksession_form, binding.Form)

    if checksession_form.SessionToken == "" {
        c.JSON(400, gin.H{"status" : "invalid format, check documentation"})
        return
    }

    err, session = db.FindSession(checksession_form.SessionToken)

    if err != nil {
        c.JSON(403, gin.H{"status" : err})
        return
    }


    c.JSON(200, gin.H{"status": "success", "session" : session})

}