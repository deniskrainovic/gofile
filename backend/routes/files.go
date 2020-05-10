package routes

import (
	"archive/zip"
	"backend/models"
	"fmt"
	"io"
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
	err := c.ShouldBindJSON(&link)
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
		//serve files
	} else {
		if link.Accesspassword == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		isPasswordCorrect, err := link.CheckPassword(&conn)
		if err != nil {
			fmt.Println("Here")
			fmt.Println(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if isPasswordCorrect == true {
			//serve files
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func ServeTest(c *gin.Context) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "uploads")
	path = filepath.Join(path, "1.txt")

	c.Writer.Header().Set("Content-type", "application/octet-stream")
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=filename.zip")
	ar := zip.NewWriter(c.Writer)
	path = filepath.Join(path, "1.txt")
	file1, _ := os.Open("filename1")
	path = filepath.Join(path, "2.txt")
	file2, _ := os.Open("filename2")
	path = filepath.Join(path, "1.txt")
	f1, _ := ar.Create("filename1")
	io.Copy(f1, file1)
	path = filepath.Join(path, "2.txt")
	f2, _ := ar.Create("filename2")
	io.Copy(f2, file2)
	ar.Close()
}
