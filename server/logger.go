package server

type Logger struct {
}

func (lg Logger) Log(message string) {
	println(message)
}

func GetLogger() Logging {
	return Logger{}
}
