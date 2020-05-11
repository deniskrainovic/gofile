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

	_, err := conn.Exec(context.Background(), "WITH lid AS (INSERT INTO links (folderpath, accesspassword, expirationdate, cookie) VALUES ($1, $2, $3, $4) RETURNING id,cookie) UPDATE files SET linkid = (SELECT id FROM lid) WHERE cookie = (SELECT cookie FROM lid)", l.Folderpath, l.Accesspassword, l.ExpirationDate, l.Cookie)

	return err
}

func (l *Link) CheckLink(conn *pgxpool.Pool) (bool, bool) {
	var exists bool = false
	var passwordSet bool = false
	var password string
	err := conn.QueryRow(context.Background(), "SELECT accesspassword FROM links WHERE cookie=$1", l.Cookie.String()).Scan(&password)
	if err != pgx.ErrNoRows {
		exists = true
	}
	if password != "" {
		passwordSet = true
	}
	return exists, passwordSet
}

//CheckPassword compares passwords and returns if its correct or not
func (l *Link) CheckPassword(conn *pgxpool.Pool) (bool, error) {
	isPasswordCorrect := false
	var dbPassword string
	err := conn.QueryRow(context.Background(), "SELECT accesspassword FROM links WHERE cookie=$1", l.Cookie.String()).Scan(&dbPassword)
	if dbPassword == l.Accesspassword {
		isPasswordCorrect = true
	}
	return isPasswordCorrect, err
}
