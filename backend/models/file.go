package models

import (
	"context"
	"time"

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
