// 깃헙 시그니처 검증
package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

// 웹훅 github 시그니처 검증
func VerifySignature(body []byte, signature string) bool {
	secret := os.Getenv("GITHUB_WEBHOOK_SECRET")
	if secret == "" || signature == "" {
		return false
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	hash := mac.Sum(nil)

	expected := "sha256=" + hex.EncodeToString(hash)

	return hmac.Equal([]byte(signature), []byte(expected))
}
