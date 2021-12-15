package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
}

//error作返回值信息，描述能力更强，error=nil则没问题
//关联用户登录方法
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	fmt.Printf("userId = %d , passwd = %s\n", userId, userPwd)

	//1.获取客户端登录输入的账户名、密码，连接到服务器的端口
	conn, err1 := net.Dial("tcp", "localhost:8889")
	if err1 != nil {
		fmt.Println("客户端连接服务器失败··· error:", err1)
		return err1
	}
	defer conn.Close()
	//把账户名密码写入登录消息结构体，选择登录消息类型，序列化为message的data
	//计算字节长度发送给服务器

	//2.连接服务器成功,发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建loginmes结构体 写入账号密码 ，序列化后写入message的data
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//4将loginMes转json
	loginMes_json, err2 := json.Marshal(loginMes)
	if err2 != nil {
		fmt.Println(" loginMes json.Marshal error :", err2)
		return err2
	}
	//5将loginmes序列化后写入mes.Data
	mes.Data = string(loginMes_json)

	//6.将mes序列化
	mes_json, err3 := json.Marshal(mes)
	if err3 != nil {
		fmt.Println("mes json.Marshal error :", err3)
		return err3
	}
	//7.计算序列化后的mes中Data的长度，发送给服务器
	//使用binary包import "encoding/binary"binary包实现了简单的数字与字节序列的转换以及变长值的编解码。
	/*
				type ByteOrder interface {
		    Uint16([]byte) uint16
		    Uint32([]byte) uint32
		    Uint64([]byte) uint64
		    PutUint16([]byte, uint16)
		    PutUint32([]byte, uint32) **
		    PutUint64([]byte, uint64)
		    String() string
		}
	*/
	var pkgLen uint32 = uint32(len(mes_json)) //定义发送包的大小
	var buf [4]byte
	//byte=unint8 1个字节 这里定义unint32是4个字节，把字节拆分放入 字节数切片中
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err4 := conn.Write(buf[:])
	if n != 4 || err4 != nil {
		fmt.Println("conn.Write error: ", err4)
		return err4
	}

	fmt.Println("客户端发送消息的长度结束···", len(mes_json))
	fmt.Println("mes_json:", string(mes_json))

	//发送消息本身:mes_json
	_, err5 := conn.Write(mes_json)
	if err5 != nil {
		fmt.Println("conn.Write(mes_json) error: ", err5)
		return err5
	}

	//处理服务器端返回的消息
	//调用工具类读包
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()

	if err != nil {
		fmt.Println("readPkg(conn) error :", err)
	}

	//将mes的data反序列化为loginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data),&loginResMes) error :", err)
		return
	}
	//200登录成功
	if loginResMes.Code == 200 {
		//初始化curUser
		curUser.Conn = conn
		curUser.Id=userId
		curUser.Status = message.UserOnline



		//fmt.Println("用户登录成功")
		//1.循环显示登录成功的菜单
		fmt.Println("用户在线列表:")
		for _,v :=range loginResMes.UsersId {
			if v==userId {
				continue
			}
			fmt.Println("当前在线用户,Id:",v)
			//完成客户端onlineUsers的初始化
			//遍历在线用户列表放入onlineUsers
			user := &message.User{
				Id : v,
				Status: message.UserOnline,
			}
			onlineUsers[v]=user
		}
		//一旦登录成功，就需要实时的读取服务端发送的消息并处理
		//,因此专门其动一个goroutine,保持和服务器通讯
		go ProcessServerMessage(conn)
		for {
			ShowMenu()
		}
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
		return
	} else if loginResMes.Code == 403 {
		fmt.Println(loginResMes.Error)
		return
	}
	return

}
func (this *UserProcess) Regist(userId int, userPwd string, userName string) (err error) {
	//1.获取客户端登录输入的账户名、密码，连接到服务器的端口
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("客户端连接服务器失败··· error:", err)
		return err
	}
	defer conn.Close()

	//2.连接服务器成功,发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.赋值
	var registerMes = message.RegisterMes{}
	registerMes.User.Id = userId
	registerMes.User.Name = userName
	registerMes.User.Pwd = userPwd

	//4将registerMes转json
	registerMes_json, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println(" registerMes json.Marshal error :", err)
		return err
	}
	//5将registerMes序列化后写入mes.Data
	mes.Data = string(registerMes_json)

	//6.将mes序列化
	mes_json, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json.Marshal error :", err)
		return err
	}

	//处理服务器端返回的消息
	//调用工具类读包
	tf := &utils.Transfer{
		Conn: conn,
	}
	//7.发送消息长度
	//8.发送消息本身:mes_json
	err = tf.WritePkg(mes_json)
	if err != nil {
		fmt.Println("注册发送消息错误:", err)
	}

	//读取服务器返回的注册消息
	mes, err = tf.ReadPkg() //mes是服务器对客户端注册时的回应消息registerResMes

	if err != nil {
		fmt.Println("readPkg(conn) error :", err)
	}

	//将mes的data反序列化为registerResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data),&registerResMes) error :", err)
		return
	}
	//200登录成功
	if registerResMes.Code == 200 {
		//fmt.Println("用户注册成功")
		//返回登录菜单

	} else if registerResMes.Code == 400 {
		fmt.Println("用户注册失败,已经存在···", registerResMes.Error)
		//return
	} else {
		fmt.Println("用户注册失败(其他原因)···", registerResMes.Error)
	}
	return
}
