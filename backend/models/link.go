package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type link struct {
	CreatedAt      time.Time
	ID             uuid.UUID
	Foldername     string
	Accesspassword string
	ExpirationDate time.Time
	FileID         uuid.UUID
}

func (l *link) Generate(conn *pgx.Conn) string {
	return ""
}
