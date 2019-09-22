package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"sync/atomic"
	"time"
)

var uid int32

func GetSend(protocol int32, content string) []byte {
	var c = []byte(content)
	var cmd = Int2Byte(protocol)
	var length = int32(len(c))
	var l = Int2Byte(length)
	return BytesJoin(l, cmd, c)
}

func CombineSend(protocol int32, content []byte) []byte {
	var cmd = Int2Byte(protocol)
	var length = int32(len(content))
	var l = Int2Byte(length)
	return BytesJoin(l, cmd, content)
}



func BytesJoin(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func BytesCombine(pBytes ...[]byte) []byte {
	var total int
	var size int
	count := len(pBytes)
	counts := make([]int, count)

	for i := 0; i < count; i++{
		size = len(pBytes[i])
		total += size
		counts[i] = size
	}
	str := make([]byte, total)
	var pos int
	for i := 0; i < count; i++ {
		size = counts[i]
		for j := 0; j < size; j++ {
			str[pos] = pBytes[i][j]
			pos++
		}
	}
	return str
}

func BytesCopy(pBytes ...[]byte) []byte {
	var buf bytes.Buffer
	for _, data := range pBytes{
		var buf bytes.Buffer
		buf.Write(data)
	}
	return buf.Bytes()
}


func ByteToInt(by []byte) int32 {
	buf := bytes.NewBuffer(by)
	var num int32
	binary.Read(buf, binary.BigEndian, &num)
	return num
}

func IntToByte(num *int32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian,num)
	return buf.Bytes()
}

//bigEndian
func Int2Byte(v int32)(ret []byte){
	var src = make([]byte, 4)
	src[0] = byte(v >> 24)
	src[1] = byte(v >> 16)
	src[2] = byte(v >> 8)
	src[3] = byte(v)
	return src
}

//bigEndian
func Byte2Int(b []byte)int32{
	return int32(b[3]) | int32(b[2])<<8 | int32(b[1])<<16 | int32(b[0])<<24
}

func UidGen()int32  {
	return atomic.AddInt32(&uid, 1)
}

func TestCombine()  {
	count := 1000000
	oneSerail := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	str0 := make([][]byte, 8)
	str0[0] = oneSerail[0:4]
	str0[1] = oneSerail[4:6]
	str0[2] = oneSerail[6:10]
	str0[3] = oneSerail[0:2]
	str0[4] = oneSerail[2:6]
	str0[5] = oneSerail[6:8]
	str0[6] = oneSerail[0:4]
	str0[7] = oneSerail[8:10]
	s0 := time.Now()
	for i := 0; i < count; i++ {
		BytesJoin(str0[0], str0[1], str0[2])
	}
	e0 := time.Now()
	d0 := e0.Sub(s0)
	fmt.Printf("time of way(0)=%v\n", d0)

	s1 := time.Now()
	for i := 0; i < count; i++ {
		BytesCombine(str0[0], str0[1], str0[2])
	}
	e1 := time.Now()
	d1 := e1.Sub(s1)
	fmt.Printf("time of way(1)=%v\n", d1)

	s2 := time.Now()
	for i := 0; i < count; i++ {
		BytesCopy(str0[0], str0[1], str0[2])
	}
	e2 := time.Now()
	d2 := e2.Sub(s2)
	fmt.Printf("time of way(2)=%v\n", d2)
	oneSerail1 := []byte{0, 1, 2}
	a := BytesJoin(oneSerail1, oneSerail1)
	b := BytesCombine(oneSerail1, oneSerail1)
	fmt.Print(string(a))
	fmt.Print(string(b))
}

func TestInt2Byte()  {
	var count = 10000000
	var num int32
	num = 32
	s0 := time.Now()
	for i := 0; i < count; i++ {
		Int2Byte(num)
	}
	e0 := time.Now()
	d0 := e0.Sub(s0)
	fmt.Printf("time of way(0)=%v\n", d0)

	s2 := time.Now()
	for i := 0; i < count; i++ {
		IntToByte(&num)
	}
	e2 := time.Now()
	d2 := e2.Sub(s2)
	fmt.Printf("time of way(2)=%v\n", d2)


	fmt.Println(Int2Byte(num))
	fmt.Println(IntToByte(&num))
}

func TestByte2Int()  {
	var count = 10000000
	oneSerail := []byte{0, 0, 0, 32}
	s0 := time.Now()
	for i := 0; i < count; i++ {
		Byte2Int(oneSerail)
	}
	e0 := time.Now()
	d0 := e0.Sub(s0)
	fmt.Printf("time of way(0)=%v\n", d0)

	s2 := time.Now()
	for i := 0; i < count; i++ {
		Byte2Int(oneSerail)
	}
	e2 := time.Now()
	d2 := e2.Sub(s2)
	fmt.Printf("time of way(2)=%v\n", d2)


	fmt.Println(Byte2Int(oneSerail))
	fmt.Println(Byte2Int(oneSerail))
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
