package Crypto

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

// DesEncrypt 使用DES CFB模式进行加密, 需要8位的key和iv
func DesEncrypt(plainText, key, iv []byte) ([]byte, error) {
	// 创建一个DES密码分组
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 对原始数据进行填充
	plainText = pkcs7Padding(plainText, block.BlockSize())

	// 创建一个CFB加密模式
	cfb := cipher.NewCFBEncrypter(block, iv)

	// 加密数据
	encrypted := make([]byte, len(plainText))
	cfb.XORKeyStream(encrypted, plainText)

	return encrypted, nil
}

// DesDecrypt 使用DES CFB模式进行解密
func DesDecrypt(encryptedData, key, iv []byte) ([]byte, error) {
	// 创建一个DES密码分组
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建一个CFB解密模式
	cfb := cipher.NewCFBDecrypter(block, iv)

	// 解密数据
	decrypted := make([]byte, len(encryptedData))
	cfb.XORKeyStream(decrypted, encryptedData)

	// 对解密后的数据进行去填充
	decrypted = pkcs7UnPadding(decrypted)

	return decrypted, nil
}

// PKCS7填充
func pkcs7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, paddingText...)
}

// PKCS7去填充
func pkcs7UnPadding(rawData []byte) []byte {
	length := len(rawData)
	unPadding := int(rawData[length-1])
	return rawData[:(length - unPadding)]
}
