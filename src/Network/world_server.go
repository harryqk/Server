package main

func Login(connection connection)  {
	//var data = CombineSend(pLogin, []byte(strconv.Itoa(int(connection.uid)) + "login"))
	var data = CombineSend(pLogin, Int2Byte(connection.uid))
	connection.conn.Write(data)
}

func PlayerJoin(connection connection)  {
	//var data = CombineSend(pPlayerJoin, []byte(strconv.Itoa(int(connection.uid)) + "join"))
	var data = CombineSend(pPlayerJoin, Int2Byte(connection.uid))
	broadCastExcept(connection.uid, data)
}

func PlayerLeave(connection connection)  {
	//var data = CombineSend(pPlayerLeave, []byte(strconv.Itoa(int(connection.uid)) + "leave"))
	var data = CombineSend(pPlayerLeave, Int2Byte(connection.uid))
	broadCastExcept(connection.uid, data)
}
