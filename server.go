package main
import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/harlov/letitbe/db"
    "github.com/harlov/letitbe/sessions"
    "github.com/harlov/letitbe/users"
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
    user_router := router.Group("/user")
    session_router := router.Group("/session")

    users.SetupRoutes(user_router)
    sessions.SetupRoutes(session_router)
}

func rootAction(c *gin.Context) {
    c.String(200, "Auth&permissions microservice")
}