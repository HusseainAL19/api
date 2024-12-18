package managerwscomp

import (
	"encoding/json"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"

	"github.com/gorilla/websocket"
)

type AllManagerInfoObj struct {
	ManagerProfiles     globalObject.ManagersProfile
	DisProfiles         []globalObject.DisProfile
	SchoolOwnerProfiles []globalObject.SchoolOwnerProfile
	SchoolProfiles      []globalObject.SchoolsProfile
	//TeacherProfiles     []globalObject.TeacherProfile
	//BusProfiles         []globalObject.BusProfile
}

type ConnectionsList struct {
	GodConnectionNum   int
	GodConnectionsList []*websocket.Conn
}

type connectionsProfile struct {
}

func assembleInfo(managerProfile globalObject.ManagersProfile,
	connection *websocket.Conn) AllManagerInfoObj {
	allManagerInfo := AllManagerInfoObj{}

	if connection == nil {
		return allManagerInfo
	}

	managerInfo := managerProfile
	//disInfo := GetAllDisInfo(godInfo, connection)
	//schoolOwnersInfo := GetAllSchoolOwnerInfo(godInfo, connection)

	//allgodinfo.GodProfile = godInfo
	allManagerInfo.ManagerProfiles = managerInfo
	//allgodinfo.DisProfiles = disInfo
	//allgodinfo.SchoolOwnerProfiles = schoolOwnersInfo

	return allManagerInfo
}

func GetAllManagerInfo(connection *websocket.Conn,
	managerInfo globalObject.ManagersProfile) {

	var allInfoRow AllManagerInfoObj
	allInfoRow = assembleInfo(managerInfo, connection)

	allInfoJson, ctjsonError := json.Marshal(allInfoRow)
	if ctjsonError != nil {
		libErrors.ReturnError(connection)
	}

	writeError := connection.WriteMessage(1, allInfoJson)
	if writeError != nil {
		libErrors.ReturnError(connection)
	}
}
