package model

import "time"

type Categories struct {
	ID          uint64 `gorm:"column:id;primary_key;autoIncrement"`
	Name        string `gorm:"column:name;size:100;not null;unique"`
	Description string `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	Product 	[]Products `gorm:"foreignKey:category_id;references:id"`
}