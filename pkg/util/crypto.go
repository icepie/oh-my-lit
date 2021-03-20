package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"github.com/icepie/lit-edu-go/conf"
)

// 加密
func UserEncrypt(stuid string, password string) string {
	src := fmt.Sprintf("%s,%s", stuid, password)
	key := []byte(conf.ProConf.JWT.Key)
	iv := []byte(conf.ProConf.JWT.IV)
	result, err := Aes128Encrypt([]byte(src), key, iv)
	if err != nil {
		panic(err)
	}
	return base64.RawStdEncoding.EncodeToString(result)
}

// 解密
func UserDecrypt(src string) string {
	key := []byte(conf.ProConf.JWT.Key)
	iv := []byte(conf.ProConf.JWT.IV)
	var result []byte
	var err error
	result, err = base64.RawStdEncoding.DecodeString(src)
	if err != nil {
		panic(err)
	}
	origData, err := Aes128Decrypt(result, key, iv)
	if err != nil {
		panic(err)
	}

	return string(origData)
}

func Aes128Encrypt(origData, key []byte, IV []byte) ([]byte, error) {
	if key == nil || len(key) != 16 {
		return nil, nil
	}
	if IV != nil && len(IV) != 16 {
		return nil, nil
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, IV[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func Aes128Decrypt(crypted, key []byte, IV []byte) ([]byte, error) {
	if key == nil || len(key) != 16 {
		return nil, nil
	}
	if IV != nil && len(IV) != 16 {
		return nil, nil
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, IV[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
