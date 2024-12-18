package schoolhttpcomp

import (
	"encoding/json"
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	schoolwscomp "iqdev/ss/websocket/school/schoolWSComp"
	"net/http"
)

type checkDisStruct struct {
	SchoolKey string `json:"schoolKey"`
}

func CheckSchoolExsit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("checking school owner")
	postDecoder := json.NewDecoder(r.Body)
	var decodeSOValue checkDisStruct

	decodeError := postDecoder.Decode(&decodeSOValue)
	if decodeError != nil {
		libErrors.ReturnHttpError(w)
	}

	schoolProfile := schoolwscomp.GetSchoolInfo(decodeSOValue.SchoolKey, nil)

	fmt.Println(schoolProfile.SchoolExsist)
	if schoolProfile.SchoolExsist == false {
		libErrors.ReturnHttpError(w)
	}

}
