package processor

import (
	"com.kuroro/usersCommu/common"
	"com.kuroro/usersCommu/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (p *Processor) ReadAndWrite() (err error) {

	for {
		t := utils.Transfer{
			Conn: p.Conn,
		}
		msg, err := t.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端退出")
				return err
			}
			fmt.Println("发生了其他错误：", err)
			return err
		}
		fmt.Println("服务接收到消息：", msg)
		err = p.serverProcessMsg(msg)
		if err != nil {
			return err
		}

	}

}

//接收完消息后处理消息
func (p *Processor) serverProcessMsg(msg common.Message) (err error) {

	t := msg.Type
	switch t {
	case common.LoginMesType:
		up := UserProcessor{
			Conn: p.Conn,
		}
		err = up.ProcessLogin(msg)
	case common.RegisterMesType:
		up := UserProcessor{
			Conn: p.Conn,
		}
		err = up.ProcessRegister(msg)

	case common.SmsMesType:
		fmt.Println("准备群发消息：", msg.Data)
		sp := SmsProcessor{}
		err = sp.SendMsgToAll(msg)

	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return

}
