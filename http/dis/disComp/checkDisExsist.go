package dishttpcomp

import (
	"encoding/json"
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	diswscomp "iqdev/ss/websocket/dis/disWSComp"
	"net/http"
)

type checkDisStruct struct {
	DisKey string `json:"disKey"`
}

func CheckDisExsit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("checking dis")
	postDecoder := json.NewDecoder(r.Body)
	var decodeAddSOValue checkDisStruct

	decodeError := postDecoder.Decode(&decodeAddSOValue)
	if decodeError != nil {
		libErrors.ReturnHttpError(w)
	}

	fmt.Println(decodeAddSOValue.DisKey)

	disProfile := diswscomp.GetDisInfo(decodeAddSOValue.DisKey, nil)

	if disProfile.DisExsist == false {
		libErrors.ReturnHttpError(w)
	}
}
