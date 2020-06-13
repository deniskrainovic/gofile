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
	// Get/Set cookie and bind to request
	router.Use(middleware.CookieMiddleware())
	// Bind db session to request
	router.Use(middleware.DbMiddleware(*conn))
	// Enable CORS for react
	router.Use(middleware.CORSMiddleware())

	router.POST("api/uploads/:uploadID/download", routes.DownloadAllFiles)
	router.GET("api/uploads/:uploadID", routes.GetLinkFiles)
	router.GET("api/uploads/:uploadID/checkpassword", routes.GetCheckIfPasswordNeeded)
	router.POST("api/upload/file", routes.UploadFile)
	router.POST("api/link/generate", routes.GenerateLink)

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
