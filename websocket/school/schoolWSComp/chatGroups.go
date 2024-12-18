package schoolwscomp

import (
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllStudyChatGroupsInfo(
	schoolId int,
	conn *websocket.Conn) []globalObject.StudyChatGroupProfile {

	studentChatGroup := []globalObject.StudyChatGroupProfile{}
	studentChatGroupTmp := globalObject.StudyChatGroupProfile{}

	// sql connection

	// query
	getManagerQuery := "SELECT * FROM study_chat_group where school_id = ?;"
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
			&studentChatGroupTmp.StudyChatGroupId,
			&studentChatGroupTmp.StudyChatGroupName,
			&studentChatGroupTmp.StudyChatGroupTotalStudent,
			&studentChatGroupTmp.StudyGroupId,
			&studentChatGroupTmp.SchoolId,
			&studentChatGroupTmp.TeacherId,
			&studentChatGroupTmp.TeacherName)
		if err != nil {
      fmt.Println("bus error")
			libErrors.ReturnError(conn)
		}

		studentChatGroup = append(studentChatGroup, globalObject.StudyChatGroupProfile{
			studentChatGroupTmp.StudyChatGroupId,
			studentChatGroupTmp.StudyChatGroupName,
			studentChatGroupTmp.StudyChatGroupTotalStudent,
			studentChatGroupTmp.StudyGroupId,
			studentChatGroupTmp.SchoolId,
			studentChatGroupTmp.TeacherId,
			studentChatGroupTmp.TeacherName})

	}
	return studentChatGroup
}
