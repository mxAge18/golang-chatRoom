package processes

import "fmt"

var (
	ServerUserManger *UserManger
)
type UserManger struct {
	onlineUsers map[string]*UserProcessor
}

func init() {
	ServerUserManger = &UserManger {
		onlineUsers : make(map[string]*UserProcessor, 1024),
	}
}

func (this *UserManger) addOnlineUser(up *UserProcessor) {
	this.onlineUsers[up.UserId] = up
}
func (this *UserManger) deleteOnlineUser(userId string) {
	delete(this.onlineUsers, userId)
}

func (this *UserManger) getAllOnlineUsers() map[string]*UserProcessor {
	return this.onlineUsers
}
func (this *UserManger) getOnlineUserWithUserId(userId string) (userProcess *UserProcessor, err error) {
	userProcess, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("user not exist or not online")
		return
	}
	return
}