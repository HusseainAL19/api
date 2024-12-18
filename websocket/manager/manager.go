package managerWS

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	libErrors "iqdev/ss/libs/errors"
	managerwscomp "iqdev/ss/websocket/manager/managerWSComp"

	"github.com/gorilla/websocket"
)

type authManagerMsg struct {
	ManKey string `json:"manKey"`
}

type managerConnectionstype struct {
	connectionCounter int
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     nil}

func ManagerHandler() {

	managercConnList := managerConnectionstype{}
	//idCounter := 0
	tickerCounter := 0

	http.HandleFunc("/ws/manager", func(w http.ResponseWriter, r *http.Request) {
		conn, connErr := upgrader.Upgrade(w, r, nil)
		if connErr != nil {
			libErrors.ReturnError(conn)
			return
		}

		managercConnList.connectionCounter++

		defer conn.Close()

		var ticker *time.Ticker
		done := make(chan bool)

		for {
			_, message, msgError := conn.ReadMessage()
			//_, message, msgError := conn.NextReader()

			if msgError != nil {
				managercConnList.connectionCounter--
				if managercConnList.connectionCounter > 0 {
					ticker.Stop()
					done <- true
					return
				}
				ticker.Stop()
				done <- true
				tickerCounter = 0
				break
			}

			decodeGodKeyValue := authManagerMsg{}
			json.Unmarshal([]byte(message), &decodeGodKeyValue)
			managerInfo := managerwscomp.GetManagerProfile(decodeGodKeyValue.ManKey, conn)

			tickerCounter++
			//if tickerCounter == 1 {
			ticker = time.NewTicker(time.Duration(1) * time.Second)
			//}

			if managerInfo.ManagerExsist == false {
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
						fmt.Println("sheesh manager")
						managerwscomp.GetAllManagerInfo(conn, managerInfo.ManagerInfo)
					}
				}
			}()
		}
	})
}
