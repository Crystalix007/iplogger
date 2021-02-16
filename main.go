package main

import iplogger "iplogger/server"

func main() {
	println("Starting iplogger...")
	server := iplogger.CreateServer()
	println("Checking migrations...")
	server.Migrate()
	println("Running server...")
	server.Run()
}