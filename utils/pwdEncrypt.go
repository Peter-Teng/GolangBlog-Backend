package utils

import (
	"MarvelousBlog-Backend/config"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"unsafe"
)

func Encrypt(username string, pwd string) (string, error) {
	salt := Str2bytes(config.SaltPrefix + username + config.SaltPostfix)
	encryptedPwd, err := scrypt.Key([]byte(pwd), salt, 32768, 8, 1, 32)
	return base64.StdEncoding.EncodeToString(encryptedPwd), err
}

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
