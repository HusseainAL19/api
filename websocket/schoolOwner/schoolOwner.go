package schoolOwnerWS

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	schoolOwnerwscomp "iqdev/ss/websocket/schoolOwner/schoolOwnerWSComp"

	"github.com/gorilla/websocket"
)

type authManagerMsg struct {
	SchoolOwnerKey string `json:"soKey"`
}

type managerConnectionstype struct {
	connectionCounter int
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     nil}

func SchoolOwnerHandler() {

	managercConnList := managerConnectionstype{}
	//idCounter := 0
	tickerCounter := 0

	http.HandleFunc("/ws/schoolOwner", func(w http.ResponseWriter, r *http.Request) {
		conn, connErr := upgrader.Upgrade(w, r, nil)
		fmt.Println(connErr)

		if connErr != nil {
			fmt.Println("error connection upgrae")
			fmt.Println(connErr)
			libErrors.ReturnError(conn)
			return
		}

		var prevCount int = 0
		var currentCount int = 1
		var sendConn int = 0

		var soCurrent globalObject.SchoolOwnerProfile
		var soPrev globalObject.SchoolOwnerProfile

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
			schoolOwnerInfo := schoolOwnerwscomp.GetSchoolOwnerInfo(
				decodeDisKeyValue.SchoolOwnerKey,
				conn,
			)

			tickerCounter++
			//if tickerCounter == 1 {
			ticker = time.NewTicker(time.Duration(1) * time.Second)
			//}

			if schoolOwnerInfo.SchoolOwnerExsist == false {
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
						shouldUpdate, rcurrentcount, rprevCount, rsoPrev := schoolOwnerwscomp.GetAllSchoolOwnerInfo(
							conn,
							schoolOwnerInfo.SchoolOwnerInfo,
							currentCount,
							prevCount,
							soCurrent,
							soPrev,
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
						soPrev = rsoPrev

						fmt.Println("sending so data")
					}
				}
			}()
		}
	})
}
