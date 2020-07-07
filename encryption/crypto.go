package encryption

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

//HmacSha256  用HMAC-SHA256算法对message签名
func HmacSha256(secret string,message string) []byte {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return h.Sum(nil)
}

//HmacSha256_base64 用HMAC-SHA256签名，然后将生成的签名串使用 Base64 进行编码
func HmacSha256_Base64(secret string, message string) string {
	sha := HmacSha256(message, secret)
	return base64.StdEncoding.EncodeToString(sha)
}
