package schoolOwnerwscomp

import (
	"fmt"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"
)

type schoolOwnerStruct struct {
	SchoolOwnerExsist bool
	SchoolOwnerInfo   globalObject.SchoolOwnerProfile
}

func GetSchoolOwnerInfo(
	schoolOwnerKey string,
	conn any) schoolOwnerStruct {

	schoolOwnerProfile := schoolOwnerStruct{}

	// sql connection

	// query
	getSchoolOwnerQuery := "SELECT * FROM school_owner where school_owner_key = ?;"
	// sql connection
	connObj := sql.InitConnection()

	defer connObj.Connections.Close()
	// query manager info
	sqlError := connObj.Connections.QueryRow(getSchoolOwnerQuery, schoolOwnerKey).Scan(
		&schoolOwnerProfile.SchoolOwnerInfo.SchoolOwnerId,
		&schoolOwnerProfile.SchoolOwnerInfo.SchoolOwnerName,
		&schoolOwnerProfile.SchoolOwnerInfo.SchoolOwnerLocation,
		&schoolOwnerProfile.SchoolOwnerInfo.SchoolOwnerTotalSchools,
		&schoolOwnerProfile.SchoolOwnerInfo.SchoolOwnerKey,
		&schoolOwnerProfile.SchoolOwnerInfo.SchoolOwnerBirthDate,
		&schoolOwnerProfile.SchoolOwnerInfo.DisId,
		&schoolOwnerProfile.SchoolOwnerInfo.ManagerId,
		&schoolOwnerProfile.SchoolOwnerInfo.SchoolOwnerACtive,
		&schoolOwnerProfile.SchoolOwnerInfo.RegisterDate,
		&schoolOwnerProfile.SchoolOwnerInfo.SchoolOwnerPhoneNumber,
		&schoolOwnerProfile.SchoolOwnerInfo.SchoolOwnerPersonalId,
	)

	if sqlError != nil {
    fmt.Println(schoolOwnerKey)
    fmt.Println(sqlError)
		schoolOwnerProfile.SchoolOwnerExsist = false
		return schoolOwnerProfile
	}

	schoolOwnerProfile.SchoolOwnerExsist = true

	return schoolOwnerProfile
}
