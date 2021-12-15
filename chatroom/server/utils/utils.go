package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//将方法关联到传输结构体中
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //传输时使用的缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 1024*8)
	fmt.Printf("等待读取客户端%v发送的数据···\n", this.Conn.RemoteAddr().String())
	//服务器读取用户发来的数据长度
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//err = errors.New("conn.Read error：")
		return
	}
	//fmt.Println("读到的buf为:", buf[:4])
	//根据读到的buf长度，转换成一个uint32，
	var pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	//再转换客户端发送的数据本身长度，与buf长度做比对，来读取消息内容
	//读取conn中的内容放入this.Buf
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("this.Conn.Read(this.Buf[:pkgLen]) error")
		return
	}

	//将buf[:pkgLen]反序列化为message结构体
	//mes要引用类型***
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	//fmt.Println("buf[:pkgLen]", string(buf[:pkgLen]))
	fmt.Println("mes", mes)
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen],&mes) error :", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//服务器端发送数据的长度给客户端，确认需要接受的数据长度
	var pkgLen uint32 = uint32(len(data)) //定义发送包的大小
	//var buf [4]byte
	//byte=unint8 1个字节 这里定义unint32是4个字节，把字节拆分放入 字节数切片中
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write error： ", err)
		return
	}

	//发送消息本身:data
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("this.Conn.Write(data) error： ", err)
		return
	}
	return
}
