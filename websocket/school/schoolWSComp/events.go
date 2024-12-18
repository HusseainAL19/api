package schoolwscomp

import (
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllEventsInfo(
	schoolId int,
	conn *websocket.Conn) []globalObject.SchoolEvents {

	schoolEvent := []globalObject.SchoolEvents{}
	schoolEventtmp := globalObject.SchoolEvents{}

	// sql connection

	// query
	getManagerQuery := "SELECT * FROM school_events where school_id = ?;"
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
			&schoolEventtmp.SchoolEventId,
			&schoolEventtmp.SchoolEventTitle,
			&schoolEventtmp.SchoolEventPerson,
			&schoolEventtmp.SchoolEventDesc,
			&schoolEventtmp.SchoolId,
			&schoolEventtmp.SchoolEventRegisterDate)
		if err != nil {
      fmt.Println("bus error")
			libErrors.ReturnError(conn)
		}

		schoolEvent = append(schoolEvent, globalObject.SchoolEvents{
			schoolEventtmp.SchoolEventId,
			schoolEventtmp.SchoolEventTitle,
			schoolEventtmp.SchoolEventPerson,
			schoolEventtmp.SchoolEventDesc,
			schoolEventtmp.SchoolId,
			schoolEventtmp.SchoolEventRegisterDate})

	}
	return schoolEvent
}
