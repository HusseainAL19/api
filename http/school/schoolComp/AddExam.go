package schoolhttpcomp

import (
	"fmt"
	"io"
	genKey "iqdev/ss/libs/key"
	"iqdev/ss/libs/sql"
	schoolwscomp "iqdev/ss/websocket/school/schoolWSComp"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func returnExamPath(file multipart.File,
	fileHeader *multipart.FileHeader,
	fileError error, w http.ResponseWriter) bool {

	if fileError != nil {
		http.Error(w, "faild to get upload file ", http.StatusInternalServerError)
	}
	defer file.Close()

	destenation := "./images/teacher/"
	fileInfoPath := genKey.RandomKey(20) + fileHeader.Filename

	cfile, cfileErr := os.OpenFile(
		fileInfoPath,
		os.O_WRONLY|os.O_CREATE,
		0666,
	)
	if cfileErr != nil {
		http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
	}
	_, cpyErr := io.Copy(cfile, file)
	if cpyErr != nil {
		http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
	}
	size, cpyErr := io.Copy(cfile, file)
	fmt.Println("done copy")
	if cpyErr != nil {
		http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
	}
	fmt.Println(size)

	cmd := exec.Command("mv", "./"+fileInfoPath, destenation)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(stdout))
	return true
}

func AddExam() {
	http.HandleFunc("POST /http/school/addExam", func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Expect") == "100-continue" {
			w.WriteHeader(http.StatusContinue)
		}
		r.ParseMultipartForm(100000000)

		//studentProfilePic, header, fileError := r.FormFile("exam_")

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
