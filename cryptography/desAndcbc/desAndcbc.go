package main

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func DesEncrypt_CBC(src, key []byte) []byte {
	block, error := des.NewCipher(key)
	if error != nil {
		panic(error)
	}
	iv := []byte{'1', '2', '3', '4', '5', '6', '7', '8'} //初始化向量8byte
	//填充字符
	src = PKCS5Padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, iv)
	blockmode.CryptBlocks(src, src)
	return src
}

//填充字符
func PKCS5Padding(src []byte, blockSize int) []byte {
	paddingLen := blockSize - (len(src) % blockSize)
	paddingText := make([]byte, paddingLen)
	for i := range paddingText {
		paddingText[i] = byte(paddingLen)
	}
	src = append(src, paddingText...)
	return src
}

func DesDecrypt_CBC(cipherText, key []byte) []byte {
	block, error := des.NewCipher(key)
	if error != nil {
		panic(error)
	}
	iv := []byte{'1', '2', '3', '4', '5', '6', '7', '8'} //初始化向量8byte
	blockmode := cipher.NewCBCDecrypter(block, iv)
	blockmode.CryptBlocks(cipherText, cipherText)
	painText := cipherText
	//去除填充
	painText = PKCS5UnPadding(painText)
	return painText
}

//删除字符
func PKCS5UnPadding(origData []byte) []byte {
	endChar := origData[len(origData)-1]
	endInt := int(endChar)
	return origData[:len(origData)-endInt]
}
func main() {
	src := []byte("这是要加密的明文！")
	key := []byte("11111111") //密钥8byte
	cipherText := DesEncrypt_CBC(src, key)
	fmt.Printf("加密后的密文是：%x\n", cipherText)
	painText := DesDecrypt_CBC(cipherText, key)
	fmt.Printf("解密后的明文是：%s", painText)
}
