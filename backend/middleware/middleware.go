package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func DbMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

func CookieMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session")
		if err != nil {
			cookie = uuid.New().String()
			c.SetCookie("session", cookie, 86400, "/", "localhost", false, true)
		}
		c.Set("cookie", cookie)
		c.Next()
	}
}
