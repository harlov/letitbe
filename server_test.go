package main
import (
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/harlov/letitbe/users"
    "github.com/harlov/letitbe/sessions"
)

func TestSetupRoutes(t *testing.T) {
    r := gin.Default()
    router := r.Group("/v1")

    r.GET("/", rootAction)
    user_router := router.Group("/user")
    session_router := router.Group("/session")

    users.SetupRoutes(user_router)
    sessions.SetupRoutes(session_router)
}

func TestDummy(t *testing.T) {
    /*resp, err := http.Get("http://127.0.0.1:8080/")
    if err != nil {
        t.Error("error get / page")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if string(body) != "hello" {
        t.Error("not valid answer on / page")
    }*/
}