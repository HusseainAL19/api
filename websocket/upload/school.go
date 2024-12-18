package upload

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	schoolwscomp "iqdev/ss/websocket/school/schoolWSComp"

	"github.com/gorilla/websocket"
)

type authManagerMsg struct {
	SchoolKey string `json:"schoolKey"`
	VideoName string `json:"schoolKey"`
	ListId string | nil`json:"schoolKey"`
}

type managerConnectionstype struct {
	connectionCounter int
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     nil}

func SchoolUpHandler() {

	managercConnList := managerConnectionstype{}
	//idCounter := 0
	tickerCounter := 0

	http.HandleFunc("/upload/uVideo", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Expect") == "100-continue" {
			w.WriteHeader(http.StatusContinue)
		}
		r.ParseMultipartForm(100000000)

		conn, connErr := upgrader.Upgrade(w, r, nil)
		if connErr != nil {
			libErrors.ReturnError(conn)
			return
		}

		managercConnList.connectionCounter++

		defer conn.Close()

		var ticker *time.Ticker
		done := make(chan bool)

		var prevCount int = 0
		var currentCount int = 1
		var sendConn int = 0

		var schoolCurrent globalObject.SchoolsProfile
		var schoolPrev globalObject.SchoolsProfile

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
			schoolInfo := schoolwscomp.GetSchoolInfo(decodeDisKeyValue.SchoolKey, conn)

			tickerCounter++
			//if tickerCounter == 1 {
			ticker = time.NewTicker(time.Duration(1) * time.Second)
			//}

			if schoolInfo.SchoolExsist == false {
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
						fmt.Println("sheesh dis")
						schoolwscomp.GetAllSchoolInfo(
							conn,
							schoolInfo.SchoolProfile,
							currentCount,
							prevCount,
							schoolCurrent,
							schoolPrev,
							sendConn,
						)
					}
				}
			}()
		}
	})
}
