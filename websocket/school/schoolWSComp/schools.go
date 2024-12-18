package schoolwscomp

import (
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"
)

type SchoolCheckStruct struct {
	SchoolExsist  bool
	SchoolProfile globalObject.SchoolsProfile
}

func GetSchoolInfo(
	schoolKey string,
	conn any) SchoolCheckStruct {

	schoolProfile := SchoolCheckStruct{}

	// sql connection

	// query
	getSchoolOwnerQuery := "SELECT * FROM school where school_key = ?;"
	// sql connection
	connObj := sql.InitConnection()
	// query manager info
	sqlError := connObj.Connections.QueryRow(getSchoolOwnerQuery, schoolKey).Scan(
		&schoolProfile.SchoolProfile.SchoolId,
		&schoolProfile.SchoolProfile.SchoolName,
		&schoolProfile.SchoolProfile.SchoolManagerName,
		&schoolProfile.SchoolProfile.SchoolOwnerName,
		&schoolProfile.SchoolProfile.SchoolManagerBirthDate,
		&schoolProfile.SchoolProfile.SchoolOwnerBirthDate,
		&schoolProfile.SchoolProfile.SchoolLocation,
		&schoolProfile.SchoolProfile.SchoolManagerLocation,
		&schoolProfile.SchoolProfile.SchoolManagerId,
		&schoolProfile.SchoolProfile.SchoolManagerCurrentLocation,
		&schoolProfile.SchoolProfile.SchoolTotalStudent,
		&schoolProfile.SchoolProfile.SchoolTotalTeachers,
		&schoolProfile.SchoolProfile.SchoolTotalBus,
		&schoolProfile.SchoolProfile.SchoolTotalAcc,
		&schoolProfile.SchoolProfile.SchoolOwnerId,
		&schoolProfile.SchoolProfile.ManagerId,
		&schoolProfile.SchoolProfile.DisId,
		&schoolProfile.SchoolProfile.SchoolKey,
		&schoolProfile.SchoolProfile.SchoolActive,
		&schoolProfile.SchoolProfile.SchoolStoreActive)

	if sqlError != nil {
		schoolProfile.SchoolExsist = false
    return schoolProfile;
	}

	defer connObj.Connections.Close()

	schoolProfile.SchoolExsist = true
	return schoolProfile
}
