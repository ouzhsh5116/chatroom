package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

type smsProcess struct {
}

//群发消息
func (this *smsProcess) SendGroupMes(content string) (err error) {
	
	var mes message.Message
	mes.Type=message.SmsMesType

	//创建SmsMes实例
	var smsMes message.SmsMes 
	smsMes.Content = content
	smsMes.Id = curUser.Id
	smsMes.Status = curUser.Status
	//发送mes给服务器
	smsMes_json,err :=json.Marshal(smsMes)
	if err!= nil {
		fmt.Println("smsMes 序列化失败 error :",err)
		return
	}
	mes.Data=string(smsMes_json)

	mes_json ,err :=json.Marshal(mes)
	if err!= nil {
		fmt.Println("mes 序列化失败 error :",err)
		return
	}

	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}

	err = tf.WritePkg(mes_json)
	if err!= nil {
		fmt.Println("客户端发送群聊消息失败 error :",err)
		return
	}
	return
}


