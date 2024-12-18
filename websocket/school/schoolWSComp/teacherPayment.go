package schoolwscomp

import (
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllTeacherPaymentsInfo(
	schoolId int,
	conn *websocket.Conn) []globalObject.TeacherPaymentProfile {

	teacherPayment := []globalObject.TeacherPaymentProfile{}
	teacherPaymentTmp := globalObject.TeacherPaymentProfile{}

	// sql connection

	// query
	getManagerQuery := "SELECT * FROM teacher_payment where school_id = ?;"
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
			&teacherPaymentTmp.TeacherPaymentId,
			&teacherPaymentTmp.TeacherPaymentTitle,
			&teacherPaymentTmp.TeacherPaymentDesc,
			&teacherPaymentTmp.TeacherPaymentAmmount,
			&teacherPaymentTmp.TeacherPaymentDate,
			&teacherPaymentTmp.TeacherRegisterDate,
			&teacherPaymentTmp.SchoolId,
			&teacherPaymentTmp.TeacherId,
			&teacherPaymentTmp.TeacherName)
		if err != nil {
			libErrors.ReturnError(conn)
		}

		teacherPayment = append(teacherPayment, globalObject.TeacherPaymentProfile{
			teacherPaymentTmp.TeacherPaymentId,
			teacherPaymentTmp.TeacherPaymentTitle,
			teacherPaymentTmp.TeacherPaymentDesc,
			teacherPaymentTmp.TeacherPaymentAmmount,
			teacherPaymentTmp.TeacherPaymentDate,
			teacherPaymentTmp.TeacherRegisterDate,
			teacherPaymentTmp.SchoolId,
			teacherPaymentTmp.TeacherId,
			teacherPaymentTmp.TeacherName})

	}
	return teacherPayment
}
