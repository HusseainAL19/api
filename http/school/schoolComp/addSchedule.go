package schoolhttpcomp

import (
	"encoding/json"
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/sql"
	schoolwscomp "iqdev/ss/websocket/school/schoolWSComp"
	"net/http"
)

type addSchedulePostStruct struct {
	SchoolKey           string `json:"school_key"`
	StudentStudyGroupId int    `json:"study_group_id"`
	Day                 int    `json:"day"`
	First               int    `json:"first_lesston"`
	Second              int    `json:"second_lesston"`
	Thrid               int    `json:"thrid_lesston"`
	Fourth              int    `json:"fourth_lesston"`
	Fifith              int    `json:"fifth_lesston"`
	Sixth               int    `json:"sixth_lesston"`
	Seven               int    `json:"seven_lesston"`
	SchoolId            int    `json:"school_id"`
}

func AddScheduleGroup() {
	http.HandleFunc("POST /http/school/addSchedule", func(w http.ResponseWriter, r *http.Request) {
		postDecoder := json.NewDecoder(r.Body)
		var decodeAddSchoolValue addSchedulePostStruct

		decodeError := postDecoder.Decode(&decodeAddSchoolValue)
		if decodeError != nil {
			fmt.Println(decodeError)
			fmt.Println("decode error")
		}

		schoolProfile := schoolwscomp.GetSchoolInfo(decodeAddSchoolValue.SchoolKey, nil)
		if schoolProfile.SchoolExsist == false {
			libErrors.ReturnHttpError(w)
		}

		addSOQuery := `INSERT INTO schedule(
      study_group_id,
      day,
      first_lesston,
      second_lesston,
      thrid_lesston,
      fourth_lesston,
      fifth_lesston,
      sixth_lesston,
      seven_lesston,
      school_id) VALUE (
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

		sqlconnection := sql.InitConnection().Connections

		res, sqlError := sqlconnection.Exec(addSOQuery,
			decodeAddSchoolValue.StudentStudyGroupId,
			decodeAddSchoolValue.Day,
			decodeAddSchoolValue.First,
			decodeAddSchoolValue.Second,
			decodeAddSchoolValue.Thrid,
			decodeAddSchoolValue.Fourth,
			decodeAddSchoolValue.Fifith,
			decodeAddSchoolValue.Sixth,
			decodeAddSchoolValue.Seven,
			decodeAddSchoolValue.SchoolId)

		fmt.Println(res)
		fmt.Println(sqlError)

		if sqlError != nil {
			fmt.Println("return sql error")
			libErrors.ReturnHttpError(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("تم اضافة الجدول"))
	})
}
