package main

import "strconv"

func Login(connection connection)  {
	var data = CombineSend(pLogin, []byte(strconv.Itoa(int(connection.uid)) + "login"))
	connection.conn.Write(data)
}

func PlayerJoin(connection connection)  {
	var data = CombineSend(pPlayerJoin, []byte(strconv.Itoa(int(connection.uid)) + "join"))
	broadCastExcept(connection.uid, data)
}

func PlayerLeave(connection connection)  {
	var data = CombineSend(pPlayerLeave, []byte(strconv.Itoa(int(connection.uid)) + "leave"))
	broadCastExcept(connection.uid, data)
}
