package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// encoding an integer in a format suitable for binary operations or storage
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)                        //dynamic buffer to store the bytes
	err := binary.Write(buff, binary.BigEndian, num) //writing binary representation of int to buff buffer
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes() //returns the bytes from the buffer
}
