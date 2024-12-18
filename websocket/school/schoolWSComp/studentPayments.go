package schoolwscomp

import (
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllStudentPaymentsInfo(
	schoolId int,
	conn *websocket.Conn) []globalObject.StudentPaymentProfile {

	studentPayment := []globalObject.StudentPaymentProfile{}
	studentPaymentTmp := globalObject.StudentPaymentProfile{}

	// sql connection

	// query
	getManagerQuery := "SELECT * FROM student_payments where school_id = ?;"
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
			&studentPaymentTmp.StudentPaymentId,
			&studentPaymentTmp.StudentPaymentRegisterPerson,
			&studentPaymentTmp.StudentPaymentRegisterTitle,
			&studentPaymentTmp.StudentPaymentRegisterDesc,
			&studentPaymentTmp.StudentPaymentRegisterCurrentDate,
			&studentPaymentTmp.StudentPaymentRegisterCurrentDate,
			&studentPaymentTmp.StudentPaymentTotalAmmount,
			&studentPaymentTmp.StudentPaymentDiscount,
			&studentPaymentTmp.SchoolId,
			&studentPaymentTmp.StudentId)
		if err != nil {
      fmt.Println("bus error")
			libErrors.ReturnError(conn)
		}

		studentPayment = append(studentPayment, globalObject.StudentPaymentProfile{
			studentPaymentTmp.StudentPaymentId,
			studentPaymentTmp.StudentPaymentRegisterPerson,
			studentPaymentTmp.StudentPaymentRegisterTitle,
			studentPaymentTmp.StudentPaymentRegisterDesc,
			studentPaymentTmp.StudentPaymentRegisterCurrentDate,
			studentPaymentTmp.StudentPaymentRegisterCurrentDate,
			studentPaymentTmp.StudentPaymentTotalAmmount,
			studentPaymentTmp.StudentPaymentDiscount,
			studentPaymentTmp.SchoolId,
			studentPaymentTmp.StudentId})

	}
	return studentPayment
}
