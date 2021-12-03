package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func AesEncrypt_CTR(src, key []byte) []byte {
	block, error := aes.NewCipher(key)
	if error != nil {
		panic(error)
	}
	iv := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '1', '2', '3', '4', '5', '6', '7', '8'} //初始化向量8byte
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(src, src)
	return src
}

func AesDecrypt_CTR(src, key []byte) []byte {
	block, error := aes.NewCipher(key)
	if error != nil {
		panic(error)
	}
	iv := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '1', '2', '3', '4', '5', '6', '7', '8'} //初始化向量8byte
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(src, src)
	return src
}

func main() {
	src := []byte("这是要加密的明文！")
	key := []byte("1111111111111111") //密钥8byte
	cipherText := AesEncrypt_CTR(src, key)
	fmt.Printf("加密后的密文是：%x\n", cipherText)
	fmt.Printf("加密后的密文是：%v\n", cipherText)
	painText := AesDecrypt_CTR(cipherText, key)
	fmt.Printf("解密后的明文是：%s", painText)
}
