package main

import (
	"chatroom/server/model"
	"chatroom/server/processor"
	"fmt"
	"net"
	"time"
)

// func readPkg(conn net.Conn) (mes message.Message, err error) {
// 	buf := make([]byte, 1024*8)
// 	fmt.Printf("等待读取客户端%v发送的数据···\n", conn.RemoteAddr().String())
// 	//服务器读取用户发来的数据长度
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		//err = errors.New("conn.Read error：")
// 		return
// 	}
// 	//fmt.Println("读到的buf为:", buf[:4])
// 	//根据读到的buf长度，转换成一个uint32，
// 	var pkgLen = binary.BigEndian.Uint32(buf[:4])

// 	//再转换客户端发送的数据本身长度，与buf长度做比对，来读取消息内容
// 	//读取conn中的内容放入buf
// 	n, err := conn.Read(buf[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		//err = errors.New("conn.Read(buf[:pkgLen]) error")
// 		return
// 	}

// 	//将buf[:pkgLen]反序列化为message结构体
// 	//mes要引用类型***
// 	err = json.Unmarshal(buf[:pkgLen], &mes)
// 	//fmt.Println("buf[:pkgLen]", string(buf[:pkgLen]))
// 	fmt.Println("mes", mes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal(buf[:pkgLen],&mes) error :", err)
// 		return
// 	}
// 	return
// }

// func writePkg(conn net.Conn, data []byte) (err error) {
// 	//服务器端发送数据的长度给客户端，确认需要接受的数据长度
// 	var pkgLen uint32 = uint32(len(data)) //定义发送包的大小
// 	var buf [4]byte
// 	//byte=unint8 1个字节 这里定义unint32是4个字节，把字节拆分放入 字节数切片中
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
// 	n, err := conn.Write(buf[:])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write error： ", err)
// 		return
// 	}

// 	//发送消息本身:data
// 	_, err = conn.Write(data)
// 	if err != nil {
// 		fmt.Println("conn.Write(data) error： ", err)
// 		return
// 	}
// 	return
// }

//处理登录消息
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
// 	//1.从mes中取出mes.Data 并反序列化为 message.LoginMes
// 	var loginMes message.LoginMes

// 	err = json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal([]byte(mes.Data),& loginMes) error :", err)
// 		return
// 	}

// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType

// 	var loginResMes message.LoginResMes

// 	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
// 		//合法
// 		loginResMes.Code = 200 //合法账号密码
// 		loginResMes.Error = ""
// 	} else {
// 		//不合法
// 		loginResMes.Code = 500 //500状态码表示用户不存在
// 		loginResMes.Error = "用户不存在,请注册"
// 	}
// 	loginResMes_json, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal(loginResMes) error :", err)
// 		return
// 	}
// 	//登录结果信息序列化放入message.Message的data中
// 	resMes.Data = string(loginResMes_json)
// 	//将message.Message序列化后发送给客户端
// 	resMes_json, err := json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal(resMes) error :", err)
// 		return
// 	}
// 	//发送消息封装装为函数
// 	err = writePkg(conn, resMes_json)
// 	if err != nil {
// 		fmt.Println("writePkg(conn,resMes_json) error :", err)
// 		return
// 	}
// 	return
// }

//ServerProcessMes 函数
//根据mes的类型调用不同的消息分发处理函数
// func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 	case message.LoginMesType:
// 		//处理登录的逻辑
// 		err = serverProcessLogin(conn, mes)
// 	case message.RegisterMesType:
// 		//处理注册逻辑
// 	default:
// 		fmt.Println("消息类型不存在无法处理···")
// 	}
// 	return
// }

//服务器处理和客户端的通讯
func mainProcess(conn net.Conn) {
	//读取客户端发送的信息
	//延时关闭
	defer conn.Close()
	p := &processor.Processor{
		Conn: conn,
	}
	err := p.Process2()
	if err != nil {
		fmt.Println("p.Process2() error :", err)
		return
	}
	//循环读取客户端消息
	// for {
	// 	fmt.Printf("等待读取客户端%v发送的数据···\n", conn.RemoteAddr().String())
	// 	mes, err := readPkg(conn)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			fmt.Println("客户端退出，服务器端也退出··")
	// 			return
	// 		} else {
	// 			fmt.Println("readPkg(conn) error :", err)
	// 			return
	// 		}
	// 	}
	// 	err = serverProcessMes(conn,&mes)
	// 	if err != nil {
	// 		fmt.Println("serverProcessMes(conn,&mes) error :",err)
	// 	}
	// 	//fmt.Println("mes:", mes)
	// }

}

//这里编写一个函数，完成对UserDao的初始化任务
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
func init() {
	//服务器启动时，初始化连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	//初始化全局UserDao
	initUserDao()
}
func main() {

	fmt.Println("服务器[新结构]正在开启8889端口启动监听服务······")
	listener, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("服务器启动监听端口失败··· error：", err)
		return
	}
	defer listener.Close()
	//监听成功，循环等待客服端连接服务器

	for {
		fmt.Println("服务器已在在8889端口启动监听,等待客户端连接······")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("服务端监听器获取通讯信道失败··· error：", err)
		}
		//连接成功在主线程中开辟一个携程，实现服务端和客户端通讯
		go mainProcess(conn)
	}
}
