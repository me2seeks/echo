package tool

import (
	"crypto/md5"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

func Md5ByString(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		panic(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

// 加密密码并生成盐
func EncryptWithBcrypt(str string) (string, error) {
	// GenerateFromPassword 会自动生成盐并将其包含在哈希值中
	hashedStr, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedStr), nil
}

// 提取盐的函数
//
//nolint:unused
func extractSalt(hashed []byte) string {
	// bcrypt 哈希值的前 29 个字符包含了盐的信息
	return string(hashed[:29])
}

// 验证密码
func CheckStrHash(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
