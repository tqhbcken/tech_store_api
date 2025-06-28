package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint           `json:"id" gorm:"primaryKey" column:"id"`
	CreatedAt time.Time      `json:"created_at" column:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" column:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" column:"deleted_at"`
}
