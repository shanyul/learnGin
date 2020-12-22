package models

// Auth 用户信息
type Auth struct {
	ID	int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `josn:"password"`
}

// CheckAuth 验证用户
func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)

	if auth.ID > 0 {
		return true
	}

	return false
}