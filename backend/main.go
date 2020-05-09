package main

import (
	"context"
	"fmt"

	"backend/middleware"
	"backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	router := gin.Default()
	router.Use(middleware.CookieMiddleware())
	uploadsGroup := router.Group("uploads")
	{
		uploadsGroup.POST("file", middleware.DbMiddleware(*conn), routes.UploadsPostFile)
	}
	router.Run(":8080")
}

func connectDB() (c *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:Test123@localhost:5432/gofile")
	if err != nil {
		fmt.Println("Error connecting to DB")
		fmt.Println(err.Error())
	}
	err = conn.Ping(context.Background())
	return conn, err
}
