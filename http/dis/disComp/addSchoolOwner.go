package dishttpcomp

import (
	"encoding/json"
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	genKey "iqdev/ss/libs/key"
	"iqdev/ss/libs/sql"
	diswscomp "iqdev/ss/websocket/dis/disWSComp"
	"net/http"
	"time"
)

type addSchoolOwnerPostStruct struct {
	DisKey                  string `json:"disKey"`
	SchoolOwnerName         string `json:"soName"`
	SchoolOwnerBirthDate    string `json:"soBD"`
	SchoolOwnerLocation     string `json:"soPlace"`
	SchoolOwnerACtive       bool   `json:"soActive"`
	SchoolOwnerTotalSchools int    `json:"schoolCount"`
	SchoolOwnerPhoneNumber  string `json:"phoneNumber"`
	SchoolOwnerIdCard       string `json:"idCard"`
}

func AddSchoolOwner(w http.ResponseWriter, r *http.Request) {
	postDecoder := json.NewDecoder(r.Body)
	var decodeAddSOValue addSchoolOwnerPostStruct

	decodeError := postDecoder.Decode(&decodeAddSOValue)
	if decodeError != nil {
    fmt.Println(decodeError)
    fmt.Println("decode error")
	}

	disProfile := diswscomp.GetDisInfo(decodeAddSOValue.DisKey, nil)

	addSOQuery := `INSERT INTO school_owner (
    school_owner_name,
    school_owner_location,
    school_owner_total_schools,
    school_owner_key,
    school_owner_birth_date,
    dis_id,
    manager_id,
    school_owner_active,
    school_owner_register_date,
    school_owner_phone_number,
    school_owner_personal_id
  ) VALUES (
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
    ?
  );`

	sqlconnection := sql.InitConnection().Connections

	res, sqlError := sqlconnection.Exec(addSOQuery,
		decodeAddSOValue.SchoolOwnerName,
		decodeAddSOValue.SchoolOwnerLocation,
		decodeAddSOValue.SchoolOwnerTotalSchools,
		genKey.RandomKey(120),
		decodeAddSOValue.SchoolOwnerBirthDate,
		disProfile.DisProfile.DisId,
		disProfile.DisProfile.ManagerId,
		decodeAddSOValue.SchoolOwnerACtive,
		time.Now(),
		decodeAddSOValue.SchoolOwnerPhoneNumber,
		decodeAddSOValue.SchoolOwnerIdCard)

    fmt.Println(res)
    fmt.Println(sqlError)

	if sqlError != nil {
		fmt.Println("return sql error")
		libErrors.ReturnHttpError(w)
		return
	}
}
