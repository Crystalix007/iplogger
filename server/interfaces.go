package server

type Logging interface {
	Log(message string)
}

type Server interface {
	Run()
}