package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
	//Conn net.Conn
}

//转发客户端的群消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//遍历onlineUsers map[int]*UserProcess转发消息
	//取出mes里面的data 做id判断是不是发送者本身
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("转发消息反序列化失败·· error :", err)
		return
	}

	mes_json, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("转发消息序列化失败·· error :", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		//过滤发送者自己
		if smsMes.Id == id {
			continue
		}

		this.SendMesToEachOnlineUser(mes_json, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(info []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(info)
	if err != nil {
		fmt.Println("转发消息失败·· error :", err)
		return
	}
}
