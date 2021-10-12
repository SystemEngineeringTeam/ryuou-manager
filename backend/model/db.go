package model

import "time"

type Question struct {
	ID          int64     `gorm:"id"`
	Title       string    `gorm:"title"`
	Description string    `gorm:"description"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}
