package work

import (
	"fmt"
	"net"
	"os"
)

func Start()  {
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
		go handleClient(conn)
	}
	
}

func handleClient(conn net.Conn){
	defer conn.Close()

	var data = GetSend(3, "i accept you")
	conn.Write(data)
	receive(conn)
}


func checkError(err error){
	if err != nil{
		fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func receive(conn net.Conn){
	for{
		buf := make([]byte, 4)
		_, err := conn.Read(buf)
		checkError(err)
		len := ByteToInt(buf)
		_, err = conn.Read(buf)
		checkError(err)
		buf = make([]byte, len)
		_, err = conn.Read(buf)
		fmt.Println(string(buf[:]))
		var data = GetSend(3, string(buf[:]))
		conn.Write(data)
	}
}


