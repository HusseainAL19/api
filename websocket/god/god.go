package godWS

import (
	"encoding/json"
	"net/http"
	"time"

	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	godwscomp "iqdev/ss/websocket/god/godWSComp"

	"github.com/gorilla/websocket"
)

type authGodMsg struct {
	GodKey string `json:"godKey"`
}

type godProfileWithId struct {
	ConnId        int
	GodProfile    globalObject.GodProfile
	Connection    *websocket.Conn
	ConnectionErr error
	MsgError      error
}

type godConnectionstype struct {
	connectionCounter  int
	connectionProfiles []godProfileWithId
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     nil}

func GodWSHander() {

	godConnList := godConnectionstype{}
	//idCounter := 0
	tickerCounter := 0

	http.HandleFunc("/ws/god", func(w http.ResponseWriter, r *http.Request) {
		conn, connErr := upgrader.Upgrade(w, r, nil)
		if connErr != nil {
			libErrors.ReturnError(conn)
			return
		}

		godConnList.connectionCounter++

		defer conn.Close()

		var ticker *time.Ticker
		done := make(chan bool)

		godwscomp.CurrentLen = 0
		godwscomp.PrevLen = 0

		for {
			_, message, msgError := conn.ReadMessage()
			//_, message, msgError := conn.NextReader()

			if msgError != nil {
				godConnList.connectionCounter--
				if godConnList.connectionCounter > 0 {
					ticker.Stop()
					done <- true
					break
				}
				ticker.Stop()
				done <- true
				tickerCounter = 0
			}

			decodeGodKeyValue := authGodMsg{}
			json.Unmarshal([]byte(message), &decodeGodKeyValue)
			godInfo := godwscomp.GetGodInfo(decodeGodKeyValue.GodKey)

			tickerCounter++
			//if tickerCounter == 1 {
			ticker = time.NewTicker(time.Duration(1) * time.Second)
			//}

			if godInfo.GodExsist == false {
				if conn != nil {
					libErrors.ReturnUnAuth(conn)
				}
				ticker.Stop()
				done <- true
				break
			}

			go func() {
				for {
					select {
					case <-done:
						ticker.Stop()
						return
					case <-ticker.C:
						godwscomp.GetAllGodInfo(conn, godInfo.GodInfo)
					}
				}
			}()
		}
	})
}
