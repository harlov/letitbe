package security
import (
    "github.com/harlov/letitbe/forms"
    "github.com/harlov/letitbe/db"

)


func compare_pass(orig, input string) bool {
    return orig == input
}


func AuthUser(loginForm forms.LoginForm) (error, db.UserSession ) {
    var err error
    err, user_entity := db.FindUser(loginForm.Username)
    var user_session db.UserSession
    if user_entity.PassHash == loginForm.Password {
        user_session,err = db.StartSession(user_entity)
    }

    return err, user_session
}