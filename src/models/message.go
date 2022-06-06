package models

import (
	"fmt"
	"time"
)

type AuthMessage struct {
	Status int
	Data   User
}

type ChatMessage struct {
	Username  string
	Data      string
	Timestamp int64
}

// ToString converts ChatMessage struct as string.
func (m *ChatMessage) ToString() string {
	t := time.Unix(0, m.Timestamp)
	return fmt.Sprintf("%d:%d, %s: %s", t.Hour(), t.Minute(), m.Username, m.Data)
}
