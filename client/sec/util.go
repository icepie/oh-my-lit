package sec

import (
	"encoding/base64"

	"github.com/wumansgy/goEncrypt"
)

// 登录参数加密
func loginCrypto(plain string, key string, iv string) (data string, err error) {

	cryptText, err := goEncrypt.AesCbcEncrypt([]byte(plain), []byte(key), []byte(iv))
	if err != nil {
		return
	}

	data = base64.StdEncoding.EncodeToString(cryptText)

	return

}
