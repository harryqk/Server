package work

import (
	"bytes"
	"encoding/binary"
)

func GetSend(protocol int32, content string) []byte {
	var c = []byte(content)
	var cmd = IntToByte(&protocol)
	var length = int32(len(c) + 8)
	var l = IntToByte(&length)
	return BytesCombine(l, cmd, c)
}

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}


func ByteToInt(by []byte, num *int32)  {
	buf := bytes.NewBuffer(by)
	binary.Read(buf, binary.BigEndian, num)
}

func IntToByte(num *int32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian,num)
	return buf.Bytes()
}
