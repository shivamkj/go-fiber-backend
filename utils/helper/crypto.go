package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHMAC(data string, key []byte) string {
	hmacHash := hmac.New(sha256.New, key)
	hmacHash.Write([]byte(data))
	hashInBytes := hmacHash.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}
