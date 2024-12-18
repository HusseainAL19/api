package diswscomp

import (
	"encoding/json"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"reflect"

	"github.com/gorilla/websocket"
)

type AllDisInfoObj struct {
	DisProfiles         globalObject.DisProfile           `json:"disProfile"`
	SchoolOwnerProfiles []globalObject.SchoolOwnerProfile `json:"school_owner_profile`
	SchoolProfiles      []globalObject.SchoolsProfile     `json:"schoolsProfile"`
	//TeacherProfiles     []globalObject.TeacherProfile
	//BusProfiles         []globalObject.BusProfile
}

type ConnectionsList struct {
	GodConnectionNum   int
	GodConnectionsList []*websocket.Conn
}

type connectionsProfile struct {
}

func assembleInfo(disProfile globalObject.DisProfile,
	connection *websocket.Conn) AllDisInfoObj {
	allDisInfo := AllDisInfoObj{}

	if connection == nil {
		return allDisInfo
	}

	disInfo := disProfile
	schoolOwnerInfo := GetAllSchoolOwnerInfo(disProfile.DisId, connection)
	schoolsInfo := GetAllSchoolInfo(disProfile.DisId, connection)

	allDisInfo.DisProfiles = disInfo
	allDisInfo.SchoolOwnerProfiles = schoolOwnerInfo
	allDisInfo.SchoolProfiles = schoolsInfo

	return allDisInfo
}

func GetAllDisInfo(connection *websocket.Conn,
	disInfo globalObject.DisProfile,
	currCounter int, prevCounter int,
	disCurrent globalObject.DisProfile,
	disPrev globalObject.DisProfile,
	connCounter int) (bool, int, int, globalObject.DisProfile) {
	var allInfoRow AllDisInfoObj
	allInfoRow = assembleInfo(disInfo, connection)

	disCurrent = allInfoRow.DisProfiles
	compareDis := reflect.DeepEqual(disCurrent, disPrev)

	currCounter = len(allInfoRow.SchoolProfiles) + len(allInfoRow.SchoolOwnerProfiles)

	if connCounter > 3 {
		if currCounter == prevCounter && compareDis {
			return false, currCounter, prevCounter, disCurrent
		}
	}
	allInfoJson, ctjsonError := json.Marshal(allInfoRow)
	if ctjsonError != nil {
		libErrors.ReturnError(connection)
	}

	prevCounter = currCounter
	disPrev = disCurrent

	writeError := connection.WriteMessage(1, allInfoJson)
	if writeError != nil {
		libErrors.ReturnError(connection)
	}
	return true, currCounter, prevCounter, disPrev
}
