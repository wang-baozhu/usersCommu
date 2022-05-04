package processor

import (
	"com.kuroro/usersCommu/common"
	"encoding/json"
	"fmt"
)

func outShowSms(msg common.Message) {
	var sms common.SmsResMessage
	err := json.Unmarshal([]byte(msg.Data), &sms)
	if err != nil {
		return
	}

	content := fmt.Sprintf("收到来自 %d 的群发消息：%s", sms.UserId, sms.Content)
	fmt.Println(content)

}
