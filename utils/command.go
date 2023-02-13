package utils

import (
	"encoding/binary"
	"net"
)

const (
	GET = "get"
	SET = "set"
	DEL = "del"
)

func EncodeCommand(args []string) []byte {
	raw := make([]byte, MAX_LENGTH_SECTION_SIZE)

	log.Debug(args, len(args))
	binary.LittleEndian.PutUint16(raw, uint16(len(args)))
	for _, arg := range args {
		tmp, err := WriteData(arg)
		if err != nil {
			log.Errorln(err)
		}

		raw = append(raw, tmp...)
	}

	return raw
}

func DecodeCommand(raw []byte) []string {
	words := make([]string, 0)
	sizeMessage := binary.LittleEndian.Uint16(raw[:MAX_LENGTH_SECTION_SIZE])
	start := MAX_LENGTH_SECTION_SIZE
	for i := 0; i < int(sizeMessage); i++ {
		sizeWord := binary.LittleEndian.Uint16(raw[start : start+MAX_LENGTH_SECTION_SIZE])
		start = start + MAX_LENGTH_SECTION_SIZE
		word := raw[start : start+int(sizeWord)]
		words = append(words, string(word))
		start = start + int(sizeWord)
	}

	return words
}

func SendCommand(cmds []byte) string {

	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		log.Fatal("An error occurred when contacting server:", err)
	}
	defer connection.Close()

	log.Debug("Connecting to:", SERVER_HOST+":"+SERVER_PORT)
	_, err = connection.Write(cmds)
	buffer := make([]byte, MAX_MESSAGE_LENGTH+MAX_LENGTH_SECTION_SIZE)
	_, err = connection.Read(buffer)
	if err != nil {
		log.Error("Error reading:", err.Error())
	}

	_, resp, err := ReadData(buffer)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
