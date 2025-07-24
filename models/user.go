package models

import (
	"errors"
	"time"
)

type User struct {
	CreatedAt time.Time `json:"created_at"`
	Password  string    `json:"receiver"`
	Login     string    `json:"content"`
	ID        int       `json:"id"`
}

func (u *User) Valid() error {
	if len(u.Password) < 4 {
		return errors.New("Password must be at least 4 characters")
	}
	if len(u.Login) <= 3 {
		return errors.New("Login should be at least 3 characters")
	}

	return nil
}
