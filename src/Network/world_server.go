package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

var(
	NetWorkFrame *Queue
	quitSyncFrame chan bool
	actionData []byte
	lock sync.Mutex
	syncTimer time.Ticker
)


func Login(connection *connection)  {
	//var data = CombineSend(pLogin, []byte(strconv.Itoa(int(connection.uid)) + "login"))
	var data = CombineSend(pLogin, Int2Byte(connection.uid))
	connection.conn.Write(data)
	AddExistPLayer(mapConnected, connection)
	AddNewPlayer(connection)
}

func PlayerJoin(connection *connection)  {
	//var data = CombineSend(pPlayerJoin, []byte(strconv.Itoa(int(connection.uid)) + "join"))
	var data = CombineSend(pPlayerJoin, Int2Byte(connection.uid))
	broadCastExcept(connection.uid, data)
}

func PlayerLeave(connection *connection)  {
	//var data = CombineSend(pPlayerLeave, []byte(strconv.Itoa(int(connection.uid)) + "leave"))
	var data = CombineSend(pPlayerLeave, Int2Byte(connection.uid))
	broadCastExcept(connection.uid, data)
}

func  playerUpdate(connection *connection, data []byte)  {
	var send = BytesJoin(Int2Byte(connection.uid), data)
	setActionData(send)
}

func AddExistPLayer(conns mapConn, connection *connection) {
	var buf bytes.Buffer
	var count int
	for _, conn := range conns.m{
		if connection.uid != conn.uid{
			buf.Write(Int2Byte(conn.uid))
			count++
		}
	}
	if count == 0{
		return
	}
	var data = CombineSend(pAddPlayer, buf.Bytes())
	connection.conn.Write(data)

}

func AddNewPlayer(connection *connection)  {
	var data = CombineSend(pAddPlayer, Int2Byte(connection.uid))
	broadCastExcept(connection.uid, data)
}

func RemovePlayer(connection *connection) {
	var data = CombineSend(pRemovePlayer, Int2Byte(connection.uid))
	connection.conn.Write(data)
}

func  setActionData(new []byte)  {
	lock.Lock()
	defer lock.Unlock()
	if len(new) == 0{
		actionData = make([]byte, 0)
	} else{
		actionData = BytesJoin(actionData, new)
	}
}

func  clearActionData()  {
	lock.Lock()
	defer lock.Unlock()
}


func StartGame()  {
	var data = CombineSend(pStartGame, []byte("start game"))
	broadCast(data)
}

func  InitWorld()  {
	StartGame()
	NetWorkFrame = NewQueue()
	quitSyncFrame = make(chan bool)
	go func() {
		syncFrame()
	}()

	//Time.Sleep(time.Second)
	//quitSyncFrame <- true
	//close(quitSyncFrame)
}

func syncFrame1()  {
	ticker := time.NewTicker(time.Millisecond * 250)
	var wg sync.WaitGroup
	wg.Add(1)
	//s0 := time.Now()
	go func() {
		defer wg.Done()
		for _ = range ticker.C {
			select {
			case <-quitSyncFrame:
				fmt.Println("close receve")
				return
			default:
				data := collectFrameData()
				broadCast(data)
				NetWorkFrame.push(data)
			}

		}
	}()
	wg.Wait()
	//s1 := time.Now()
	//s2 := s0.Sub(s1)
	//fmt.Println("run time", s2)
}

//1.protocol move
//2.player num
//3.player uid, actionNums, action1,action2,action3
func  collectFrameData()[]byte  {
	var data []byte
	var num int
	num = len(mapConnected.m)
	var actions int
	for _,connection := range mapConnected.m{
		actions = len(connection.data) / 4
		if actions > 0 {
			data = BytesJoin(data, append(BytesJoin(Int2Byte(connection.uid), Int2Byte(int32(actions)), connection.data)))
			connection.data = make([]byte, 0)
		}else {
			data = BytesJoin(data, append(BytesJoin(Int2Byte(connection.uid), Int2Byte(1),Int2Byte(0))))
		}
	}

	return CombineSend(pMove, BytesJoin(Int2Byte(int32(num)), data))
}

func  syncFrame()  {
	ticker := time.NewTicker(time.Millisecond * 50)
	var wg sync.WaitGroup
	wg.Add(1)
	//s0 := time.Now()
	go func() {
		defer wg.Done()
		for _ = range ticker.C {
			select {
			case <-quitSyncFrame:
				fmt.Println("close syncFrame2")
				return
			default:
				var send []byte
				if	len(actionData) == 0{
					send = Int2Byte(0)
				} else{
					send = BytesJoin(Int2Byte(1), actionData)
				}
				var empty []byte
				setActionData(empty)
				send = CombineSend(pUpdate, send)
				broadCast(send)
				NetWorkFrame.push(send)
			}

		}
	}()

	wg.Wait()
	//s1 := time.Now()
	//s2 := s0.Sub(s1)
	//fmt.Println("run time", s2)
}
