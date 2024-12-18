package disWS

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	diswscomp "iqdev/ss/websocket/dis/disWSComp"

	"github.com/gorilla/websocket"
)

type authManagerMsg struct {
	DisKey string `json:"disKey"`
}

type managerConnectionstype struct {
	connectionCounter int
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     nil}

func DisHandler() {

	managercConnList := managerConnectionstype{}
	//idCounter := 0
	tickerCounter := 0

	http.HandleFunc("/ws/dis", func(w http.ResponseWriter, r *http.Request) {

		conn, connErr := upgrader.Upgrade(w, r, nil)
		fmt.Println("starting connection")
		if connErr != nil {
			libErrors.ReturnError(conn)
			return
		}

		fmt.Println(conn)

		var prevCount int = 0
		var currentCount int = 1
		var sendConn int = 0

		var disCurrent globalObject.DisProfile
		var disPrev globalObject.DisProfile

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

			decodeDisKeyValue := authManagerMsg{}
			json.Unmarshal([]byte(message), &decodeDisKeyValue)
			disInfo := diswscomp.GetDisInfo(decodeDisKeyValue.DisKey, conn)

			tickerCounter++
			//if tickerCounter == 1 {
			ticker = time.NewTicker(time.Duration(1) * time.Second)
			//}

			if disInfo.DisExsist == false {
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
						shouldUpdate, rcurrentcount, rprevCount, rdisPrev := diswscomp.GetAllDisInfo(
							conn,
							disInfo.DisProfile,
							currentCount,
							prevCount,
							disCurrent,
							disPrev,
							sendConn,
						)
						if sendConn < 5 {
							sendConn++
						}
						if shouldUpdate == false {
							continue
						}
						currentCount = rcurrentcount
						prevCount = rprevCount
						disPrev = rdisPrev
					}
				}
			}()
		}
	})
}
