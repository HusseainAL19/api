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
)

func AddTeacher(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Expect") == "100-continue" {
		w.WriteHeader(http.StatusContinue)
	}
	r.ParseMultipartForm(100000000)

	teacherProfilePic, header, fileError := r.FormFile("teacher_profile_pic")
	if fileError != nil {
		fmt.Println(fileError)
		fmt.Println("faild to get upload file ")
		http.Error(w, "faild to get upload file ", http.StatusInternalServerError)
		return
	}
	defer teacherProfilePic.Close()

	destenation := "./images/teachers/"
	fileInfoPath := genKey.RandomKey(20) + header.Filename

	cfile, cfileErr := os.OpenFile(
		fileInfoPath,
		os.O_WRONLY|os.O_CREATE,
		0666,
	)
	if cfileErr != nil {
		fmt.Println(cfileErr)
		http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
		return
	}
	size, cpyErr := io.Copy(cfile, teacherProfilePic)
	fmt.Println("done copy")
	if cpyErr != nil {
		fmt.Println(cpyErr)
		http.Error(w, "faild to read upload file ", http.StatusUnsupportedMediaType)
		return
	}
	fmt.Println(size)

	cmd := exec.Command("mv", "./"+fileInfoPath, destenation)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(stdout))

	teacherFullName := r.FormValue("teacher_name")
	teacherBirthDate := r.FormValue("teacher_birth_day")
	teacherLocation := r.FormValue("teacher_location")
	teacherPhoneNumber := r.FormValue("teacher_phone_number")
	teacherDegree := r.FormValue("teacher_degree")
	teacherMajor := r.FormValue("teacher_major")
	schoolKey := r.FormValue("school_key")

	schoolInfo := schoolwscomp.GetSchoolInfo(schoolKey, nil)

	if schoolInfo.SchoolExsist == false {
		http.Error(w, "school does not exsist", http.StatusUnauthorized)
		return
	}
	if schoolInfo.SchoolProfile.SchoolActive == false {
		http.Error(w, "faild to read upload file ", http.StatusLocked)
		return
	}
	defer teacherProfilePic.Close()

	addteacherSqlQuery := `INSERT INTO teacher( 
      teacher_full_name,
      teacher_birth_day,
      teacher_location,
      teacher_image_path,
      teacher_id_xnumber,
      teacher_degree,
      teacher_major,
      teacher_key,
      teacher_total_groups,
      teacher_total_score,
      teacher_total_student,
      school_id,
      owner_id,
      teacher_active) VALUE (
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

	_, sqlErr := sqlConn.Exec(addteacherSqlQuery,
		teacherFullName,
		teacherBirthDate,
		teacherLocation,
		destenation+fileInfoPath,
		teacherPhoneNumber,
		teacherDegree,
		teacherMajor,
		genKey.RandomKey(120),
		0,
		0,
		0,
		schoolInfo.SchoolProfile.SchoolId,
		schoolInfo.SchoolProfile.SchoolOwnerId,
		true)
	if sqlErr != nil {
		fmt.Println("sql shit")
		fmt.Println(sqlErr)
		http.Error(w, "faild to insert dis into database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("تم اضافة مدرس"))
}
