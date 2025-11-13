package models

import "time"

type UserSession struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	UserID     uint       `json:"user_id"`
	LoginTime  time.Time  `json:"login_time"`
	LogoutTime *time.Time `json:"logout_time,omitempty"`
	DeviceInfo string     `json:"device_info"`
}
