package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Link struct {
	CreatedAt      time.Time
	ID             uuid.UUID
	Folderpath     string
	Cookie         uuid.UUID
	Accesspassword string `json:"password" form:"password"`
	ExpirationDate time.Time
	ExpirationDays *int `json:"expirationDays" form:"expirationDays" binding:"required,min=1,max=7"`
}

//WriteToDB writes the generated link to the database
// TODO: Implement Bcrypt
func (l *Link) WriteToDB(conn *pgxpool.Pool) error {
	l.ExpirationDate = time.Now().Local().Add(time.Hour * time.Duration(*l.ExpirationDays*24))

	_, err := conn.Exec(context.Background(), "WITH lid AS (INSERT INTO links (folderpath, accesspassword, expirationdate, cookie) VALUES ($1, $2, $3, $4) RETURNING id,cookie) UPDATE files SET link = (SELECT id FROM lid) WHERE cookie = (SELECT cookie FROM lid)", l.Folderpath, l.Accesspassword, l.ExpirationDate, l.Cookie)

	return err
}

func (l *Link) CheckLink(conn *pgxpool.Pool) (bool, bool) {
	var exists bool = false
	var passwordSet bool = false
	err := conn.QueryRow(context.Background(), "SELECT password FROM links WHERE id=$1", l.ID.String()).Scan(&l.Accesspassword)
	if err != pgx.ErrNoRows {
		exists = true
	}
	if l.Accesspassword != "" {
		passwordSet = true
	}
	return exists, passwordSet
}
