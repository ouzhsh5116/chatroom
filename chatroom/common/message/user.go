package message

//一个用户信息的结构体
type User struct {
	Id   int    `json:"userId"`
	Pwd  string `json:"userPwd"`
	Name string `json:"userName"`
	//后续可以加入的字段
	// Nick      string `json:"nick"`
	// Sex       string `json:"sex"`
	// Header    string `json:"header"`
	// LastLogin string `json:"last_login"`
	Status int `json:"status"`
}
