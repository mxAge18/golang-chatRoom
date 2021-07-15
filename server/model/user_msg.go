package model

import "time"

type UserSmsMsg struct {
	From User
	To	string `json:"toUser"`
	Body string `json:"body"`
	PostTime time.Time `json:"postTime"`
}