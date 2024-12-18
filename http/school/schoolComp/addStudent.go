package schoolhttpcomp

import (
	"fmt"
	"io"
	genKey "iqdev/ss/libs/key"
	"iqdev/ss/libs/sql"
	schoolwscomp "iqdev/ss/websocket/school/schoolWSComp"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type addStudentPostStruct struct {
	StudentFullName          string `json:"student_full_name"`
	StudentBirthDate         string `json:"student_birth_date"`
	StudentParentFullName    string `json:"student_parent_full_name"`
	StudentPhoneNumber       string `json:"student_phone_number"`
	StudentParentPhoneNumber string `json:"student_parent_phone_number"`
	StudentLocation          string `json:"student_location"`
	StudentClass             string `json:"student_class"`
	StudentStudyGroupId      int    `json:"student_study_group_id"`
	StudentIdBack            string `json:"student_id_back"`
	StudentIdFront           string `json:"student_id_front"`
	SchoolKey                string `json:"school_key"`
	BusId                    int    `json:"bus_id"`
	StudentProfilePic        string `json:"student_profile_pic"`
}

func AddStudent() {
	http.HandleFunc("POST /http/school/addStudent", func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Expect") == "100-continue" {
			w.WriteHeader(http.StatusContinue)
		}
		r.ParseMultipartForm(100000000)

		studentProfilePic, header, fileError := r.FormFile("student_profile_pic")
		if fileError != nil {
			http.Error(w, "faild to get upload file ", http.StatusInternalServerError)
			return
		}
		defer studentProfilePic.Close()

		destenation := "./images/students/"
		fileInfoPath := genKey.RandomKey(20) + header.Filename

		cfile, cfileErr := os.OpenFile(
			fileInfoPath,
			os.O_WRONLY|os.O_CREATE,
			0666,
		)
		if cfileErr != nil {
			http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
			return
		}
		_, cpyErr := io.Copy(cfile, studentProfilePic)
		if cpyErr != nil {
			http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
			return
		}
		size, cpyErr := io.Copy(cfile, studentProfilePic)
		fmt.Println("done copy")
		if cpyErr != nil {
			http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
			return
		}
		fmt.Println(size)

		cmd := exec.Command("mv", "./"+fileInfoPath, destenation)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(stdout))

		studentFullName := r.FormValue("student_full_name")
		studentBirthDate := r.FormValue("student_birth_date")
		studentParentFullName := r.FormValue("student_parent_full_name")
		studentPhoneNumber := r.FormValue("student_phone_number")
		studentParentPhoneNumber := r.FormValue("student_parent_phone_number")
		studentLocation := r.FormValue("student_location")
		studentClass := r.FormValue("student_class")
		studentStudyGroupId := r.FormValue("student_study_group_id")
		schoolKey := r.FormValue("school_key")
		busId := r.FormValue("bus_id")

		schoolInfo := schoolwscomp.GetSchoolInfo(schoolKey, nil)

		if schoolInfo.SchoolExsist == false {
			http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
			return
		}
		if schoolInfo.SchoolProfile.SchoolActive == false {
			http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
			return
		}
		if fileError != nil {
			http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
			return
		}
		defer studentProfilePic.Close()

		addStudentSqlQuery := `
    student_full_name,
    student_birth_date,
    student_parent_full_name,
    student_phone_number,
    student_parent_phone_number,
    student_location,
    student_class,
    student_study_group_id,
    student_actve,
    student_key,
    student_overall_num,
    school_id,
    dis_id,
    manager_id,
    bus_id,
    student_profile_pic,
    student_register_date,
    student_money) VALUE (
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
		sqlConn := sql.InitConnection().Connections

		_, sqlErr := sqlConn.Exec(addStudentSqlQuery,
			studentFullName,
			studentBirthDate,
			studentParentFullName,
			studentPhoneNumber,
			studentParentPhoneNumber,
			studentLocation,
			studentClass,
			studentStudyGroupId,
			true,
			genKey.RandomKey(120),
			0,
			schoolInfo.SchoolProfile.SchoolId,
			schoolInfo.SchoolProfile.DisId,
			schoolInfo.SchoolProfile.ManagerId,
			busId,
			destenation+fileInfoPath,
			time.Now(),
			0)
		if sqlErr != nil {
			http.Error(w, "faild to insert dis into database", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("تم اضافه الطالب"))
	})
}
