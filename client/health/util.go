package health

import (
	"crypto/sha256"
	"encoding/hex"
)

func getSha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
