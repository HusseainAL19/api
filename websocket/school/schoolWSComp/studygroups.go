package schoolwscomp

import (
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllStudyGroupsInfo(
	schoolId int,
	conn *websocket.Conn) []globalObject.StudyGroupProfile {

	studentGroup := []globalObject.StudyGroupProfile{}
	studentGroupTmp := globalObject.StudyGroupProfile{}

	// sql connection

	// query
	getManagerQuery := "SELECT * FROM study_groups where school_id = ?;"
	// sql connection
	connObj := sql.InitConnection()
	// query manager info
	rows, sqlError := connObj.Connections.Query(getManagerQuery, schoolId)

	if sqlError != nil {
		fmt.Println("bus error")
		libErrors.ReturnError(conn)
	}

	defer rows.Close()
	defer connObj.Connections.Close()

	for rows.Next() {
		err := rows.Scan(
			&studentGroupTmp.StudentGroupId,
			&studentGroupTmp.StudentGroupName,
			&studentGroupTmp.StudentGroupTotalStudent,
			&studentGroupTmp.StudentGroupTeacherName,
			&studentGroupTmp.SchoolId)
		if err != nil {
			fmt.Println("bus error")
			libErrors.ReturnError(conn)
		}

		studentGroup = append(studentGroup, globalObject.StudyGroupProfile{
			studentGroupTmp.StudentGroupId,
			studentGroupTmp.StudentGroupName,
			studentGroupTmp.StudentGroupTotalStudent,
			studentGroupTmp.StudentGroupTeacherName,
			studentGroupTmp.SchoolId})

	}
	return studentGroup
}
