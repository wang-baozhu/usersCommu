package processor

import (
	"fmt"
)

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcessor
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcessor, 512),
	}
}

//添加与修改
func (u *UserMgr) AddOrModifyUp(up *UserProcessor) {
	u.onlineUsers[up.UserId] = up
}

//删除
func (u *UserMgr) DelUp(userId int) {

	delete(u.onlineUsers, userId)
}

//查询某个在线用户
func (u *UserMgr) GetUpById(userId int) (up *UserProcessor, err error) {

	up, ok := u.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户 %d 不在线", userId)
		return
	}
	return

}

//查询所以在线用户
func (u *UserMgr) GetAllUp() map[int]*UserProcessor {
	return u.onlineUsers
}
