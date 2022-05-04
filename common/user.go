package common

type User struct {
	UserId     int    `json:"userId"`
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	UserStatus int    `json:"userStatus"`
}
