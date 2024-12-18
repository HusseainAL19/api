package schoolOwnerwscomp

import (
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllstudentPaymentPaymentsInfo(
	schoolsProfile []globalObject.SchoolsProfile,
	conn *websocket.Conn) []globalObject.StudentPaymentProfile {

	studentPaymentProfile := []globalObject.StudentPaymentProfile{}
	studentPaymenttmp := globalObject.StudentPaymentProfile{}

	// sql connection

	//getSchoolOwnerQuery := "SELECT * FROM school where  school_owner_id = ?;"
	getstudentPaymentQuery := "SELECT * FROM student where  school_id = ?;"

	// query
	for _, element := range schoolsProfile {
		// sql connection
		connObj := sql.InitConnection()
		// query manager info
		rows, sqlError := connObj.Connections.Query(getstudentPaymentQuery, element.SchoolId)

		if sqlError != nil {
			libErrors.ReturnError(conn)
		}

		defer rows.Close()
		defer connObj.Connections.Close()

		for rows.Next() {
			err := rows.Scan(
				&studentPaymenttmp.StudentPaymentId,
				&studentPaymenttmp.StudentPaymentRegisterPerson,
				&studentPaymenttmp.StudentPaymentRegisterTitle,
				&studentPaymenttmp.StudentPaymentRegisterDesc,
				&studentPaymenttmp.StudentPaymentRegisterRegisterDate,
				&studentPaymenttmp.StudentPaymentRegisterCurrentDate,
				&studentPaymenttmp.StudentPaymentTotalAmmount,
				&studentPaymenttmp.StudentPaymentDiscount,
				&studentPaymenttmp.SchoolId,
				&studentPaymenttmp.StudentId)
			if err != nil {
				libErrors.ReturnError(conn)
			}

			studentPaymentProfile = append(
				studentPaymentProfile,
				globalObject.StudentPaymentProfile{
					studentPaymenttmp.StudentPaymentId,
					studentPaymenttmp.StudentPaymentRegisterPerson,
					studentPaymenttmp.StudentPaymentRegisterTitle,
					studentPaymenttmp.StudentPaymentRegisterDesc,
					studentPaymenttmp.StudentPaymentRegisterRegisterDate,
					studentPaymenttmp.StudentPaymentRegisterCurrentDate,
					studentPaymenttmp.StudentPaymentTotalAmmount,
					studentPaymenttmp.StudentPaymentDiscount,
					studentPaymenttmp.SchoolId,
					studentPaymenttmp.StudentId},
			)

		}
	}
	return studentPaymentProfile
}
