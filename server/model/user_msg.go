package model

type UserSmsMsg struct {
	User User
	Body string `json:"body"`
}