package schoolwscomp

import (
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllTeacherStudyGroupsInfo(
	schoolId int,
	conn *websocket.Conn) []globalObject.TeacherStudyGroup {

	teacherGroup := []globalObject.TeacherStudyGroup{}
	teacherGroupTmp := globalObject.TeacherStudyGroup{}

	// sql connection

	// query
	getManagerQuery := "SELECT * FROM teacher_study_group where school_id = ?;"
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
			&teacherGroupTmp.TeacherStudyGroupId,
			&teacherGroupTmp.TeacherId,
			&teacherGroupTmp.StudyGroupId,
			&teacherGroupTmp.SchoolId)
		if err != nil {
			libErrors.ReturnError(conn)
		}

		teacherGroup = append(teacherGroup, globalObject.TeacherStudyGroup{
			teacherGroupTmp.TeacherStudyGroupId,
			teacherGroupTmp.TeacherId,
			teacherGroupTmp.StudyGroupId,
			teacherGroupTmp.SchoolId})

	}
	return teacherGroup
}
