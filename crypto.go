package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	mathrand "math/rand"
)

func GetRandomInt(min, max int) int {
	var n int64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	mathrand.New(mathrand.NewSource(n))
	return mathrand.Intn(max-min) + min
}

func GetHMAC256(k *string) string {
	key := *k
	h := hmac.New(sha256.New, []byte(key))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
func EncryptAES(key []byte, data []byte) []byte {
	return data
}
