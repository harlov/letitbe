package sessions
import (
    "testing"
    "github.com/gin-gonic/gin"
)


func TestSetupRoutes(t *testing.T) {
    r := gin.Default()
    router := r.Group("/v1")

    session_router := router.Group("/session")

    session_router.POST("/start", start_session)
    session_router.POST("/check", check_session)
}