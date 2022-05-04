package processor

import (
	"com.kuroro/usersCommu/common"
	"com.kuroro/usersCommu/server/model"
	"com.kuroro/usersCommu/server/utils"

	"encoding/json"
	"fmt"
	"net"
)

type UserProcessor struct {
	Conn net.Conn
	//表示当前的连接是哪个用户的
	UserId int
}

func (u *UserProcessor) NotifyOthers() {

	for id, up := range userMgr.onlineUsers {
		if id == u.UserId {
			continue
		}
		up.NotifyMeOnline(u.UserId)
	}

}

func (u *UserProcessor) NotifyMeOnline(userId int) (err error) {
	var resMsg common.Message
	resMsg.Type = common.UserStatusNotifyMesType

	var data common.UserStatusNotifyMes
	data.UserId = userId
	data.UserStatus = common.UserOnline

	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("序列化失败：", err)
		return
	}
	resMsg.Data = string(bytes)
	t := utils.Transfer{
		Conn: u.Conn,
	}
	err = t.WritePkg(resMsg)
	if err != nil {
		fmt.Println("通知上线失败:", err)
		return
	}
	return
}

func (u *UserProcessor) ProcessLogin(msg common.Message) (err error) {

	var data common.LoginMessage
	err = json.Unmarshal([]byte(msg.Data), &data)
	if err != nil {
		fmt.Println("json反序列化失败：", err)
		return
	}

	//查询redis中的用户信息
	user, err := model.MyUserDao.LoginCheck(data.Id, data.Password)

	fmt.Println(user)

	//构建返回消息
	var resMsg common.Message
	resMsg.Type = common.LoginResMesType
	var resData common.LoginResMessage

	if err != nil {
		if err == model.USER_EXIST_ERROR {
			resData.Code = 500
			//err.Error返回我们自定义的err中的text信息
			resData.Err = err.Error()
		} else if err == model.USER_PASSWORD_ERROR {
			resData.Code = 300
			resData.Err = err.Error()
		} else {
			resData.Code = 110
			resData.Err = err.Error()
		}

	} else {
		resData.Code = 200
		resData.Err = "登录成功"
		u.UserId = data.Id
		userMgr.AddOrModifyUp(u)
		//通知其他在线用户，我上线了
		u.NotifyOthers()

		for userId, _ := range userMgr.onlineUsers {

			resData.UsersId = append(resData.UsersId, userId)

		}

	}

	bytes, err := json.Marshal(resData)
	if err != nil {
		fmt.Println("序列化失败：", err)
		return
	}
	resMsg.Data = string(bytes)

	t := utils.Transfer{
		Conn: u.Conn,
	}
	err = t.WritePkg(resMsg)
	if err != nil {
		return
	}
	return

}

func (u *UserProcessor) ProcessRegister(msg common.Message) (err error) {

	var data common.RegisterMessage
	err = json.Unmarshal([]byte(msg.Data), &data)
	if err != nil {
		fmt.Println("json反序列化失败：", err)
		return
	}

	//注册redis中的用户信息
	err = model.MyUserDao.Register(data.User)

	//构建返回消息
	var resMsg common.Message
	resMsg.Type = common.RegisterResMesType
	var resData common.RegisterResMessage

	if err != nil {

		if err == model.USER_EXIST_ERROR {
			resData.Code = 400
			resData.Err = err.Error()

		} else {
			resData.Code = 100
			resData.Err = err.Error()
		}

	} else {
		resData.Code = 200
		resData.Err = "注册成功"
	}

	//返回消息
	bytes, err := json.Marshal(resData)
	if err != nil {
		fmt.Println("序列化失败：", err)
		return
	}
	resMsg.Data = string(bytes)

	t := utils.Transfer{
		Conn: u.Conn,
	}
	err = t.WritePkg(resMsg)
	if err != nil {
		return
	}
	return

}
