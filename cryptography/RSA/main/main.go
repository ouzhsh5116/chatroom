package main

import (
	_ "cryptography/RSA/RSAGen"
	"cryptography/RSA/RSA_decrypt_encrypt"
	"fmt"
)
func main () {
	//生成公私钥文件
	//rsagen.RsaGenKey(1024)

	cryptText := rsadecryptencrypt.RSAEncrypt([]byte("wocannns我我无法我"),[]byte("../main/public.pem"))
	paintText := rsadecryptencrypt.RSADecrypt(cryptText,[]byte("../main/private.pem"))
	fmt.Println("密文：",string(cryptText))
	fmt.Println("明文:",string(paintText))
}