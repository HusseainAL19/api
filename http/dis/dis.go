package httpdis

import (
	dishttpcomp "iqdev/ss/http/dis/disComp"
	"net/http"
)

func HttpDisHanlder() {
	http.HandleFunc("/http/dis/addSchoolOwner", dishttpcomp.AddSchoolOwner)
	http.HandleFunc("/http/dis/checkDis", dishttpcomp.CheckDisExsit)
	http.HandleFunc("/http/dis/delSo", dishttpcomp.DeleteSchoolOwner)
}
