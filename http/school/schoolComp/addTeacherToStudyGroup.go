package schoolhttpcomp

import (
	"encoding/json"
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/sql"
	schoolwscomp "iqdev/ss/websocket/school/schoolWSComp"
	"net/http"
)

type addTeacherToGroupPostStruct struct {
	SchoolKey    string `json:"school_key"`
	TeacherId    int    `json:"teacher_id"`
	StudyGroupId int    `json:"study_group_id"`
}

func AddTeacherToGroup() {
	http.HandleFunc(
		"POST /http/school/addTeachertoGroup",
		func(w http.ResponseWriter, r *http.Request) {
			postDecoder := json.NewDecoder(r.Body)
			var decodeAddSchoolValue addTeacherToGroupPostStruct

			decodeError := postDecoder.Decode(&decodeAddSchoolValue)
			if decodeError != nil {
				fmt.Println(decodeError)
				fmt.Println("decode error")
			}

			schoolProfile := schoolwscomp.GetSchoolInfo(decodeAddSchoolValue.SchoolKey, nil)
			if schoolProfile.SchoolExsist == false {
				libErrors.ReturnHttpError(w)
			}

			addSOQuery := `INSERT INTO teacher_study_group(
        teacher_id,
        study_group_id,
        school_id) VALUE (
          ?,
          ?,
          ?);`

			sqlconnection := sql.InitConnection().Connections

			res, sqlError := sqlconnection.Exec(addSOQuery,
				decodeAddSchoolValue.TeacherId,
				decodeAddSchoolValue.StudyGroupId,
				schoolProfile.SchoolProfile.SchoolId)

			fmt.Println(res)
			fmt.Println(sqlError)

			if sqlError != nil {
				fmt.Println("return sql error")
				libErrors.ReturnHttpError(w)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("تم اضافة المدرس الى المجموعة"))
		},
	)
}
