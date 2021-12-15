package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//保持和服务器通讯,服务器发送数据，客户端处理
func ProcessServerMessage(conn net.Conn) {
	//创建一个tf 来不停的读取服务器的信息...
	tf := &utils.Transfer{Conn: conn}
	for {
		fmt.Printf("不停的读取从服务器%s发来的信息，并准备处理...\n", conn.RemoteAddr().String())
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() error :", err)
			return
		}
		//读取到服务器的消息后
		switch mes.Type {
		case message.NotifyUsersStatusMesType:
			//fmt.Println("服务器发来的消息:", mes)
			//读取用户上线列表
			var notifyUsersStatusMes message.NotifyUsersStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUsersStatusMes)
			//将用户信息保存到客户端维护的onlineusers的map[int]user中

			if err != nil {
				fmt.Println("客户端接受在线用户消息失败··· error :", err)
			}
			updateUsersStatus(&notifyUsersStatusMes)
		case message.SmsMesType:
			//有人群发消息
			OutputGroupMes(&mes)

		default:
			fmt.Println("服务器返回的消息类型无法处理···")
		}

	}
}

//客户端登录成功后显示二级菜单

func ShowMenu() {
	fmt.Println("1. 显示在线用户列表")
	fmt.Println("2. 发送信息")
	fmt.Println("3. 信息列表")
	fmt.Println("4. 退出系统")
	fmt.Println("请选择1-4:")
	var key int
	var content string

	smsProcess := &smsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//显示用户在线列表
		outputOnlineUsers()
	case 2:
		fmt.Println("请输入你要群发的消息:")
		fmt.Scanf("%s\n\n",&content)

		smsProcess.SendGroupMes(content)
		//enterTalk(conn)
	case 3:
		fmt.Println("信息列表")
		//listUnReadMsg()
		return
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不对，请重新输入")
	}
}
