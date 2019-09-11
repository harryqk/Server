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
}

func checkError(err error){
	if err != nil{
		fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}


