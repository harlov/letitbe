package security
import (
    "github.com/harlov/letitbe/forms"
    "github.com/harlov/letitbe/db"
    "crypto/sha512"

    "encoding/hex"
    "errors"
)



func HashPass(pass string )(string) {
    h := sha512.New()
    h.Write([]byte(pass))
    return hex.EncodeToString(h.Sum(nil))
}


func compare_pass(orig, input string) bool {
    return orig == HashPass(input)

}


func AuthUser(loginForm forms.LoginForm) (error, db.UserSession ) {
    var err error
    var user_session db.UserSession
    err, user_entity := db.FindUser(loginForm.Username)
    if err != nil {
        return err, user_session
    }


    if compare_pass(user_entity.PassHash, loginForm.Password) {
        user_session,err = db.StartSession(user_entity, 86400)
    } else {
        err = errors.New("pass_fail")
    }

    return err, user_session
}