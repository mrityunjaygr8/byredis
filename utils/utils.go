package utils

import (
	"encoding/binary"
	"errors"
)

const (
	MAX_MESSAGE_LENGTH      = 1024
	ErrMsgTooBig            = "The message is too big for this server."
	MAX_LENGTH_SECTION_SIZE = 4
)

func ReadData(raw []byte) (int, string, error) {
	sizeMessage := binary.LittleEndian.Uint16(raw[:MAX_LENGTH_SECTION_SIZE])

	if sizeMessage > MAX_MESSAGE_LENGTH {
		return 0, "", errors.New(ErrMsgTooBig)
	}
	return int(sizeMessage), string(raw[MAX_LENGTH_SECTION_SIZE:]), nil

}

func WriteData(message string) ([]byte, error) {
	sizeMessage := len(message)
	if sizeMessage > MAX_MESSAGE_LENGTH {
		return nil, errors.New(ErrMsgTooBig)
	}

	size := make([]byte, MAX_LENGTH_SECTION_SIZE)
	binary.LittleEndian.PutUint16(size, uint16(sizeMessage))

	finalMessage := size
	finalMessage = append(finalMessage, []byte(message)...)
	return []byte(finalMessage), nil
}
