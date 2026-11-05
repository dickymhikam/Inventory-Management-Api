package model

import "time"

type Movements struct {
	ID        uint64    `gorm:"column:id;primary_key;autoIncrement"`
	ProductId uint64    `gorm:"column:product_id;not null"`
	UserId    uint64    `gorm:"column:user_id;not null"`
	Type      string    `gorm:"type:enum('IN','OUT');not null"`
	Quantity  int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	Product   Products  `gorm:"foreignKey:product_id;references:id"`
	User      Users     `gorm:"foreignKey:user_id;references:id"`
}
