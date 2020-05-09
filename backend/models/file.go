package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type File struct {
	Storedname   string
	Originalname string
	Extension    string
	Cookie       string
	ID           uuid.UUID
	CreatedAt    time.Time
}

func (f *File) Upload(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO files (storedname, originalname, extension, cookie) VALUES ($1, $2, $3, $4)", f.Storedname, f.Originalname, f.Extension, f.Cookie)

	return err
}
