package server

type Logging interface {
	Log(message string)
}

type Server interface {
	Migrate()
	Run()
	GetLongUrl(shortUrl string) string
	NewUrl(longUrl string) string
}

type Database interface {
	GetFullUrlFromShort(shortUrl string) string
	CreateUrl(longUrl string) string
}