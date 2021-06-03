package processes

import (
	"fmt"
	"go_code/chatPro/common/message"
)

var (
	ClientUserMangerObj *ClientUserManger
)

type ClientUserManger struct {
	onlineUsers map[string]*message.User
}

func init() {
	ClientUserMangerObj = &ClientUserManger{
		onlineUsers: make(map[string]*message.User, 10),
	}
}

func (this *ClientUserManger) OutputOnlineUsers() {
	for _, v := range this.onlineUsers {
		fmt.Println("userName:", v.UserName)
	}
}

func (this *ClientUserManger) AddOnlineUser(user *message.User) {
	_, ok := this.onlineUsers[user.UserId]
	fmt.Println("ok,user", user)
	if !ok {
		this.onlineUsers[user.UserId] = user
	}
}

func (this *ClientUserManger) DeleteOnlineUser(userId string) {
	delete(this.onlineUsers, userId)
}

func (this *ClientUserManger) GetAllOnlineUsers() map[string]*message.User {
	return this.onlineUsers
}
func (this *ClientUserManger) GetOnlineUserWithUserId(userId string) (user *message.User, err error) {
	user, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("user not exist or not online")
		return
	}
	return
}
