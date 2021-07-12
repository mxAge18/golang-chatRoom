package model

import (
	"chatPro/common/message"
	"encoding/json"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

type UserSmsMsgDao struct{
	Pool *redis.Pool
}
var (
	ThisUserMsgDao *UserSmsMsgDao
)

func NewThisUserMsgDao(pool *redis.Pool) (ThisUserMsgDao *UserSmsMsgDao) {
	ThisUserMsgDao = &UserSmsMsgDao{
		Pool: pool,
	}
	return
}
func (this *UserSmsMsgDao) buildMsg(msg message.SmsMsgSingle) (userMsg UserSmsMsg){
	userMsg.Body = msg.Body
	userMsg.From.UserId = msg.From.UserId
	userMsg.From.UserName = msg.From.UserName
	userMsg.To = msg.To
	userMsg.PostTime = time.Now()
	return
}

func (this *UserSmsMsgDao) setUnreadMsgCounter(conn redis.Conn, ToUserId string, unReadKey string) (err error){
	_, err = conn.Do("HEXISTS", ToUserId, unReadKey)
	fmt.Println("res,", err)
	if err != nil {
		num, _ := redis.Int(conn.Do("HGet", ToUserId, unReadKey))
		conn.Do("HSet", ToUserId, unReadKey, num + 1)
	} else {
		conn.Do("HSet", ToUserId, unReadKey, 1)
	}
	return
}

func (this *UserSmsMsgDao) generateKey(fromId string, toId string) (unReadKey string) {
	return toId + "_" + fromId
}


func (this *UserSmsMsgDao) StoreUnreadMsg(msg message.SmsMsgSingle) (err error) {
	fmt.Println("store unread to redis")
	conn := this.Pool.Get()
	defer conn.Close()
	userMsg := this.buildMsg(msg)
	unReadKey := this.generateKey(userMsg.From.UserId, msg.To)
	data, err := json.Marshal(userMsg)
	if err != nil {
		fmt.Println("json.Marshal(userMsg) error", err)
		return err
	}
	fmt.Println("key", unReadKey)
	_, err = conn.Do("LPush", unReadKey, string(data))
	if err != nil {
		fmt.Println("\"LPush\", unReadKey, string(data)", err)
		return err
	}
	this.setUnreadMsgCounter(conn, userMsg.To, unReadKey)
	return err
}


func (this *UserSmsMsgDao) GetUnreadMsg() {

}