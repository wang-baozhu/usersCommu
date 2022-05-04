package processor

import (
	"com.kuroro/usersCommu/client/model"
	"com.kuroro/usersCommu/client/utils"
	"com.kuroro/usersCommu/common"
	"encoding/json"
	"fmt"
)

type SmsProcessor struct {
}

func (s *SmsProcessor) SendGroupMes(content string) (err error) {

	var msg common.Message
	msg.Type = common.SmsMesType

	var data common.SmsMessage

	data.User = model.CurrentUser.User
	data.Content = content

	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("序列化失败：", err)
		return
	}
	msg.Data = string(bytes)

	t := utils.Transfer{
		Conn: model.CurrentUser.Conn,
	}
	err = t.WritePkg(msg)
	if err != nil {
		return
	}
	return
}
