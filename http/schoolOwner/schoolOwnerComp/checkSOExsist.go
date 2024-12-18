package sohttpcomp

import (
	"encoding/json"
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	schoolOwnerwscomp "iqdev/ss/websocket/schoolOwner/schoolOwnerWSComp"
	"net/http"
)

type checkDisStruct struct {
	SOKey string `json:"soKey"`
}

func CheckSOExsit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("checking school owner")
	postDecoder := json.NewDecoder(r.Body)
	var decodeSOValue checkDisStruct

	decodeError := postDecoder.Decode(&decodeSOValue)
	if decodeError != nil {
		libErrors.ReturnHttpError(w)
	}

	fmt.Println(decodeSOValue.SOKey)

	soProfile := schoolOwnerwscomp.GetSchoolOwnerInfo(decodeSOValue.SOKey, nil)

	fmt.Println(soProfile.SchoolOwnerExsist)

	if soProfile.SchoolOwnerExsist == false {
		libErrors.ReturnHttpError(w)
	}

	w.WriteHeader(http.StatusOK)
}
