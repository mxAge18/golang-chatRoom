package model

import (
	"chatPro/common/message"
	"net"
)

type CurrentUser struct {
	Conn net.Conn
	message.User
}