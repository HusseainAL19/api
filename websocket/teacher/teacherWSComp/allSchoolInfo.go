package schoolwscomp

import (
	"encoding/json"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"reflect"

	"github.com/gorilla/websocket"
)

type EventInstideStruct struct {
}

type AllTeacherInfoObj struct {
}

type ConnectionsList struct {
	GodConnectionNum   int
	GodConnectionsList []*websocket.Conn
}

type connectionsProfile struct {
}

func assembleInfo(SchoolProfile globalObject.SchoolsProfile,
	connection *websocket.Conn) AllTeacherInfoObj {
	allSchoolInfo := AllTeacherInfoObj{}

	if connection == nil {
		return allSchoolInfo
	}

	return allSchoolInfo
}

func GetAllSchoolInfo(connection *websocket.Conn,
	schoolProfile globalObject.SchoolsProfile,
	currCounter int, prevCounter int,
	schoolCurrent globalObject.SchoolsProfile,
	schoolPrev globalObject.SchoolsProfile,
	connCounter int) (bool, int, int, globalObject.SchoolsProfile) {

	var allInfoRow = assembleInfo(schoolProfile, connection)

	//schoolCurrent = allInfoRow.SchoolProfiles
	compareschool := reflect.DeepEqual(schoolCurrent, schoolPrev)

	currCounter = 0

	if connCounter > 3 {
		if currCounter == prevCounter && compareschool {
			return false, currCounter, prevCounter, schoolCurrent
		}
	}

	allInfoJson, ctjsonError := json.Marshal(allInfoRow)
	if ctjsonError != nil {
		libErrors.ReturnError(connection)
	}

	writeError := connection.WriteMessage(1, allInfoJson)
	if writeError != nil {
		libErrors.ReturnError(connection)
	}

	return true, currCounter, prevCounter, schoolPrev
}
