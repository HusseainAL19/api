package managerwscomp

import (
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"
)

type ManagerInfoStruct struct {
	ManagerExsist bool
	ManagerInfo   globalObject.ManagersProfile
}

func GetManagerProfile(
	managerKey string,
	conn any) ManagerInfoStruct {

	managerProfiles := ManagerInfoStruct{}

	// sql connection

	// query
	getManagerQuery := "SELECT * FROM manager WHERE manager_key = ?;"
	// sql connection
	connObj := sql.InitConnection()
	// query manager info
	sqlError := connObj.Connections.QueryRow(getManagerQuery, managerKey).Scan(
		&managerProfiles.ManagerInfo.ManagerId,
		&managerProfiles.ManagerInfo.ManagerName,
		&managerProfiles.ManagerInfo.ManagerBirthDay,
		&managerProfiles.ManagerInfo.ManagerCurrentLocation,
		&managerProfiles.ManagerInfo.ManagerImagePath,
		&managerProfiles.ManagerInfo.ManagerDeviceLocation,
		&managerProfiles.ManagerInfo.ManagerDeviceType,
		&managerProfiles.ManagerInfo.ManagerDeviceXNumber,
		&managerProfiles.ManagerInfo.ManagerDeviceEmulated,
		&managerProfiles.ManagerInfo.ManagerDeviceBattaryLevel,
		&managerProfiles.ManagerInfo.ManagerTotalMemory,
		&managerProfiles.ManagerInfo.ManagerUsedMemory,
		&managerProfiles.ManagerInfo.ManagerKey,
		&managerProfiles.ManagerInfo.ManagerDeviceCapacity,
		&managerProfiles.ManagerInfo.ManagerDeviceFreeDisk,
		&managerProfiles.ManagerInfo.ManagerDeviceTotalImages,
		&managerProfiles.ManagerInfo.ManagerDeviceTotalVideos,
		&managerProfiles.ManagerInfo.ManagerRegisterDate,
		&managerProfiles.ManagerInfo.ManagerActive,
		&managerProfiles.ManagerInfo.GodId,
		&managerProfiles.ManagerInfo.ManagerLastActivity,
		&managerProfiles.ManagerInfo.ManagerPhoneNumber)

	//defer connObj.Connections.Close()

	if sqlError != nil {
		managerProfiles.ManagerExsist = false
		return managerProfiles
	}

	managerProfiles.ManagerExsist = true
	return managerProfiles
}
