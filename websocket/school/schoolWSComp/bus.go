package schoolwscomp

import (
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllBusInfo(
	schoolId int,
	conn *websocket.Conn) []globalObject.BusProfile {

	BusProfile := []globalObject.BusProfile{}
	BusProfileTmp := globalObject.BusProfile{}

	// sql connection

	// query
	getManagerQuery := "SELECT * FROM bus where school_id = ?;"
	// sql connection
	connObj := sql.InitConnection()
	// query manager info
	rows, sqlError := connObj.Connections.Query(getManagerQuery, schoolId)

	if sqlError != nil {
		libErrors.ReturnError(conn)
	}

	defer rows.Close()
	defer connObj.Connections.Close()

	for rows.Next() {
		err := rows.Scan(
			&BusProfileTmp.BusId,
			&BusProfileTmp.BusName,
			&BusProfileTmp.BusDocumentId,
			&BusProfileTmp.BusKey,
			&BusProfileTmp.SchoolId)
		if err != nil {
      fmt.Println("bus error")
			libErrors.ReturnError(conn)
		}

		BusProfile = append(BusProfile, globalObject.BusProfile{
			BusProfileTmp.BusId,
			BusProfileTmp.BusName,
			BusProfileTmp.BusDocumentId,
			BusProfileTmp.BusKey,
			BusProfileTmp.SchoolId})

	}
	return BusProfile
}
