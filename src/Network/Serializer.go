package main

func Deserialize(conn *connection, contentLen int32, proto int32, data []byte)  {
	if proto == pLogin{
		Login(conn)
	}else if	proto == pMove{
		if len(conn.data) > 0{
			conn.data = BytesJoin(conn.data, data)
		} else{
			conn.data = append(data)
		}
	} else if proto == pPlayerJoin{
		PlayerJoin(conn)
	}else if proto == pPlayerLeave{
		PlayerLeave(conn)
	} else if proto == pStartGame{
		InitWorld()
	}
}
