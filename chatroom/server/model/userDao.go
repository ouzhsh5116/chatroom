package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

//服务器启动时就初始化一个userDao
var (
	MyUserDao *UserDao
)

//定义UserDao完成对user的操作
type UserDao struct {
	pool *redis.Pool
}

//使用工厂实例创建UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = new(UserDao)
	userDao.pool = pool
	return
}

//根据用户id返回user实例或error
func (this *UserDao) GetUserById(conn redis.Conn, id int) (user *User, err error) {
	redisRes, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		//没查到users中对应的用户id信息
		if err == redis.ErrNil {
			err = ErrUserNotExist
		}
		return

	}

	user = new(User)
	//将redisRes反序列化
	err = json.Unmarshal([]byte(redisRes), user)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(redisRes),user) error :", err)
		return
	}
	return
}

//登录校验
//1.login完成用户校验
//2.如果用户的id和pwd都正确，则返回一个user实例
//3.不正确，返回错误信息

func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	//连接池中取一个redis连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.GetUserById(conn, userId)
	if err != nil {
		return
	}

	//用户存在，校验密码
	if user.Pwd != userPwd {
		err = ErrInvalidPasswd
		return
	}

	return
}


func (this *UserDao) Regist(user *message.User) (err error) {

	//连接池中取一个redis连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.GetUserById(conn, user.Id)
	if err == nil {
		err = ErrUserExist
		//没有返回错误则表示用户存在，注册失败
		return
	}

	//这时用户不存在，可以完成注册，入redis库
	//序列化user
	user_json ,err :=json.Marshal(user)
	if err!= nil {
		fmt.Println("user序列化失败··")
		return
	}
	_,err =conn.Do("hset","users",user.Id,string(user_json))
	if err!= nil {
		fmt.Println("user入库失败··",err)
		return
	}
	return
}
