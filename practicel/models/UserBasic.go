package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name       string `json:"name"`
	Password   string `json:"password"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
	ClientIP   string `json:"clientIP"`
	ClientPort string `json:"clientPort"`
	LoginTime  uint64 `json:"loginTime"`
	LogOutTime uint64 `json:"logOutTime"`
	IsLogout   bool   `json:"isLogout"`
	DeviceInfo string `json:"deviceInfo"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
