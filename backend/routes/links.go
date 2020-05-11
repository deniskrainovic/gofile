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

func GenerateLink(c *gin.Context) {
	link := models.Link{}
	err := c.ShouldBind(&link)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	link.Cookie, err = uuid.Parse(fmt.Sprintf("%v", (c.MustGet("cookie"))))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	db, _ := c.Get("db")
	conn := db.(pgxpool.Pool)
	link.Folderpath = filepath.Join(wd, "uploads", link.Cookie.String())
	err = link.WriteToDB(&conn)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	cookie := uuid.New().String()
	c.SetCookie("session", cookie, 86400, "/", "localhost", false, true)

	url := "http://localhost:8080/uploads/" + link.Cookie.String()
	c.JSON(http.StatusOK, gin.H{"link": url})
}

func PostLink(c *gin.Context) {
	uploadID := c.Param("uploadID")
	link := models.Link{}
	err := c.ShouldBind(&link)
	var files []models.File
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
		files, err = models.GetAllFilesListByCookie(&conn, uploadID)
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	} else {
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
			files, err = models.GetAllFilesListByCookie(&conn, uploadID)
			if err != nil {
				fmt.Println(err.Error())
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"isPasswordNeeded": passwordSet, "files": files})
	return
}
