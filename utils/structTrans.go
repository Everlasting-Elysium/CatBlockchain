package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func Int2Byte(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	CheckErr("Int2Byte", err)
	return buffer.Bytes()
}

func CheckErr(pos string, err error) {
	if err != nil {
		fmt.Println("error,pos", err, pos)
		os.Exit(-1)
	}
}
