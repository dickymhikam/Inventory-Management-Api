package model

import "time"

type Users struct {
	ID        uint64 `gorm:"primary_key:id;column:id;autoIncrement"`
	Name      string `gorm:"column:name;size:100;not null"`
	Email     string `gorm:"column:email;size:150;not null;unique"`
	Password  string `gorm:"column:password;size:225;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	Movement []Movements `gorm:"foreignKey:user_id;references:id"`
}