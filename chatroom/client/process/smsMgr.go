package process

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

func OutputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)

	if err!= nil {
		fmt.Println("客户端输出群发消息失败 err :",err)
		return
	}

	info := fmt.Sprintf("用户id:\t%d 对大家说:\t%s",smsMes.Id,smsMes.Content)
	fmt.Println(info)

}