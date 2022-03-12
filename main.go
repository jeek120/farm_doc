package main

import (
	"encoding/binary"
	"fmt"
	"github.com/jeek120/farm/server/pb/msg"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
)

func main()  {
	conn, err := net.Dial("tcp", "127.0.0.1:9971")
	if err != nil {
		panic(err)
	}

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	message := &msg.LoginAccPasswdReq{Passwd: "123", Account: "1234"}
	data,err := proto.Marshal(message)
	if err != nil {
		panic(err)
	}

	// len + data
	m := make([]byte, 2+4+len(data))
	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)) + 2)
	binary.BigEndian.PutUint32(m[2:], uint32(message.Cmd()))

	copy(m[6:], data)

	// 发送消息
	conn.Write(m)

	bs := make([]byte, 2)
	_, err = io.ReadFull(conn, bs)
	if err != nil {
		panic(err)
	}
	l := binary.BigEndian.Uint16(bs)
	fmt.Println("长度=", l)
	bs = make([]byte, l)
	_, err = io.ReadFull(conn, bs)
	if err != nil {
		panic(err)
	}
	id := binary.BigEndian.Uint32(bs[:4])
	fmt.Println("返回id=", id)


	message1 := &msg.LoginAccPasswdResp{}
	err = proto.Unmarshal(bs[4:], message1)
	if err != nil {
		panic(err)
	}
	fmt.Println(message1.String())

}