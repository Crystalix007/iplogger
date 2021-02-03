package main

import iplogger "iplogger/server"

func main() {
	println("Starting iplogger...")
	server := iplogger.CreateServer()
	server.Run()
}