package model

import "time"

type Products struct {
	ID         uint64     `gorm:"column:id;primary_key;autoIncrement"`
	CategoryId uint64     `gorm:"not null;column:category_id"`
	Name       string     `gorm:"column:name;size:100;not null"`
	Price      float64    `gorm:"type:decimal(12,2); not null"`
	Stock      int        `gorm:"column:stock;not null;default:0"`
	CreatedAt  time.Time  `gorm:"column:created_at;autoCreateTime"`
	Category   Categories `gorm:"foreignKey:category_id;references:id"`
	Movement   []Movements  `gorm:"foreignKey:product_id;references:id"`
}
