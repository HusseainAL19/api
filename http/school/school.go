package httpSchool

import (
	schoolhttpcomp "iqdev/ss/http/school/schoolComp"
	"net/http"
)

func HttpSchoolHanlder() {
	http.HandleFunc("/http/school/schoolCheck", schoolhttpcomp.CheckSchoolExsit)
	schoolhttpcomp.AddChatGroup()
	schoolhttpcomp.AddExam()
	schoolhttpcomp.AddScheduleGroup()
	// student
	schoolhttpcomp.AddStudent()
	schoolhttpcomp.AddStudentGroup()
	// teacher
	http.HandleFunc("/http/school/addTeacher", schoolhttpcomp.AddTeacher)
	schoolhttpcomp.AddTeacherToGroup()
}
