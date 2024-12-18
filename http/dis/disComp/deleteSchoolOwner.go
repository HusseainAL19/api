package dishttpcomp

import (
	"encoding/json"
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/sql"
	diswscomp "iqdev/ss/websocket/dis/disWSComp"
	"net/http"
)

type DeleteSchoolOwnerPostStruct struct {
	DisKey        string `json:"disKey"`
	SchoolOwnerid int    `json:"soId"`
}

func DeleteSchoolOwner(w http.ResponseWriter, r *http.Request) {
	postDecoder := json.NewDecoder(r.Body)
	var decodeAddSOValue DeleteSchoolOwnerPostStruct
	decodeError := postDecoder.Decode(&decodeAddSOValue)

	fmt.Println(decodeError)

	if decodeError != nil {
		fmt.Println(decodeError)
	}

	disProfile := diswscomp.GetDisInfo(decodeAddSOValue.DisKey, nil)
	if disProfile.DisExsist == false {
		libErrors.ReturnHttpError(w)
	}

	fmt.Println(disProfile.DisProfile.DisId)
	fmt.Println(decodeAddSOValue.SchoolOwnerid)

	deleteSoQuery := `delete from school_owner where school_owner_id = ?;`

	sqlconnection := sql.InitConnection().Connections
	fmt.Println(decodeAddSOValue.SchoolOwnerid, "so id")
	res, sqlError := sqlconnection.Exec(deleteSoQuery, decodeAddSOValue.SchoolOwnerid)

	fmt.Println(res)
	fmt.Println(sqlError)

	if sqlError != nil {
		libErrors.ReturnHttpError(w)
		return
	}
	w.WriteHeader(http.StatusOK)
}
