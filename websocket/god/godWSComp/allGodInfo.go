package godwscomp

import (
	"encoding/json"
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"

	"github.com/gorilla/websocket"
)

type AllGodInfoObj struct {
	ShouldUpDate        bool
	GodProfile          globalObject.GodProfile
	ManagerProfiles     []globalObject.ManagersProfile
	DisProfiles         []globalObject.DisProfile
	SchoolOwnerProfiles []globalObject.SchoolOwnerProfile
	//SchoolProfiles      []globalObject.SchoolsProfile
	//TeacherProfiles     []globalObject.TeacherProfile
	//BusProfiles         []globalObject.BusProfile
}

type UpDateObj struct {
	GodUpdateCount         int
	ManagerUpdateCount     int
	DisUpdateCount         int
	SchoolOwnerUpdateCount int
	SchoolsUpdateCount     int
}

var PrevLen int
var CurrentLen int

func assembleInfo(godInfo globalObject.GodProfile,
	connection *websocket.Conn) AllGodInfoObj {
	allgodinfo := AllGodInfoObj{}
	if connection == nil {
		return allgodinfo
	}

	managerInfo := GetAllManagerInfo(godInfo, connection)
	disInfo := GetAllDisInfo(godInfo, connection)
	schoolOwnersInfo := GetAllSchoolOwnerInfo(godInfo, connection)

	CurrentLen = len(managerInfo) + len(disInfo) + len(schoolOwnersInfo)
	if PrevLen == CurrentLen {
		allgodinfo.ShouldUpDate = false
		return allgodinfo
	}
	fmt.Println("does not equle")

	allgodinfo.ShouldUpDate = true
	PrevLen = CurrentLen

	allgodinfo.GodProfile = godInfo
	allgodinfo.ManagerProfiles = managerInfo
	allgodinfo.DisProfiles = disInfo
	allgodinfo.SchoolOwnerProfiles = schoolOwnersInfo

	return allgodinfo
}

func GetAllGodInfo(connection *websocket.Conn,
	godInfo globalObject.GodProfile) {

	var allInfoRow AllGodInfoObj
	allInfoRow = assembleInfo(godInfo, connection)

	if allInfoRow.ShouldUpDate == false {
		return
	}

	allInfoJson, ctjsonError := json.Marshal(allInfoRow)
	if ctjsonError != nil {
		libErrors.ReturnError(connection)
	}

	writeError := connection.WriteMessage(1, allInfoJson)
	if writeError != nil {
		libErrors.ReturnError(connection)
	}
}
