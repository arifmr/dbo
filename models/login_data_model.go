package models

import "time"

type LoginData struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	CustomerID int64     `gorm:"not null" json:"customer_id"`
	LoginTime  time.Time `gorm:"autoCreateTime" json:"login_time"`
	IPAddress  string    `gorm:"size:45" json:"ip_address"`
	Token      string    `gorm:"not null" json:"token"`
}

type AuthData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
