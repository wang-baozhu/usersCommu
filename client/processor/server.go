package processor

import (
	"com.kuroro/usersCommu/client/utils"
	"com.kuroro/usersCommu/common"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
)

func ShowMenu() {

	fmt.Println("----------登录成功，欢迎使用----------")
	fmt.Println("---------1-显示在线用户列表--------")
	fmt.Println("---------2-发送消息--------")
	fmt.Println("---------3-显示信息列表--------")
	fmt.Println("---------4-退出系统--------")
	fmt.Println("请输入1-4选择功能")

	sp := SmsProcessor{}
	key := ""
	fmt.Scanln(&key)

	switch key {
	case "1":
		showOnlineUsers()
	case "2":
		fmt.Println("请输入要发送的消息，并enter发送：")
		content := ""
		fmt.Scanln(&content)
		//将SmsProcess定义在外面这样就不需要循环一次就创建一次
		sp.SendGroupMes(content)
	case "3":
		fmt.Println("显示信息列表")
	case "4":
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入选项不正确")
	}

}

func MonitorServer(conn net.Conn) {
	t := utils.Transfer{
		Conn: conn,
	}

	for {
		fmt.Println("正在监听。。。")
		msg, err := t.ReadPkg()
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("发生了其他错误：", err)
			return
		}

		switch msg.Type {
		case common.UserStatusNotifyMesType:
			fmt.Println("------------------有情况------------")
			var notifyMsg common.UserStatusNotifyMes
			json.Unmarshal([]byte(msg.Data), &notifyMsg)

			updateUserStatus(notifyMsg)
		case common.SmsResMesType:
			fmt.Println("---------------收到群发消息--------- ")
			outShowSms(msg)

		}

	}

}
