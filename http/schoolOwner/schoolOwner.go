package httpSo

import (
	sohttpcomp "iqdev/ss/http/schoolOwner/schoolOwnerComp"
	"net/http"
)

func HttpSOHanlder() {
	http.HandleFunc("/http/so/socheck", sohttpcomp.CheckSOExsit)
	http.HandleFunc("/http/so/addSchool", sohttpcomp.AddSchool)
}
