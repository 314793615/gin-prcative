package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(tempStr))
}

func MakePassword(plainwd , salt string) string{
	return Md5Encode(plainwd + salt)
}

func ValidatePassword(plainwd, salt string, password string) bool{
	return Md5Encode(plainwd + salt) == password
}