package godwscomp

import (
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllManagerInfo(
	godInfo globalObject.GodProfile,
	conn *websocket.Conn  ) []globalObject.ManagersProfile {

	managerProfiles := []globalObject.ManagersProfile{}
	managerProfileTmp := globalObject.ManagersProfile{}

	// sql connection

	// query
	getManagerQuery := "SELECT * FROM manager;"
	// sql connection
	connObj := sql.InitConnection()
	// query manager info
	rows, sqlError := connObj.Connections.Query(getManagerQuery)

	if sqlError != nil {
		libErrors.ReturnError(conn)
	}

	defer rows.Close()
	defer connObj.Connections.Close()

	for rows.Next() {
		err := rows.Scan(
			&managerProfileTmp.ManagerId,
			&managerProfileTmp.ManagerName,
			&managerProfileTmp.ManagerBirthDay,
			&managerProfileTmp.ManagerCurrentLocation,
			&managerProfileTmp.ManagerImagePath,
			&managerProfileTmp.ManagerDeviceLocation,
			&managerProfileTmp.ManagerDeviceType,
			&managerProfileTmp.ManagerDeviceXNumber,
			&managerProfileTmp.ManagerDeviceEmulated,
			&managerProfileTmp.ManagerDeviceBattaryLevel,
			&managerProfileTmp.ManagerTotalMemory,
			&managerProfileTmp.ManagerUsedMemory,
			&managerProfileTmp.ManagerKey,
			&managerProfileTmp.ManagerDeviceCapacity,
			&managerProfileTmp.ManagerDeviceFreeDisk,
			&managerProfileTmp.ManagerDeviceTotalImages,
			&managerProfileTmp.ManagerDeviceTotalVideos,
			&managerProfileTmp.ManagerRegisterDate,
			&managerProfileTmp.ManagerActive,
			&managerProfileTmp.GodId,
			&managerProfileTmp.ManagerLastActivity,
			&managerProfileTmp.ManagerPhoneNumber)
		if err != nil {
			libErrors.ReturnError(conn)
		}

		managerProfiles = append(managerProfiles, globalObject.ManagersProfile{
			managerProfileTmp.ManagerId,
			managerProfileTmp.ManagerName,
			managerProfileTmp.ManagerBirthDay,
			managerProfileTmp.ManagerCurrentLocation,
			managerProfileTmp.ManagerImagePath,
			managerProfileTmp.ManagerDeviceLocation,
			managerProfileTmp.ManagerDeviceType,
			managerProfileTmp.ManagerDeviceXNumber,
			managerProfileTmp.ManagerDeviceEmulated,
			managerProfileTmp.ManagerDeviceBattaryLevel,
			managerProfileTmp.ManagerTotalMemory,
			managerProfileTmp.ManagerUsedMemory,
			managerProfileTmp.ManagerKey,
			managerProfileTmp.ManagerDeviceCapacity,
			managerProfileTmp.ManagerDeviceFreeDisk,
			managerProfileTmp.ManagerDeviceTotalImages,
			managerProfileTmp.ManagerDeviceTotalVideos,
			managerProfileTmp.ManagerRegisterDate,
			managerProfileTmp.ManagerActive,
			managerProfileTmp.GodId,
			managerProfileTmp.ManagerLastActivity,
			managerProfileTmp.ManagerPhoneNumber})
	}

	return managerProfiles
}
