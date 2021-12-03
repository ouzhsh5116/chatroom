package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

//MD5方式1
func getMD5_1(str []byte) string {
	// 1. 直接用原数据计算数据的md5
	res := md5.Sum(str)
	fmt.Printf("%x\n", res)
	//字节切片转字符串
	res2 := fmt.Sprintf("%x", res)
	fmt.Println("res2:", res2)
	// --- 这是另外一种格式化切片的方式 数组转切片
	res3 := hex.EncodeToString(res[:])
	fmt.Println("res3: ", res3)
	return res3
}

//MD5方式2
func getMD5_2(str []byte) string {
	// 1. 创建一个使用MD5校验的Hash对象`
	myHash := md5.New()
	// 2. 通过io操作将数据写入hash对象中
	io.WriteString(myHash, "hello")
	//io.WriteString(myHash, ", world")
	myHash.Write([]byte(", world"))
	// 3. 计算结果 若在hash器里面写了数据，sum里面再写参数进去，sum中参数运算的hash会追加到hash器结果的前面
	result := myHash.Sum(nil)
	fmt.Println(result)
	// 4. 将结果转换为16进制格式字符串
	res := fmt.Sprintf("%x", result)
	fmt.Println(res)
	// --- 这是另外一种格式化切片的方式
	res = hex.EncodeToString(result)
	fmt.Println(res)
	return res
}

func main() {
	getMD5_1([]byte("123"))
	res := getMD5_2([]byte("1234"))
	fmt.Println(res)
}
