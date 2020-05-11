package main

import (
	"context"
	"fmt"

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
	router.Use(middleware.CookieMiddleware())
	router.Use(middleware.DbMiddleware(*conn))
	router.Use(middleware.CORSMiddleware())

	router.POST("uploads/:uploadID/download", routes.DownloadAllFiles)
	router.POST("uploads/:uploadID", routes.PostLink)
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
