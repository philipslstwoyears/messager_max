package repository

import (
	"awesomeProject2/models"
)

func (s *Storage) SaveUser(user models.User) (int, error) {
	q := `INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id`
	res := s.Db.QueryRow(q, user.Login, user.Password)
	var id int
	err := res.Scan(&id)
	return id, err
}
