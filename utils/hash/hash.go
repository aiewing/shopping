package hash

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

// 随机字符串
const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 使用当前时间创建seed
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// 创建salt
func CreateSalt() string {
	bytes := make([]byte, bcrypt.MaxCost)
	for i := range bytes {
		bytes[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(bytes)
}

// 使用bcrypt算法返回hash后密码
func HashPassword(passsword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passsword), bcrypt.DefaultCost)
	return string(bytes), err
}

// 检查密码是否相等
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// 创建SKU
func CreateSKU() string {
	bytes := make([]byte, bcrypt.MaxCost)
	for i := range bytes {
		bytes[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(bytes)
}
