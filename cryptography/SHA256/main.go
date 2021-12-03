package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"io"
	"os"
	"fmt"
	"encoding/hex"
)
func getSHA1(src string) string {
	// 1. 打开文件
	file ,err := os.Open(src)
	if err!=nil {
		return "文件打开失败"
	}
	// 2. 创建基于sha1算法的Hash对象
	myHash := sha1.New()
	num ,err :=io.Copy(myHash,file)
	if err != nil {
		return "拷贝文件失败"
	}
	fmt.Println("文件大小: ", num)
	// 4. 计算文件的哈希值
	tmp1 := myHash.Sum(nil)
	// 5. 数据格式转换
	result := hex.EncodeToString(tmp1)
	fmt.Println("sha1: ", result)
	return result
}
func getSHA256(src string) string {
	// 1. 打开文件
	file ,err := os.Open(src)
	if err!=nil {
		return "文件打开失败"
	}
	// 2. 创建基于sha256算法的Hash对象
	myHash := sha256.New()
	num ,err :=io.Copy(myHash,file)
	if err != nil {
		return "拷贝文件失败"
	}
	fmt.Println("文件大小: ", num)
	// 4. 计算文件的哈希值
	tmp1 := myHash.Sum(nil)
	// 5. 数据格式转换
	result := hex.EncodeToString(tmp1)
	fmt.Println("sha256: ", result)
	return result
}
func main() {
	getSHA1("../RSA/main/public.pem")
	getSHA256("../RSA/main/public.pem")
}