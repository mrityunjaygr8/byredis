package main

import (
	cobra "github.com/mrityunjaygr8/byredis/cmd/client/cobra"

	"github.com/mrityunjaygr8/byredis/utils"
)

const (
	SERVER_TYPE = "tcp"
	SERVER_HOST = "localhost"
	SERVER_PORT = "6739"
)

func main() {
	utils.SetupLogger(SERVER_HOST, SERVER_PORT)

	cobra.Execute()
	// str := "this is the end"
	// msg, err := utils.WriteData(str)
	// if err != nil {
	// 	utils.Log.Fatal(err)
	// }
	//
	// connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	// if err != nil {
	// 	utils.Log.Fatal("An error occurred when contacting server:", err)
	// }
	// defer connection.Close()
	//
	// utils.Log.Println("Connecting to:", SERVER_HOST+":"+SERVER_PORT)
	// _, err = connection.Write(msg)
	// buffer := make([]byte, utils.MAX_MESSAGE_LENGTH+utils.MAX_LENGTH_SECTION_SIZE)
	// _, err = connection.Read(buffer)
	// if err != nil {
	// 	utils.Log.Println("Error reading:", err.Error())
	// }
	//
	// _, resp, err := utils.ReadData(buffer)
	// if err != nil {
	// 	utils.Log.Fatal(err)
	// }
	// utils.Log.Println("Response: ", resp)
}
