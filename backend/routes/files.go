package routes

import (
	"backend/models"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	originalname := file.Filename
	extension := strings.Split(originalname, ".")[1]
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
		Extension:    extension,
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

func GetFiles(c *gin.Context) {

}
