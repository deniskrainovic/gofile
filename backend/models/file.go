package models

import (
	"archive/zip"
	"context"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type File struct {
	Storedname   string `json:"-"`
	Originalname string
	Cookie       string    `json:"-"`
	ID           uuid.UUID `json:"-"`
	CreatedAt    time.Time `json:"-"`
}

//WriteToDB writes the file-info to the database
func (f *File) WriteToDB(conn *pgxpool.Pool) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO files (storedname, originalname, cookie) VALUES ($1, $2, $3)", f.Storedname, f.Originalname, f.Cookie)

	return err
}

func GetAllFilesListByCookie(conn *pgxpool.Pool, cookie string) ([]File, error) {
	rows, err := conn.Query(context.Background(), "SELECT id,originalname FROM files WHERE Cookie = $1", cookie)
	var files []File
	for rows.Next() {
		var file File
		err = rows.Scan(&file.ID, &file.Originalname)
		files = append(files, file)
	}
	return files, err
}

func DownloadAllFilesByCookie(conn *pgxpool.Pool, cookie string, writer *gin.ResponseWriter) error {
	rows, err := conn.Query(context.Background(), "SELECT storedname,originalname FROM files WHERE Cookie = $1", cookie)

	wd, _ := os.Getwd()
	ar := zip.NewWriter(*writer)

	for rows.Next() {
		path := filepath.Join(wd, "uploads", cookie)
		var storedname string
		var originalname string

		err = rows.Scan(&storedname, &originalname)
		path = filepath.Join(path, storedname)
		file, _ := os.Open(path)
		f, _ := ar.Create(originalname)
		io.Copy(f, file)
	}

	ar.Close()
	return err
}
