package main

import (
	"context"
	"fmt"
	"net/http"

	"backend/middleware"
	"backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	conn, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Use(middleware.CookieMiddleware())
	router.Use(middleware.DbMiddleware(*conn))
	router.Use(middleware.CORSMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.Group("/uploads")
	{
		router.GET(":uploadID", routes.GetLink, routes.GetLinkFiles)
	}

	router.POST("upload/file", routes.UploadFile)
	router.POST("link/generate", routes.GenerateLink)

	router.Run(":8080")
}

func connectDB() (c *pgxpool.Pool, err error) {
	conn, err := pgxpool.Connect(context.Background(), "postgresql://postgres:@localhost:5432/gofile")
	if err != nil {
		fmt.Println("Error connecting to DB")
		fmt.Println(err.Error())
	}
	return conn, err
}
