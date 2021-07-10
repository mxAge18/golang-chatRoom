package model

import (
	"chatPro/common/message"

	"github.com/garyburd/redigo/redis"
)

type UserSmsMsgDao struct{
	p *redis.Pool
}

func (this *UserSmsMsgDao) StoreUnreadMsg(msg message.Message) (err error) {

	return
}

func (this *UserSmsMsgDao) GetUnreadMsg() {

}