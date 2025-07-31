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

func (s *Storage) GetUserByLogin(login string) (models.User, error) {
	var user models.User
	query := `SELECT id, login, password FROM users WHERE login = $1`
	row := s.Db.QueryRow(query, login)
	err := row.Scan(&user.ID, &user.Login, &user.Password)
	return user, err
}
