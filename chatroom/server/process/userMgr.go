package process

import "fmt"

//userMgr只有一个可以定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//初始化userMgr
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineUsers的添加/修改
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//返回所有在线用户
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {

	return this.onlineUsers
}

//根据id获取对应userprocess
func (this *UserMgr) GetOnlineUserById(userId int)(up *UserProcess,err error) {
	up , ok := this.onlineUsers[userId]
	//ok为假当前用户不在线
	if !ok {
		err = fmt.Errorf("用户%v不在线···",userId)
		return
	}
	return
}
