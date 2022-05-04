package processor

import (
	"com.kuroro/usersCommu/common"
	"fmt"
)

var UserOnlineMap = make(map[int]common.User, 10)

func updateUserStatus(notifyMsg common.UserStatusNotifyMes) {

	user, ok := UserOnlineMap[notifyMsg.UserId]
	if !ok {
		UserOnlineMap[notifyMsg.UserId] = common.User{
			UserId:     notifyMsg.UserId,
			UserStatus: notifyMsg.UserStatus,
		}
		showOnlineUsers()
		return
	}

	user.UserStatus = notifyMsg.UserStatus
	showOnlineUsers()

}

func showOnlineUsers() {
	for id, _ := range UserOnlineMap {
		fmt.Println("在线用户ID:", id)
	}
}
