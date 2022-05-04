package main

import (
	"com.kuroro/usersCommu/client/processor"
	"fmt"
)

type userClient struct {
	key         string
	loopOut     bool
	userService processor.UserProcessor
}

func (u *userClient) mainMenu() {

	for {
		fmt.Println("--------欢迎使用多人聊天系统--------")
		fmt.Println("\t\t\t 1-登录聊天室")
		fmt.Println("\t\t\t 2-注册用户")
		fmt.Println("\t\t\t 3-退出系统")
		fmt.Println("请选择功能(1-3):")
		fmt.Scanln(&u.key)

		switch u.key {
		case "1":
			fmt.Println("正在登录。。。")
			fmt.Println("请输入用户编号")
			id := 0
			fmt.Scanln(&id)

			fmt.Println("请输入用户名")
			username := ""
			fmt.Scanln(&username)

			fmt.Println("请输入密码")
			password := ""
			fmt.Scanln(&password)

			up := processor.UserProcessor{}
			up.Login(id, username, password)

		case "2":
			fmt.Println("正在注册。。。")
			fmt.Println("请输入用户编号")
			id := 0
			fmt.Scanln(&id)

			fmt.Println("请输入用户名")
			username := ""
			fmt.Scanln(&username)

			fmt.Println("请输入密码")
			password := ""
			fmt.Scanln(&password)

			up := processor.UserProcessor{}
			up.Register(id, username, password)
		case "3":
			fmt.Println("退出系统")
			u.loopOut = true
		default:
			fmt.Println("输入不正确，请重新输入")
		}

		if u.loopOut {
			break
		}
	}
	fmt.Println("已退出系统")

}

func main() {

	u := userClient{
		key:         "",
		loopOut:     false,
		userService: processor.UserProcessor{},
	}
	u.mainMenu()

}
