package sohttpcomp

import (
	"encoding/json"
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	_ "iqdev/ss/libs/key"
	genKey "iqdev/ss/libs/key"
	"iqdev/ss/libs/sql"
	schoolOwnerwscomp "iqdev/ss/websocket/schoolOwner/schoolOwnerWSComp"
	"net/http"
)

type addSchoolOwnerPostStruct struct {
	SoKey                  string `json:"soKey"`
	SchoolName             string `json:"school_name"`
	SchoolManagerName      string `json:"school_manager_name"`
	SchoolManagerBirthDate string `json:"school_manager_birth_date"`
	SchoolLocation         string `json:"school_location"`
	SchoolManagerLocation  string `json:"school_manager_location"`
	SchoolManagerId        string `json:"school_manager_id"`
	SchoolStoreActive      bool   `json:"store_active"`
}

func AddSchool(w http.ResponseWriter, r *http.Request) {
	postDecoder := json.NewDecoder(r.Body)
	var decodeAddSchoolValue addSchoolOwnerPostStruct

	decodeError := postDecoder.Decode(&decodeAddSchoolValue)
	if decodeError != nil {
		fmt.Println(decodeError)
		fmt.Println("decode error")
		w.WriteHeader(http.StatusUnauthorized)
	}

  fmt.Println(decodeAddSchoolValue)

	soInfo := schoolOwnerwscomp.GetSchoolOwnerInfo(decodeAddSchoolValue.SoKey, nil)
  fmt.Println(soInfo.SchoolOwnerInfo.SchoolOwnerId)
	if soInfo.SchoolOwnerExsist == false {
		w.WriteHeader(http.StatusUnauthorized)
	}

	addSOQuery := `INSERT INTO school (
    school_name,
    school_manager_name,
    school_owner_name,
    school_manager_birth_date,
    school_owner_birth_date,
    school_location,
    school_manager_location,
    school_manager_id,
    school_manager_current_location,
    school_total_student,
    school_total_teachers,
    school_total_bus,
    school_total_accountent,
    school_owner_id,
    manager_id,
    dis_id,
    school_key,
    school_active,
    school_store_active)
    VALUES (
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
    ?
  );`

	sqlconnection := sql.InitConnection().Connections

	res, sqlError := sqlconnection.Exec(addSOQuery,
		decodeAddSchoolValue.SchoolName,
		decodeAddSchoolValue.SchoolManagerName,
		soInfo.SchoolOwnerInfo.SchoolOwnerName,
		decodeAddSchoolValue.SchoolManagerBirthDate,
		soInfo.SchoolOwnerInfo.SchoolOwnerBirthDate,
		decodeAddSchoolValue.SchoolLocation,
		decodeAddSchoolValue.SchoolManagerLocation,
		decodeAddSchoolValue.SchoolManagerId,
		"",
		0,
		0,
		0,
		0,
		soInfo.SchoolOwnerInfo.SchoolOwnerId,
		soInfo.SchoolOwnerInfo.ManagerId,
		soInfo.SchoolOwnerInfo.DisId,
		genKey.RandomKey(120),
		true,
		decodeAddSchoolValue.SchoolStoreActive)

	fmt.Println(res)
	fmt.Println(sqlError)

	if sqlError != nil {
		fmt.Println("return sql error")
		libErrors.ReturnHttpError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
