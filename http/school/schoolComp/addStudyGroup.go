package schoolhttpcomp

import (
	"encoding/json"
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/sql"
	schoolwscomp "iqdev/ss/websocket/school/schoolWSComp"
	"net/http"
)

type addStudyGroupPostStruct struct {
	SchoolKey                string `json:"school_key"`
	StudentGroupName         string `json:"student_groups_name"`
	StudentGroupTotalStudent int    `json:"student_groups_total_student"`
	StudentGroupTeacherName  string `json:"student_groups_teacher_name"`
}

func AddStudentGroup() {
	http.HandleFunc(
		"POST /http/school/addStudentGroup",
		func(w http.ResponseWriter, r *http.Request) {
			postDecoder := json.NewDecoder(r.Body)
			var decodeAddSchoolValue addStudyGroupPostStruct

			decodeError := postDecoder.Decode(&decodeAddSchoolValue)
			if decodeError != nil {
				fmt.Println(decodeError)
				fmt.Println("decode error")
			}

			schoolProfile := schoolwscomp.GetSchoolInfo(decodeAddSchoolValue.SchoolKey, nil)
			if schoolProfile.SchoolExsist == false {
				libErrors.ReturnHttpError(w)
			}

			addSOQuery := `INSERT INTO study_groups(
        student_groups_name,
        student_groups_total_student,
        student_groups_teacher_name,
        school_id) VALUE (
          ?,
          ?,
          ?,
          ?);`

			sqlconnection := sql.InitConnection().Connections

			res, sqlError := sqlconnection.Exec(addSOQuery,
				decodeAddSchoolValue.StudentGroupName,
				decodeAddSchoolValue.StudentGroupTotalStudent,
				decodeAddSchoolValue.StudentGroupTeacherName,
				schoolProfile.SchoolProfile.SchoolId)

			fmt.Println(res)
			fmt.Println(sqlError)

			if sqlError != nil {
				fmt.Println("return sql error")
				libErrors.ReturnHttpError(w)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("تم اضافة مجموعة دراسية"))
		},
	)
}
