package main

import "grpc_practice/config"

func main() {
	config.ConnectToDatabase()
	config.APIConfig()
}