package health

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math"
	"math/rand"
	"time"

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

func getSha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// GenRandomSafeTemp 随机生成安全体温
func GenRandomSafeTemp() float64 {
	rand.Seed(time.Now().Unix())
	min, max := 36.0, 37.0
	x := math.Pow10(1)
	return math.Trunc((min+rand.Float64()*(max-min))*x) / x
}
