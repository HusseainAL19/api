package schoolOwnerwscomp

import (
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllSchoolInfo(
	schoolOwnerId int,
	conn *websocket.Conn) []globalObject.SchoolsProfile {

	schoolProfiles := []globalObject.SchoolsProfile{}
	schoolProfileTmp := globalObject.SchoolsProfile{}

	// sql connection

	// query
	getSchoolOwnerQuery := "SELECT * FROM school where  school_owner_id = ?;"
	// sql connection
	connObj := sql.InitConnection()
	// query manager info
	rows, sqlError := connObj.Connections.Query(getSchoolOwnerQuery, schoolOwnerId)

	if sqlError != nil {
		fmt.Println(sqlError)
		libErrors.ReturnError(conn)
	}

	defer rows.Close()
	defer connObj.Connections.Close()

	for rows.Next() {
		err := rows.Scan(
			&schoolProfileTmp.SchoolId,
			&schoolProfileTmp.SchoolName,
			&schoolProfileTmp.SchoolManagerName,
			&schoolProfileTmp.SchoolOwnerName,
			&schoolProfileTmp.SchoolManagerBirthDate,
			&schoolProfileTmp.SchoolOwnerBirthDate,
			&schoolProfileTmp.SchoolLocation,
			&schoolProfileTmp.SchoolManagerLocation,
			&schoolProfileTmp.SchoolManagerId,
			&schoolProfileTmp.SchoolManagerCurrentLocation,
			&schoolProfileTmp.SchoolTotalStudent,
			&schoolProfileTmp.SchoolTotalTeachers,
			&schoolProfileTmp.SchoolTotalBus,
			&schoolProfileTmp.SchoolTotalAcc,
			&schoolProfileTmp.SchoolOwnerId,
			&schoolProfileTmp.ManagerId,
			&schoolProfileTmp.DisId,
			&schoolProfileTmp.SchoolKey,
			&schoolProfileTmp.SchoolActive,
			&schoolProfileTmp.SchoolStoreActive)
		if err != nil {
			fmt.Println(err)
			libErrors.ReturnError(conn)
		}

		schoolProfiles = append(schoolProfiles, globalObject.SchoolsProfile{
			schoolProfileTmp.SchoolId,
			schoolProfileTmp.SchoolName,
			schoolProfileTmp.SchoolManagerName,
			schoolProfileTmp.SchoolOwnerName,
			schoolProfileTmp.SchoolManagerBirthDate,
			schoolProfileTmp.SchoolOwnerBirthDate,
			schoolProfileTmp.SchoolLocation,
			schoolProfileTmp.SchoolManagerLocation,
			schoolProfileTmp.SchoolManagerId,
			schoolProfileTmp.SchoolManagerCurrentLocation,
			schoolProfileTmp.SchoolTotalStudent,
			schoolProfileTmp.SchoolTotalTeachers,
			schoolProfileTmp.SchoolTotalBus,
			schoolProfileTmp.SchoolTotalAcc,
			schoolProfileTmp.SchoolOwnerId,
			schoolProfileTmp.ManagerId,
			schoolProfileTmp.DisId,
			schoolProfileTmp.SchoolKey,
			schoolProfileTmp.SchoolActive,
			schoolProfileTmp.SchoolStoreActive})
	}

	return schoolProfiles
}
