package model

import (
	"go_code/chatPro/common/message"
	"net"
)

type CurrentUser struct {
	Conn net.Conn
	message.User
}