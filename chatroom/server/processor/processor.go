package processor

import (
	"chatroom/common/message"
	"chatroom/server/process"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//ServerProcessMes 函数
//根据mes的类型调用不同的消息分发处理函数
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {
	fmt.Println("群发消息:",mes)
	switch mes.Type {
	case message.LoginMesType:
		//处理登录的逻辑
		//userProcess实例调用登录处理函数
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册逻辑
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegist(mes)
	case message.SmsMesType:
		//处理客户端的群发消息
		sp := &process.SmsProcess{
		}
		sp.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在无法处理···")
	}
	return
}
func (this *Processor) Process2 () (err error) {
	for {
		//创建transfer完成读包
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		fmt.Printf("等待读取客户端%v发送的数据···\n", this.Conn.RemoteAddr().String())
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出··")
				return err
			} else {
				fmt.Println("readPkg(conn) error :", err)
				return err
			}
		}
		err = this.ServerProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes(conn,&mes) error :",err)
			return err
		}
	}
}
