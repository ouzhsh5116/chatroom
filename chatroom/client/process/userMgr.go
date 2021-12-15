package process

import (
	"chatroom/client/model"
	"chatroom/common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 100)
var curUser model.CurUser //用户登录成功后对curUser初始化
//客户端显示在线用户列表
func outputOnlineUsers() {
	fmt.Println("当前在线用户列表:")
	for _, v := range onlineUsers {
		fmt.Printf("用户id:%v,%v,状态:%v\n", v.Id, v.Name, v.Status)
	}
}

func updateUsersStatus(notifyUsersStatusMes *message.NotifyUsersStatusMes) {

	user, ok := onlineUsers[notifyUsersStatusMes.UserId]
	if !ok {
		user = &message.User{
			Id: notifyUsersStatusMes.UserId,
		}
	}

	user.Status = notifyUsersStatusMes.UserStatus
	onlineUsers[notifyUsersStatusMes.UserId] = user

	outputOnlineUsers()
}
