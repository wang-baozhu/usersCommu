package processor

import (
	"com.kuroro/usersCommu/client/utils"
	"com.kuroro/usersCommu/common"
	"encoding/json"
	"fmt"
)

type SmsProcessor struct {
}

func (s *SmsProcessor) SendMsgToAll(msg common.Message) (err error) {

	var sms common.SmsMessage
	err = json.Unmarshal([]byte(msg.Data), &sms)
	if err != nil {
		fmt.Println("反序列化失败：", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//排除发给自己
		if id == sms.User.UserId {
			continue
		}

		err := s.sendToEach(up, sms.User.UserId, sms.Content)
		if err != nil {
			return err
		}

	}
	return

}

func (s *SmsProcessor) sendToEach(up *UserProcessor, id int, content string) (err error) {

	var msg common.Message
	msg.Type = common.SmsResMesType

	var sms common.SmsResMessage
	sms.UserId = id
	sms.Content = content

	bytes, err := json.Marshal(sms)

	if err != nil {
		return
	}
	msg.Data = string(bytes)

	t := utils.Transfer{
		Conn: up.Conn,
	}

	err = t.WritePkg(msg)
	if err != nil {
		return
	}
	return

}
