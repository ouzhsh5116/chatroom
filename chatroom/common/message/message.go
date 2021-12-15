package message

/**
定义各种公共使用的消息类型
*/
const (
	LoginMesType             = "LoginMes"
	LoginResMesType          = "LoginResMes"
	RegisterMesType          = "RegisterMes"
	RegisterResMesType       = "RegisterResMes"
	NotifyUsersStatusMesType = "NotifyUsersStatusMes"
	SmsMesType               = "SmsMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息具体数据
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code    int    `json:"code"` //返回登录结果的状态码	500：用户未注册 200：登录成功
	UsersId []int  //增加在线用户切片
	Error   string `json:"error"` //错误消息，没错误返回nil
}

type RegisterMes struct {
	User User `json:"user"`
}
type RegisterResMes struct {
	Code  int    `json:"code"`  //返回注册结果的状态码	400：用户已被注册 200：注册 成功
	Error string `json:"error"` //错误消息，没错误返回nil
}

//服务器推送用户上线通知
type NotifyUsersStatusMes struct {
	UserId     int `json:"userId"`     //用户id
	UserStatus int `json:"userStatus"` //用户状态
}

//发送的SmsMes
type SmsMes struct {
	Content string `json:"content"` //消息内容
	User           //匿名结构体 继承User

}

//定义用户状态常量
const (
	UserOnline = iota
	UserOffline
	UserBusy
)
