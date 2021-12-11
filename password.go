package util

import "golang.org/x/crypto/bcrypt"

// GeneratePassword 加密
func GeneratePassword(password string)  ([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ComparePasswords 验证密码
func ComparePasswords(hashedPassword []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return false
	}
	return true
}