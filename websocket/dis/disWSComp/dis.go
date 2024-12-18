package diswscomp

import (
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"
)

type RDisInfo struct {
	DisExsist  bool
	DisProfile globalObject.DisProfile
}

func GetDisInfo(
	disKey string,
	conn any) RDisInfo {

	disProfile := RDisInfo{}

	// sql connection

	// query
	getdisQuery := "SELECT * FROM dis WHERE dis_key = ?;"
	// sql connection
	connObj := sql.InitConnection()
	// query manager info
	sqlError := connObj.Connections.QueryRow(getdisQuery, disKey).Scan(
		&disProfile.DisProfile.DisId,
		&disProfile.DisProfile.DisFullName,
		&disProfile.DisProfile.DisBirthday,
		&disProfile.DisProfile.DisImagePath,
		&disProfile.DisProfile.DisCurrentLocation,
		&disProfile.DisProfile.DisHomeLocation,
		&disProfile.DisProfile.DisPhoneType,
		&disProfile.DisProfile.DisPhoneXnumber,
		&disProfile.DisProfile.DisPhoneIsEmulated,
		&disProfile.DisProfile.DisPhoneBattaryLevel,
		&disProfile.DisProfile.DisTotalMemory,
		&disProfile.DisProfile.DisUsedMemory,
		&disProfile.DisProfile.DisKey,
		&disProfile.DisProfile.DisPhoneCapacity,
		&disProfile.DisProfile.DisPhoneDiskFree,
		&disProfile.DisProfile.DisTotalImages,
		&disProfile.DisProfile.DisTotalVideos,
		&disProfile.DisProfile.ManagerId,
		&disProfile.DisProfile.DisRegisterDate,
		&disProfile.DisProfile.DisActive,
		&disProfile.DisProfile.DisIdNumber)

	if sqlError != nil {
		disProfile.DisExsist = false
	}

	defer connObj.Connections.Close()

	disProfile.DisExsist = true

	return disProfile
}
