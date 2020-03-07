package util

import (
	"crypto/md5"
	_ "encoding/base64"
	"encoding/hex"
)

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))

	return hex.EncodeToString(hash.Sum(nil))
}