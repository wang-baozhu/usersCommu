package utils

import (
	"com.kuroro/usersCommu/common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn   net.Conn
	Buffer [8096]byte
}

func (t *Transfer) WritePkg(resMsg common.Message) (err error) {

	bytes, err := json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json序列化失败:", err)
		return
	}

	//先发送数据长度信息
	var lenMsgBytes [4]byte
	binary.BigEndian.PutUint32(lenMsgBytes[:], uint32(len(bytes)))

	_, err = t.Conn.Write(lenMsgBytes[:])
	if err != nil {
		fmt.Println("发送数据长度信息失败:", err)
		return
	}

	_, err = t.Conn.Write(bytes)
	if err != nil {
		fmt.Println("发送消息体数据失败：", err)
		return
	}
	return

}

//服务端接收消息
func (t *Transfer) ReadPkg() (msg common.Message, err error) {

	buffer := make([]byte, 8096)
	//读取消息长度
	_, err = t.Conn.Read(buffer[:4])
	if err != nil {
		return
	}
	//转换长度
	len := binary.BigEndian.Uint32(buffer[:4])
	//读取消息体
	_, err = t.Conn.Read(buffer[:len])
	if err != nil {
		return
	}
	err = json.Unmarshal(buffer[:len], &msg)
	if err != nil {
		fmt.Println("json反序列化失败:", err)
		return
	}

	return

}

func test222222222() {
	fmt.Println("xixixixi")
	fmt.Println("zzzzzzzz")
}
