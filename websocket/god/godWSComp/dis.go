package godwscomp

import (
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllDisInfo(
	godInfo globalObject.GodProfile,
	conn *websocket.Conn) []globalObject.DisProfile {

	disProfiles := []globalObject.DisProfile{}
	disProfileTmp := globalObject.DisProfile{}

	// sql connection

	// query
	getdisQuery := "SELECT * FROM dis;"
	// sql connection
	connObj := sql.InitConnection()
	// query manager info
	rows, sqlError := connObj.Connections.Query(getdisQuery)

	if sqlError != nil {
		libErrors.ReturnError(conn)
	}

	defer rows.Close()
	defer connObj.Connections.Close()

	for rows.Next() {
		err := rows.Scan(
			&disProfileTmp.DisId,
			&disProfileTmp.DisFullName,
			&disProfileTmp.DisBirthday,
			&disProfileTmp.DisImagePath,
			&disProfileTmp.DisCurrentLocation,
			&disProfileTmp.DisHomeLocation,
			&disProfileTmp.DisPhoneType,
			&disProfileTmp.DisPhoneXnumber,
			&disProfileTmp.DisPhoneIsEmulated,
			&disProfileTmp.DisPhoneBattaryLevel,
			&disProfileTmp.DisTotalMemory,
			&disProfileTmp.DisUsedMemory,
			&disProfileTmp.DisKey,
			&disProfileTmp.DisPhoneCapacity,
			&disProfileTmp.DisPhoneDiskFree,
			&disProfileTmp.DisTotalImages,
			&disProfileTmp.DisTotalVideos,
			&disProfileTmp.ManagerId,
			&disProfileTmp.DisRegisterDate,
			&disProfileTmp.DisActive,
			&disProfileTmp.DisIdNumber)
		if err != nil {
			libErrors.ReturnError(conn)
		}

		disProfiles = append(disProfiles, globalObject.DisProfile{
			disProfileTmp.DisId,
			disProfileTmp.DisFullName,
			disProfileTmp.DisBirthday,
			disProfileTmp.DisImagePath,
			disProfileTmp.DisCurrentLocation,
			disProfileTmp.DisHomeLocation,
			disProfileTmp.DisPhoneType,
			disProfileTmp.DisPhoneXnumber,
			disProfileTmp.DisPhoneIsEmulated,
			disProfileTmp.DisPhoneBattaryLevel,
			disProfileTmp.DisTotalMemory,
			disProfileTmp.DisUsedMemory,
			disProfileTmp.DisKey,
			disProfileTmp.DisPhoneCapacity,
			disProfileTmp.DisPhoneDiskFree,
			disProfileTmp.DisTotalImages,
			disProfileTmp.DisTotalVideos,
			disProfileTmp.ManagerId,
			disProfileTmp.DisRegisterDate,
			disProfileTmp.DisActive,
			disProfileTmp.DisIdNumber})
	}

	return disProfiles
}
