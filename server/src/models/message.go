package models

type AuthMessage struct {
	Status int
	Data   User
}

type ChatMessage struct {
	Username  string
	Data      string
	Timestamp int64
}
