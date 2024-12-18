package schoolOwnerwscomp

import (
	"encoding/json"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"reflect"

	"github.com/gorilla/websocket"
)

type AllSchoolOwnerInfoObj struct {
	SchoolOwnerProfiles globalObject.SchoolOwnerProfile
	SchoolProfiles      []globalObject.SchoolsProfile
	TeacherProfiles     []globalObject.TeacherProfile
	StudentProfiles     []globalObject.StudentProfile
	StudentPayments     []globalObject.StudentPaymentProfile
	//BusProfiles         []globalObject.BusProfile
}

type ConnectionsList struct {
	GodConnectionNum   int
	GodConnectionsList []*websocket.Conn
}

type connectionsProfile struct {
}

func assembleInfo(SchoolOwnerProfile globalObject.SchoolOwnerProfile,
	connection *websocket.Conn) AllSchoolOwnerInfoObj {
	allSchoolOwnerInfo := AllSchoolOwnerInfoObj{}

	if connection == nil {
		return allSchoolOwnerInfo
	}

	schoolOwnerInfo := SchoolOwnerProfile
	schoolsInfo := GetAllSchoolInfo(schoolOwnerInfo.SchoolOwnerId, connection)
	studentInfo := GetAllStudentInfo(schoolsInfo, connection)
	teacherInfo := GetAllTeaacherInfo(schoolsInfo, connection)
	studentPayment := GetAllstudentPaymentPaymentsInfo(schoolsInfo, connection)

	allSchoolOwnerInfo.SchoolOwnerProfiles = SchoolOwnerProfile
	allSchoolOwnerInfo.SchoolProfiles = schoolsInfo
	allSchoolOwnerInfo.StudentProfiles = studentInfo
	allSchoolOwnerInfo.TeacherProfiles = teacherInfo
	allSchoolOwnerInfo.StudentPayments = studentPayment

	return allSchoolOwnerInfo
}

func GetAllSchoolOwnerInfo(connection *websocket.Conn,
	schoolOwnerInfo globalObject.SchoolOwnerProfile,
	currCounter int, prevCounter int,
	soCurrent globalObject.SchoolOwnerProfile,
	soPrev globalObject.SchoolOwnerProfile,
	connCounter int) (bool, int, int, globalObject.SchoolOwnerProfile) {

	var allInfoRow AllSchoolOwnerInfoObj
	allInfoRow = assembleInfo(schoolOwnerInfo, connection)

	soCurrent = allInfoRow.SchoolOwnerProfiles
	compareDis := reflect.DeepEqual(soCurrent, soPrev)

	currCounter = len(
		allInfoRow.SchoolProfiles,
	) + len(
		allInfoRow.StudentProfiles,
	) + len(
		allInfoRow.TeacherProfiles,
	) + len(
		allInfoRow.StudentPayments,
	)

	if connCounter > 3 {
		if currCounter == prevCounter && compareDis {
			return false, currCounter, prevCounter, soCurrent
		}
	}

	allInfoJson, ctjsonError := json.Marshal(allInfoRow)
	if ctjsonError != nil {
		libErrors.ReturnError(connection)
	}

	prevCounter = currCounter
	soPrev = soCurrent

	writeError := connection.WriteMessage(1, allInfoJson)
	if writeError != nil {
		libErrors.ReturnError(connection)
	}
	return true, currCounter, prevCounter, soPrev
}
