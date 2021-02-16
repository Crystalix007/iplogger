package db

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	LongUrl string
	ShortUrl string `gorm:"index"`
}

