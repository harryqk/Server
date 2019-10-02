package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"sync"
)

type connection struct {
	uid int32
	conn net.Conn
	data []byte
}

type mapConn struct {
	m map[int32] *connection
	//lock *sync.Mutex
	lock sync.RWMutex
}

type broadCastWork struct {
	*connection
	data []byte
}

func (m *broadCastWork)Task()  {
	m.connection.conn.Write(m.data)
}

func (c *mapConn) Get(key int32) *connection {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.m[key]
}

func (c *mapConn) Set(key int32, val *connection) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.m[key] = val
}

func (c *mapConn) Del(key int32) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.m, key)
}

func (c *mapConn) getLen() int  {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.m)
}

func (c *mapConn) Range(f func(key, value interface{}) bool) {
	c.lock.RLock()
	for k, v := range c.m {
		f(k, v)
	}
	c.lock.RUnlock()
}

var mapConnected = mapConn{
	m : make(map[int32] *connection),
}

func Start()  {
	runtime.GOMAXPROCS(6)
	//mapConnected = mapConn{
		//m : make(map[int32]connection),
	//}
	port := ":1500"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)


	checkError(err)
	fmt.Println("start")
	for{
		conn, err := listener.Accept()
		fmt.Println("client accept")
		if err != nil{
			continue
		}
		go onClientConnected(conn)
	}
	
}

func onClientConnected(conn net.Conn){
	defer conn.Close()
	newConnect := connection{
		uid : UidGen(),
		conn: conn,
	}
	mapConnected.Set(newConnect.uid, &newConnect)
	//var data = GetSend(3, "i accept you")
	//conn.Write(data)
	fmt.Println(strconv.Itoa(int(newConnect.uid)) + " client connected")
	//Login(newConnect)
	//PlayerJoin(newConnect)
	runReceive(&newConnect)
}

func onClientDisconnected(connect *connection){
	connect.conn.Close()
	mapConnected.Del(connect.uid)
	PlayerLeave(connect)

	if mapConnected.getLen() == 0 && quitSyncFrame != nil{
		quitSyncFrame <- true
		close(quitSyncFrame)
		quitSyncFrame = nil
	}
	fmt.Println(strconv.Itoa(int(connect.uid)) + "client close")
}


func checkError(err error){
	if err != nil{
		fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func runReceive(connect *connection)  {
	conn := connect.conn
	bufInt := make([]byte, 4)
	slice := make([]byte, 1024)
	for{

		_, err := conn.Read(bufInt)
		if err != nil{
			onClientDisconnected(connect)
			return
		}
		len := ByteToInt(bufInt)
		_, err = conn.Read(bufInt)
		if err != nil{
			onClientDisconnected(connect)
			return
		}
		cmd := ByteToInt(bufInt)
		slice1 := slice[0:len]
		_, err = conn.Read(slice1)
		if err != nil{
			onClientDisconnected(connect)
			return
		}
		//fmt.Println(string(slice1[:]))
		//var data = CombineSend(pMove, slice1[:])
		//var data = BytesJoin(Int2Byte(len), Int2Byte(cmd), slice1[:])
		//connect.data = append(data)
		Deserialize(connect, len, cmd, slice1[:])
		//broadCast(data)
		//broadCast(data)
		//conn.Write(data)
	}
}
