package main

import (
	"net"

	"github.com/mrityunjaygr8/byredis/utils"
	"github.com/sirupsen/logrus"
)

const (
	SERVER_TYPE = "tcp"
	SERVER_HOST = "localhost"
	SERVER_PORT = "6739"
)

func init() {
	log = logrus.WithFields(logrus.Fields{
		"host": SERVER_HOST,
		"port": SERVER_PORT,
	})
	log.Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

}

var log *logrus.Entry

func main() {

	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		log.Fatal("An error occurred when creating server:", err)
	}
	defer server.Close()
	log.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	log.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		log.Println("client connected")
		go processClient(connection)
	}
}

func processClient(connection net.Conn) {
	buffer := make([]byte, utils.MAX_MESSAGE_LENGTH+utils.MAX_LENGTH_SECTION_SIZE)
	_, err := connection.Read(buffer)
	if err != nil {
		log.Println("Error reading:", err)
	}
	_, msg, err := utils.ReadData(buffer)
	log.Println("Received: ", msg)
	resp, err := utils.WriteData("world")
	if err != nil {
		log.Fatal(err)
	}
	_, err = connection.Write(resp)
	connection.Close()
}
