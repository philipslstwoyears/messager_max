package models

import (
	"errors"
	"time"
)

type Message struct {
	CreatedAt  time.Time `json:"created_at"`
	SenderID   int       `json:"sender_id"`
	ReceiverID int       `json:"receiver_id"`
	Content    string    `json:"content"`
	ID         int       `json:"id"`
}

func (m *Message) Valid() error {
	if m.SenderID == 0 {
		return errors.New("SenderID is zero")
	}
	if m.ReceiverID == 0 {
		return errors.New("ReceiverID is zero")
	}
	if m.Content == "" {
		return errors.New("Content is empty")
	}
	return nil
}
