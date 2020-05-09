package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GenerateLink(c *gin.Context) {
	cookie := fmt.Sprintf("%v", (c.MustGet("cookie")))
}
