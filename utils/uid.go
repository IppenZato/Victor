package utils

import (
	"crypto/rand"
	"io"
	"fmt"
	"crypto/sha256"
	"encoding/hex"
)

func newUUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return ""
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func GenerateSignature(payload []byte, apiSecret string) string {
	h := sha256.New()
	h.Write(payload)
	h.Write([]byte(apiSecret))
	return hex.EncodeToString(h.Sum(nil))
}