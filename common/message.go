package common

const (
	LoginMesType    = "LoginMessage"
	LoginResMesType = "LoginResMessage"

	RegisterMesType    = "RegisterMessage"
	RegisterResMesType = "RegisterResMessage"

	UserStatusNotifyMesType = "UserStatusNotifyMessage"

	SmsMesType = "SmsMessage"

	SmsResMesType = "SmsResMessage"
)

const (
	UserOffline = -1
	UserOnline  = 1
	UserBusy    = 0
)

type Message struct {
	//消息类型
	Type string `json:"type"`
	//消息内容
	Data string `json:"data"`
}

type LoginMessage struct {
	Id int `json:"id"`

	UserName string `json:"userName"`

	Password string `json:"password"`
}

type LoginResMessage struct {
	//500表示登录失败，200表示登录成功
	Code int `json:"code"`

	Err string `json:"err"`
	//存储登录用户id
	UsersId []int `json:"usersId"`
}

type RegisterMessage struct {
	User User `json:"user"`
}

type RegisterResMessage struct {
	//400表示用户名已存在，注册失败；200表示注册成功
	Code int `json:"code"`

	Err string `json:"err"`
}

type UserStatusNotifyMes struct {
	UserId     int `json:"userId""`
	UserStatus int `json:"userStatus"`
}

type SmsMessage struct {
	User    User   `json:"user"`
	Content string `json:"content"`
}

type SmsResMessage struct {
	UserId  int    `json:"userId"`
	Content string `json:"content"`
}
