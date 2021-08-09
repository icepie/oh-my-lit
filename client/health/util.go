package health

import (
	"crypto/sha256"
	"encoding/hex"
	"math"
	"math/rand"
	"time"
)

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
