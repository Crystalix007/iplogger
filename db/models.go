package db

import (
	"gorm.io/gorm"
	"time"
)

type URL struct {
	gorm.Model
	fullUrl string
	shortUrl string
}

type Logged struct {
	gorm.Model
	IP string
	timestamp time.Time
}
