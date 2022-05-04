package main

import (
	"com.kuroro/usersCommu/server/model"
	"com.kuroro/usersCommu/server/processor"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"io"
	"net"
	"time"
)

func initUserDao(pool *redis.Pool) {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	//初始化连接池，创建了一个全局的pool
	pool := processor.InitPool("192.168.56.10:6379", 8, 0, 100*time.Millisecond)
	initUserDao(pool)

	fmt.Println("服务器在8000端口监听")
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("监听失败:", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接失败：", err)
			return
		}

		go process(conn)

	}
}

func process(conn net.Conn) {

	defer conn.Close()
	p := processor.Processor{
		Conn: conn,
	}
	err := p.ReadAndWrite()
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("协程处理消息失败:", err)
		return
	}

}
