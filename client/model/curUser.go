package model

import (
	"com.kuroro/usersCommu/common"
	"net"
)

//登录成功以后进行初始化，用于群发消息
var (
	CurrentUser CurUser
)

type CurUser struct {
	User common.User
	Conn net.Conn
}
