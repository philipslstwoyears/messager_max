package repository

import (
	"awesomeProject2/models"
)

func (s *Storage) SaveUser(user models.User) error {
	q := `INSERT INTO users (login, password) VALUES ($1, $2)`
	_, err := s.Db.Exec(q, user.Login, user.Password)
	return err
}
