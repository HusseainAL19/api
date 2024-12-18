package schoolOwnerwscomp

import (
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllStudentInfo(
	schoolsProfile []globalObject.SchoolsProfile,
	conn *websocket.Conn) []globalObject.StudentProfile {

	StudentProfile := []globalObject.StudentProfile{}
	Studenttmp := globalObject.StudentProfile{}

	// sql connection

	//getSchoolOwnerQuery := "SELECT * FROM school where  school_owner_id = ?;"
	getStudentQuery := "SELECT * FROM student where  school_id = ?;"

	// query
	for _, element := range schoolsProfile {
		// sql connection
		connObj := sql.InitConnection()
		// query manager info
		rows, sqlError := connObj.Connections.Query(getStudentQuery, element.SchoolId)

		if sqlError != nil {
			libErrors.ReturnError(conn)
		}

		defer rows.Close()
		defer connObj.Connections.Close()

		for rows.Next() {
			err := rows.Scan(
				&Studenttmp.StudentId,
				&Studenttmp.StudentFullName,
				&Studenttmp.StudentBirthDate,
				&Studenttmp.StudentParentFullName,
				&Studenttmp.StudentPhoneNumber,
				&Studenttmp.StudentParentPhoneNumber,
				&Studenttmp.StudentLocation,
				&Studenttmp.StudentCurrentLocation,
				&Studenttmp.StudentClass,
				&Studenttmp.StudentStudyGroupId,
				&Studenttmp.StudentIdBack,
				&Studenttmp.StudentIdFront,
				&Studenttmp.StudentDeviceType,
				&Studenttmp.StudentDeviceOsNum,
				&Studenttmp.StudentActive,
				&Studenttmp.StudentKey,
				&Studenttmp.StudentOverAllNum,
				&Studenttmp.SchoolId,
				&Studenttmp.DisId,
				&Studenttmp.StudentProfilePic,
				&Studenttmp.StudentRegisterDate)
			if err != nil {
				libErrors.ReturnError(conn)
			}

			StudentProfile = append(StudentProfile, globalObject.StudentProfile{
				Studenttmp.StudentId,
				Studenttmp.StudentFullName,
				Studenttmp.StudentBirthDate,
				Studenttmp.StudentParentFullName,
				Studenttmp.StudentPhoneNumber,
				Studenttmp.StudentParentPhoneNumber,
				Studenttmp.StudentLocation,
				Studenttmp.StudentCurrentLocation,
				Studenttmp.StudentClass,
				Studenttmp.StudentStudyGroupId,
				Studenttmp.StudentIdBack,
				Studenttmp.StudentIdFront,
				Studenttmp.StudentDeviceType,
				Studenttmp.StudentDeviceOsNum,
				Studenttmp.StudentActive,
				Studenttmp.StudentKey,
				Studenttmp.StudentOverAllNum,
				Studenttmp.SchoolId,
				Studenttmp.DisId,
				Studenttmp.ManagerId,
				Studenttmp.BusId,
				Studenttmp.StudentProfilePic,
				Studenttmp.StudentRegisterDate})

		}
	}
	return StudentProfile
}
