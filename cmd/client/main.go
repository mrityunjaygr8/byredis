package main

import (
	"net"

	"github.com/mrityunjaygr8/byredis/utils"

	"github.com/sirupsen/logrus"
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

const (
	SERVER_TYPE = "tcp"
	SERVER_HOST = "localhost"
	SERVER_PORT = "6739"
)

func main() {
	str := "this is the end"
	msg, err := utils.WriteData(str)
	if err != nil {
		log.Fatal(err)
	}

	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		log.Fatal("An error occurred when contacting server:", err)
	}
	defer connection.Close()

	log.Println("Connecting to:", SERVER_HOST+":"+SERVER_PORT)
	_, err = connection.Write(msg)
	buffer := make([]byte, utils.MAX_MESSAGE_LENGTH+utils.MAX_LENGTH_SECTION_SIZE)
	_, err = connection.Read(buffer)
	if err != nil {
		log.Println("Error reading:", err.Error())
	}

	_, resp, err := utils.ReadData(buffer)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response: ", resp)
}
