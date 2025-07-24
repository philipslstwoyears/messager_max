package repository

import (
	"awesomeProject2/models"
	"database/sql"
	"errors"
)

type Storage struct {
	Db *sql.DB
}

func NewDb(db *sql.DB) *Storage {
	return &Storage{
		Db: db,
	}
}

func (s *Storage) AddMessage(message models.Message) (int, error) {
	var q = `
	INSERT INTO messages (receiver, content, sender) values ($1, $2, $3) RETURNING id;
	`
	result := s.Db.QueryRow(q, message.ReceiverID, message.Content, message.SenderID)
	var id int
	err := result.Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *Storage) GetMessage(id int) (models.Message, bool) {
	message := models.Message{}
	var q = `
	SELECT id, receiver, content, sender, created_at FROM messages WHERE id = $1
	`
	err := s.Db.QueryRow(q, id).Scan(&message.ID, &message.ReceiverID, &message.Content, &message.SenderID, &message.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Message{}, false
		}
		return models.Message{}, false
	}
	return message, true

}

func (s *Storage) GetAll() []models.Message {
	rows, err := s.Db.Query("SELECT id, receiver, content, sender, created_at FROM messages")
	if err != nil {
		return []models.Message{}
	}
	defer rows.Close()
	messages := []models.Message{}
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.ID, &m.ReceiverID, &m.Content, &m.SenderID, &m.CreatedAt); err != nil {
			return []models.Message{}
		}
		messages = append(messages, m)
	}
	if rows.Err() != nil {
		return []models.Message{}
	}
	return messages
}

func (s *Storage) DeleteMessage(id int) error {
	result, err := s.Db.Exec("DELETE FROM messages WHERE id = ?", id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("message not found")
	}
	return nil
}
