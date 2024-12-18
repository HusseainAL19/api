package httpgod

import (
	"encoding/json"
	libErrors "iqdev/ss/libs/errors"
	genKey "iqdev/ss/libs/key"
	"iqdev/ss/libs/sql"
	godwscomp "iqdev/ss/websocket/god/godWSComp"
	"net/http"
	"time"
)

type addManagerPostStruct struct {
	GodKey                    string `json:"godKey"`
	ManagerName               string `json:"manager_name"`
	ManagerBirthDay           string `json:"manager_birth_day"`
	ManagerCurrentLocation    string `json:"manager_current_location"`
	ManagerImagePath          string `json:"manager_image_path"`
	ManagerDeviceLocation     string `json:"manager_device_location"`
	ManagerDeviceType         string `json:"manager_device_type"`
	ManagerDeviceXNumber      string `json:"manager_device_xnumber"`
	ManagerDeviceEmulated     bool   `json:"manager_device_emulated"`
	ManagerDeviceBattaryLevel string `json:"manager_device_battary_level"`
	ManagerTotalMemory        string `json:"manager_total_memory"`
	ManagerUsedMemory         string `json:"manager_used_memory"`
	ManagerDeviceCapacity     string `json:"manager_device_capacity"`
	ManagerDeviceFreeDisk     string `json:"manager_device_free_disk"`
	ManagerDeviceTotalImages  string `json:"manager_total_images"`
	ManagerDeviceTotalVideos  string `json:"manager_total_videos"`
	ManagerRegisterDate       string `json:"manager_register_date"`
	ManagerActive             bool   `json:"manager_active"`
	ManagerLastActivity       string `json:"manager_last_activity"`
	ManagerPhoneNumber        string `json:"manager_phone_number"`
}

func HttpGodHanlder() {
	http.HandleFunc("POST /http/addManager", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var godKeyStruct addManagerPostStruct
		err := decoder.Decode(&godKeyStruct)
		if err != nil {
			// error
			libErrors.ReturnHttpError(w)
		}
		// checking god value
		godValue := godwscomp.GetGodInfo(godKeyStruct.GodKey)
		if godValue.GodExsist == false {
		}
		// getting god info
		addManagerQuery := `INSERT INTO manager (manager_name,
    manager_birth_date,
    manager_current_location,
    manager_image_path,
    manager_device_location,
    manager_device_type,
    manager_device_xnumber,
    manager_device_emulated,
    manager_device_battary_level,
    manager_total_memory,
    manager_used_memory,
    manager_key,
    manager_device_capacity,
    manager_device_free_disk,
    manager_total_images,
    manager_total_videos,
    manager_register_date,
    manager_active,
    god_id,
    manager_last_activity,
    manager_phone_number) VALUE (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?);`

		sqlConn := sql.InitConnection().Connections
		_, sqlErr := sqlConn.Exec(addManagerQuery,
			godKeyStruct.ManagerName,
			godKeyStruct.ManagerBirthDay,
			godKeyStruct.ManagerCurrentLocation,
			godKeyStruct.ManagerImagePath,
			godKeyStruct.ManagerDeviceLocation,
			godKeyStruct.ManagerDeviceType,
			godKeyStruct.ManagerDeviceXNumber,
			godKeyStruct.ManagerDeviceEmulated,
			godKeyStruct.ManagerDeviceBattaryLevel,
			godKeyStruct.ManagerTotalMemory,
			godKeyStruct.ManagerUsedMemory,
			genKey.RandomKey(120),
			godKeyStruct.ManagerDeviceCapacity,
			godKeyStruct.ManagerDeviceFreeDisk,
			godKeyStruct.ManagerDeviceTotalImages,
			godKeyStruct.ManagerDeviceTotalVideos,
			time.Now(),
			true,
			godValue.GodInfo.GodId,
			time.Now(),
			godKeyStruct.ManagerPhoneNumber)

		if sqlErr != nil {
			w.Write([]byte("something went wrong"))
		}

		w.Header().Set("Content-Type", "application/json")
		newR, _ := json.Marshal(godKeyStruct)
		w.Write(newR)

		// store querys
	})
}
