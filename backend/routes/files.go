package routes

import (
	"backend/models"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	originalname := file.Filename
	storedname := uuid.New().String()
	cookie := fmt.Sprintf("%v", (c.MustGet("cookie")))

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	wd, err := os.Getwd()

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	path := filepath.Join(wd, "uploads", cookie)
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	path = filepath.Join(path, storedname)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	db, _ := c.Get("db")
	conn := db.(pgxpool.Pool)
	fileInfo := models.File{
		Storedname:   storedname,
		Originalname: originalname,
		Cookie:       cookie,
	}
	err = fileInfo.WriteToDB(&conn)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func DownloadAllFiles(c *gin.Context) {
	uploadID := c.Param("uploadID")
	link := models.Link{}
	err := c.ShouldBind(&link)
	id, err := uuid.Parse(uploadID)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	link.Cookie = id
	db, _ := c.Get("db")
	conn := db.(pgxpool.Pool)
	linkExists, passwordSet := link.CheckLink(&conn)
	if linkExists != true {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if passwordSet == false {
		c.Writer.Header().Set("Content-type", "application/octet-stream")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=gofile_download.zip")
		models.DownloadAllFilesByCookie(&conn, link.Cookie.String(), &c.Writer)
		return
	}
	if link.Accesspassword == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	isPasswordCorrect, err := link.CheckPassword(&conn)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if isPasswordCorrect == true {
		c.Writer.Header().Set("Content-type", "application/octet-stream")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=gofile_download.zip")
		models.DownloadAllFilesByCookie(&conn, link.Cookie.String(), &c.Writer)
		return
	}
	c.AbortWithStatus(http.StatusUnauthorized)
	return

}
