package diswscomp

import (
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllSchoolOwnerInfo(
	disId int,
	conn *websocket.Conn) []globalObject.SchoolOwnerProfile {

	schoolOwnerProfiles := []globalObject.SchoolOwnerProfile{}
	schoolOwnerProfileTmp := globalObject.SchoolOwnerProfile{}

	// sql connection

	// query
	getSchoolOwnerQuery := "SELECT * FROM school_owner where dis_id = ?;"
	// sql connection
	connObj := sql.InitConnection()
	// query manager info
	rows, sqlError := connObj.Connections.Query(getSchoolOwnerQuery, disId)

	if sqlError != nil {
		libErrors.ReturnError(conn)
	}

	defer rows.Close()
	defer connObj.Connections.Close()

	for rows.Next() {
		err := rows.Scan(
			&schoolOwnerProfileTmp.SchoolOwnerId,
			&schoolOwnerProfileTmp.SchoolOwnerName,
			&schoolOwnerProfileTmp.SchoolOwnerLocation,
			&schoolOwnerProfileTmp.SchoolOwnerTotalSchools,
			&schoolOwnerProfileTmp.SchoolOwnerKey,
			&schoolOwnerProfileTmp.SchoolOwnerBirthDate,
			&schoolOwnerProfileTmp.DisId,
			&schoolOwnerProfileTmp.ManagerId,
			&schoolOwnerProfileTmp.SchoolOwnerACtive,
			&schoolOwnerProfileTmp.RegisterDate,
			&schoolOwnerProfileTmp.SchoolOwnerPhoneNumber,
			&schoolOwnerProfileTmp.SchoolOwnerPersonalId)
		if err != nil {
			fmt.Println("error during reading schoolowner from database")
			libErrors.ReturnError(conn)
		}

		schoolOwnerProfiles = append(schoolOwnerProfiles, globalObject.SchoolOwnerProfile{
			schoolOwnerProfileTmp.SchoolOwnerId,
			schoolOwnerProfileTmp.SchoolOwnerName,
			schoolOwnerProfileTmp.SchoolOwnerLocation,
			schoolOwnerProfileTmp.SchoolOwnerTotalSchools,
			schoolOwnerProfileTmp.SchoolOwnerKey,
			schoolOwnerProfileTmp.SchoolOwnerBirthDate,
			schoolOwnerProfileTmp.DisId,
			schoolOwnerProfileTmp.ManagerId,
			schoolOwnerProfileTmp.SchoolOwnerACtive,
			schoolOwnerProfileTmp.RegisterDate,
			schoolOwnerProfileTmp.SchoolOwnerPhoneNumber,
			schoolOwnerProfileTmp.SchoolOwnerPersonalId})
	}

	return schoolOwnerProfiles
}
