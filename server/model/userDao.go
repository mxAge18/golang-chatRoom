package model

import (
	"encoding/json"
	"fmt"
	"chatPro/common/message"

	"github.com/garyburd/redigo/redis"
)

type UserDBO struct {
	pool *redis.Pool
}

var (
	MyUserDBO *UserDBO
)

func NewUserDBO(pool *redis.Pool) (userDBO *UserDBO) {
	userDBO = &UserDBO{
		pool: pool,
	}
	return
}

func (this *UserDBO) create(userName, userId, userPwd string)(user *User, err error) {
	return
}
func (this *UserDBO) update() {

}
func (this *UserDBO) delete() {

}
func (this *UserDBO) getByFiledUserId(conn redis.Conn, userId string) (user *User, err error) {
	result, err := redis.String(conn.Do("HGet", "users", userId))
	fmt.Println(result)
	if err != nil {
		if redis.ErrNil == err {
			err = ERROR_USER_NOT_EXIST
		}
		return
	}
	user = &User{}
	// 对redis 结果反序列化
	err = json.Unmarshal([]byte(result), &user)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(result), &user) error", err)
		return
	}
	return
}

func (this *UserDBO) Login(userId string, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getByFiledUserId(conn, userId)
	if err != nil {
		return
	}
	if userPwd != user.UserPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDBO) Register(user *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getByFiledUserId(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXIST
		return
	} else {
		data, err := json.Marshal(user)
		if err != nil {
			fmt.Println("json.Marshal(user) error", err)
			return err
		}
		_, err = conn.Do("HSet", "users", user.UserId, string(data))
		if err != nil {
			fmt.Println("\"HSet\", \"users\", user.UserId, string(data)", err)
			return err
		}
		return err
	}

}
func (this *UserDBO) getByName() {

}
func (this *UserDBO) getAll() {
}
