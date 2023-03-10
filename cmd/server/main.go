package main

import (
	"net"

	"github.com/mrityunjaygr8/byredis/utils"
)

const (
	SERVER_TYPE = "tcp"
	SERVER_HOST = "localhost"
	SERVER_PORT = "6739"
)

func main() {

	utils.SetupLogger(SERVER_HOST, SERVER_PORT)
	log := utils.GetLogger()

	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		log.Fatal("An error occurred when creating server:", err)
	}
	defer server.Close()

	loop := utils.NewLoop()
	loop.RunLoop()
	log.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	log.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		log.Println("client connected")
		loop.AddEvent(connection)
	}
}
