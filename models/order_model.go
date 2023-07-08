package models

import "time"

type Order struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	CustomerID  int64     `gorm:"not null" json:"customer_id"`
	ProductName string    `gorm:"not null" json:"product_name"`
	Quantity    int64     `gorm:"not null" json:"quantity"`
	TotalPrice  float64   `gorm:"not null" json:"total_price"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
