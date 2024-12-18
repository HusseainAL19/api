package schoolwscomp

import (
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllStudentHistoryInfo(
	schoolId int,
	conn *websocket.Conn) []globalObject.StudentHistory {

	studentHistory := []globalObject.StudentHistory{}
	studentHistoryTmp := globalObject.StudentHistory{}

	// sql connection

	// query
	getManagerQuery := "SELECT * FROM student_history where school_id = ?;"
	// sql connection
	connObj := sql.InitConnection()
	// query manager info
	rows, sqlError := connObj.Connections.Query(getManagerQuery,schoolId)

	if sqlError != nil {
		libErrors.ReturnError(conn)
	}

	defer rows.Close()
	defer connObj.Connections.Close()

	for rows.Next() {
		err := rows.Scan(
			&studentHistoryTmp.StudentHistoryId,
			&studentHistoryTmp.StudentHistoryTitle,
			&studentHistoryTmp.StudentHistoryMaterial,
			&studentHistoryTmp.StudentHistoryDate,
			&studentHistoryTmp.StudentId,
			&studentHistoryTmp.SchoolId)
		if err != nil {
      fmt.Println("bus error")
			libErrors.ReturnError(conn)
		}

		studentHistory = append(studentHistory, globalObject.StudentHistory{
			studentHistoryTmp.StudentHistoryId,
			studentHistoryTmp.StudentHistoryTitle,
			studentHistoryTmp.StudentHistoryMaterial,
			studentHistoryTmp.StudentHistoryDate,
			studentHistoryTmp.StudentId,
			studentHistoryTmp.SchoolId})

	}
	return studentHistory
}
