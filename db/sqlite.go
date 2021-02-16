package db

import (
	"crypto/sha256"
	"errors"
)
import "gorm.io/gorm"
import "gorm.io/driver/sqlite"

type DB struct {
	underlying *gorm.DB
}

func (db *DB) GetFullUrlFromShort(shortUrl string) (string, error) {
	var url Url
	val := db.underlying.
		Where("short_url = ?", shortUrl).
		First(&url)
	if val != nil {
		return url.ShortUrl, nil
	} else {
		return "", errors.New("record not found")
	}
}

// CreateUrl creates a shortened url for the given longUrl and returns the shortened url
func (db *DB) CreateUrl(longUrl string) string {
	sha := sha256.New()
	sha.Write([]byte(longUrl))
	url:= Url{
		LongUrl: longUrl,
		ShortUrl: string(sha.Sum(nil)[0:10]),
	}
	db.underlying.Create(&url)
	return url.ShortUrl
}

func (db *DB) Migrate() {
	db.underlying.AutoMigrate(&Url{})
}

func New() (DB, error) {
	db, err := gorm.Open(sqlite.Open("logger.db"), &gorm.Config{})
	return DB{db}, err
}

