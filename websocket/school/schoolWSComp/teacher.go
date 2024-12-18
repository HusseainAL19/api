package schoolwscomp

import (
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllTeacherInfo(
	schoolId int,
	conn *websocket.Conn) []globalObject.TeacherProfile {

	teacherProfile := []globalObject.TeacherProfile{}
	teachertmp := globalObject.TeacherProfile{}

	// sql connection

	// query
	getManagerQuery := "SELECT * FROM teacher where school_id = ?;"
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
			&teachertmp.TeacherId,
			&teachertmp.TeacherFullName,
			&teachertmp.TeacherBirthDay,
			&teachertmp.TeacherLocation,
			&teachertmp.TeacherDeviceLocation,
			&teachertmp.TeacherTotalStudent,
			&teachertmp.TeacherTotalScore,
			&teachertmp.TeacherTotalPresent,
			&teachertmp.TeacherImagePath,
			&teachertmp.TeacherIdXNumber,
			&teachertmp.TeacherIdImagePathFront,
			&teachertmp.TeacherIdImagePathBack,
			&teachertmp.TeacherDegree,
			&teachertmp.TeacherMajor,
			&teachertmp.TeacherKey,
			&teachertmp.TeacherTotalGroups,
			&teachertmp.SchoolId,
			&teachertmp.OwnerId,
			&teachertmp.TeacherActive)
		if err != nil {
			libErrors.ReturnError(conn)
		}

		teacherProfile = append(teacherProfile, globalObject.TeacherProfile{
			teachertmp.TeacherId,
			teachertmp.TeacherFullName,
			teachertmp.TeacherBirthDay,
			teachertmp.TeacherLocation,
			teachertmp.TeacherDeviceLocation,
			teachertmp.TeacherTotalStudent,
			teachertmp.TeacherTotalScore,
			teachertmp.TeacherTotalPresent,
			teachertmp.TeacherImagePath,
			teachertmp.TeacherIdXNumber,
			teachertmp.TeacherIdImagePathFront,
			teachertmp.TeacherIdImagePathBack,
			teachertmp.TeacherDegree,
			teachertmp.TeacherMajor,
			teachertmp.TeacherKey,
			teachertmp.TeacherTotalGroups,
			teachertmp.SchoolId,
			teachertmp.OwnerId,
			teachertmp.TeacherActive})

	}
	return teacherProfile
}
