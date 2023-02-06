package utils

import "encoding/binary"

func EncodeCommand(args []string) []byte {
	raw := make([]byte, MAX_LENGTH_SECTION_SIZE)

	log.Println(args, len(args))
	binary.LittleEndian.PutUint16(raw, uint16(len(args)))
	for _, arg := range args {
		// raw = append(raw, Write)
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
