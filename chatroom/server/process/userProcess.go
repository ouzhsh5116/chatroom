package process

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	//增加Conn关联的用户
	UserId int
}

//通知所有在线用户
//userId上线后，通知其他人
func (this *UserProcess) NotifyOthersOnlineUsers(userId int) {
	//遍历UserMgr的OnlineUsers,一个一个发送NotifyUsersStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == userId {
			continue
		}
		//开始通知其他人
		//用其他用户的up发送消息给他们
		up.NotifyMeOnline(userId)
	}
}

//通知别人，自己(userId)上线了
func (this *UserProcess) NotifyMeOnline(userId int) {
	//组织通知上线消息
	var mes message.Message
	mes.Type = message.NotifyUsersStatusMesType

	var notifyUsersStatusMes message.NotifyUsersStatusMes
	notifyUsersStatusMes.UserId = userId
	notifyUsersStatusMes.UserStatus = message.UserOnline

	//将notifyUsersStatusMes序列化放入mes的data中
	notifyUsersStatusMes_json, err := json.Marshal(notifyUsersStatusMes)
	if err != nil {
		fmt.Println("通知用户状态信息序列化失败 error:", err)
		return
	}
	mes.Data = string(notifyUsersStatusMes_json)

	mes_json, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("通知用户状态信息序列化失败 error:", err)
		return
	}

	//向up发送mes_json
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(mes_json)
	if err != nil {
		fmt.Println("通知别人上线出错 error :", err)
	}

}

//处理登录消息
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1.从mes中取出mes.Data 并反序列化为 message.LoginMes
	var loginMes message.LoginMes

	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data),& loginMes) error :", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	//根据用户的登录信息和redis数据库比对完成验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		//不合法
		if err == model.ErrUserNotExist {
			loginResMes.Code = 500 //500状态码表示用户不存在
			loginResMes.Error = err.Error()
		} else if err == model.ErrInvalidPasswd {
			loginResMes.Code = 403 //403状态码表示密码不正确
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "未知错误信息"
		}
	} else {
		//合法
		loginResMes.Code = 200 //合法账号密码
		loginResMes.Error = ""
		fmt.Println("登录成功 user:", user)

		//将登录成功的userid赋值给Userprocess
		this.UserId = loginMes.UserId
		//用户登录,将用户放入UserMgr,作为在线用户列表
		userMgr.AddOnlineUser(this)
		//某个用户上线，通知其他在线用户
		this.NotifyOthersOnlineUsers(loginMes.UserId)
		//将userId放入loginResmes的usersid中
		//遍历userMgr。onlineUsers
		for id := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}

	}

	loginResMes_json, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) error :", err)
		return
	}
	//登录结果信息序列化放入message.Message的data中

	resMes.Data = string(loginResMes_json)
	//将message.Message序列化后发送给客户端
	resMes_json, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) error :", err)
		return
	}
	//发送消息封装装为函数
	//MVC结构 调用utils里的transfer来读写包
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(resMes_json)
	if err != nil {
		fmt.Println("writePkg(conn,resMes_json) error :", err)
		return
	}
	return
}

//处理注册消息
func (this *UserProcess) ServerProcessRegist(mes *message.Message) (err error) {
	var registMes message.RegisterMes

	err = json.Unmarshal([]byte(mes.Data), &registMes)
	if err != nil {
		fmt.Println("注册消息反序列化失败 error :", err)
		return
	}

	//声明一个注册返回的结果信息
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	//声明注册结果信息
	var registerResMes message.RegisterResMes

	//根据用户的登录信息和redis数据库比对完成验证
	err = model.MyUserDao.Regist((&registMes.User))

	if err != nil {
		if err == model.ErrUserExist {
			registerResMes.Code = 400
			registerResMes.Error = model.ErrUserExist.Error()
		} else {
			registerResMes.Code = 505
			registerResMes.Error = "未知错误*"
		}
	} else {
		registerResMes.Code = 200
	}

	registerResMes_json, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println(" json.Marshal(registerResMes) error :", err)
		return
	}
	//登录结果信息序列化放入message.Message的data中

	resMes.Data = string(registerResMes_json)
	//将message.Message序列化后发送给客户端
	resMes_json, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) error :", err)
		return
	}
	//发送消息封装装为函数
	//MVC结构 调用utils里的transfer来读写包
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(resMes_json)
	if err != nil {
		fmt.Println("writePkg(conn,resMes_json) error :", err)
		return
	}

	return
}
