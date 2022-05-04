package processor

import (
	"com.kuroro/usersCommu/client/model"
	"com.kuroro/usersCommu/client/utils"
	"com.kuroro/usersCommu/common"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcessor struct {
}

//校验用户登录信息
func (u *UserProcessor) Login(id int, username string, password string) (err error) {
	//1.建立连接
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("拨号失败：", err)
		return
	}
	defer conn.Close()

	//2..构建消息
	msg := common.Message{}
	msg.Type = common.LoginMesType
	bytes, err := json.Marshal(common.LoginMessage{id, username, password})
	if err != nil {
		fmt.Println("json序列化失败:", err)
		return
	}
	msg.Data = string(bytes)

	t := utils.Transfer{
		Conn: conn,
	}
	//3..发送消息
	err = t.WritePkg(msg)
	if err != nil {
		fmt.Println("客户端发送消息失败：", err)
	}
	fmt.Println("客户端发送消息成功：", string(bytes))

	//4..接收数据
	resMsg, err := t.ReadPkg()
	if err != nil {
		fmt.Println("客户端接收数据失败：", err)
		return err
	}
	fmt.Println("客户端接收到消息：", resMsg)

	//5..反序列化拿到状态码
	var m common.LoginResMessage
	json.Unmarshal([]byte(resMsg.Data), &m)
	if m.Code == 200 {
		fmt.Println("当前在线用户列表如下")
		for _, val := range m.UsersId {
			if val == id {
				continue
			}
			fmt.Println("在线用户ID：", val)
			user := common.User{
				UserId:     val,
				UserStatus: common.UserOnline,
			}
			UserOnlineMap[val] = user

		}

		go MonitorServer(conn)
		//为登录成功后的功能做准备
		c := common.User{
			UserId:     id,
			UserName:   username,
			UserStatus: common.UserOnline,
		}
		model.CurrentUser = model.CurUser{
			User: c,
			Conn: conn,
		}

		for {
			ShowMenu()
		}
	} else {
		fmt.Println(m.Err)

	}
	return
}

func (u *UserProcessor) Register(id int, username string, password string) (err error) {

	//1.建立连接
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("拨号失败：", err)
		return
	}
	defer conn.Close()

	//2..构建消息
	msg := common.Message{}
	msg.Type = common.RegisterMesType
	user := common.User{UserId: id, UserName: username, Password: password}
	bytes, err := json.Marshal(common.RegisterMessage{user})
	if err != nil {
		fmt.Println("json序列化失败:", err)
		return
	}
	msg.Data = string(bytes)
	fmt.Println(msg.Data)

	//3..发送消息
	t := utils.Transfer{
		Conn: conn,
	}

	err = t.WritePkg(msg)
	if err != nil {
		fmt.Println("客户端发送消息失败：", err)
	}
	fmt.Println("客户端发送消息成功：", string(bytes))

	//4..接收数据
	resMsg, err := t.ReadPkg()
	if err != nil {
		fmt.Println("客户端接收数据失败：", err)
		return err
	}
	fmt.Println("客户端接收到消息：", resMsg)

	//5..反序列化拿到状态码
	var m common.RegisterResMessage
	json.Unmarshal([]byte(resMsg.Data), &m)
	if m.Code == 200 {
		fmt.Println("注册成功，请重新登录")

	} else {

		fmt.Println(m.Err)

	}
	return

}
