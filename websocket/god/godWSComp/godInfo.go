package godwscomp

import (
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"
)

type godInfoStruct struct {
	GodExsist bool
	GodInfo   globalObject.GodProfile
}

func GetGodInfo(god_key string) godInfoStruct {
	godInfo := godInfoStruct{}
	checkGodQ := `SELECT * FROM god WHERE god_key = ?`

	if len(god_key) < 1 {
		godInfo.GodExsist = false
		return godInfo
	}

	sqlConnObj := sql.InitConnection()
	defer sqlConnObj.Connections.Close()

	err := sqlConnObj.Connections.QueryRow(checkGodQ, god_key).Scan(
		&godInfo.GodInfo.GodId,
		&godInfo.GodInfo.GodName,
		&godInfo.GodInfo.GodKey,
		&godInfo.GodInfo.GodTotalSchools,
		&godInfo.GodInfo.GodTotalStudents,
		&godInfo.GodInfo.GodTotalTeachers,
		&godInfo.GodInfo.GodTotalBus,
		&godInfo.GodInfo.GodLastActive)

	//defer sqlConnObj.Connections.Close()

	if err != nil {
		godInfo.GodExsist = false
		return godInfo
	}

	godInfo.GodExsist = true

	return godInfo
}
