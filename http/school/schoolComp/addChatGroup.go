package schoolhttpcomp

import (
	"encoding/json"
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/sql"
	schoolwscomp "iqdev/ss/websocket/school/schoolWSComp"
	"net/http"
)

type addChatGroupPostStruct struct {
	SchoolKey          string `json:"school_key"`
	StudyChatGroupName string `json:"student_groups_name"`
	StudyGroupId       int    `json:"student_groups_id"`
	TeacherId          int    `json:"teacher_id"`
	TeacherName        int    `json:"teacher_name"`
}

func AddChatGroup() {
	http.HandleFunc("POST /http/school/addChatGroup", func(w http.ResponseWriter, r *http.Request) {
		postDecoder := json.NewDecoder(r.Body)
		var decodeAddSchoolValue addChatGroupPostStruct

		decodeError := postDecoder.Decode(&decodeAddSchoolValue)
		if decodeError != nil {
			fmt.Println(decodeError)
			fmt.Println("decode error")
		}

		schoolProfile := schoolwscomp.GetSchoolInfo(decodeAddSchoolValue.SchoolKey, nil)
		if schoolProfile.SchoolExsist == false {
			libErrors.ReturnHttpError(w)
		}

		addSOQuery := `INSERT INTO study_chat_group(
      student_groups_name,
      student_groups_total_student,
      student_groups_id,
      school_id,
      teacher_id,
      teacher_name) VALUES (
        ?,
        ?,
        ?,
        ?,
        ?,
        ?
      );`

		sqlconnection := sql.InitConnection().Connections

		res, sqlError := sqlconnection.Exec(addSOQuery,
			decodeAddSchoolValue.StudyChatGroupName,
			0,
			decodeAddSchoolValue.StudyGroupId,
			schoolProfile.SchoolProfile.SchoolId,
			decodeAddSchoolValue.TeacherId,
			decodeAddSchoolValue.TeacherName)

		fmt.Println(res)
		fmt.Println(sqlError)

		if sqlError != nil {
			fmt.Println("return sql error")
			libErrors.ReturnHttpError(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("تم اضافة مجموعة"))
	})
}
