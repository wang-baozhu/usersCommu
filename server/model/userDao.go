package model

import (
	"com.kuroro/usersCommu/common"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var MyUserDao *UserDao

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) *UserDao {
	return &UserDao{
		pool: pool,
	}
}

//用户登录校验
//1.根据ID查询用户
//2.校验密码

func (u *UserDao) LoginCheck(id int, password string) (user common.User, err error) {

	conn := u.pool.Get()
	defer conn.Close()

	user, err = u.GetById(conn, id)

	if err != nil {
		return
	}

	if user.Password != password {
		err = USER_PASSWORD_ERROR
		return
	}

	return

}

//根据用户Id到redis中查询是否有

func (u *UserDao) GetById(conn redis.Conn, id int) (user common.User, err error) {

	r, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = USER_NOT_EXIST_ERROR
		}
		return
	}

	//反序列化
	err = json.Unmarshal([]byte(r), &user)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return
	}
	return
}

func (u *UserDao) Register(user common.User) (err error) {
	conn := u.pool.Get()
	defer conn.Close()

	_, err = u.GetById(conn, user.UserId)

	if err == nil {
		err = USER_EXIST_ERROR
		return
	}
	if err != USER_NOT_EXIST_ERROR {
		return
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json序列化失败：", err)
		return
	}

	_, err = conn.Do("HSet", "users", user.UserId, string(bytes))
	if err != nil {
		fmt.Println("redis注册用户失败：", err)
		return
	}
	return

}
