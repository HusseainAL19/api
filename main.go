package main

import (
	"fmt"
	httpdis "iqdev/ss/http/dis"
	httpgod "iqdev/ss/http/god"
	httpmanager "iqdev/ss/http/manager"
	httpSchool "iqdev/ss/http/school"
	httpSo "iqdev/ss/http/schoolOwner"
	testuplaod "iqdev/ss/test"

	//"iqdev/ss/libs/sql"
	disWS "iqdev/ss/websocket/dis"
	godWS "iqdev/ss/websocket/god"
	managerWS "iqdev/ss/websocket/manager"
	schoolWS "iqdev/ss/websocket/school"
	schoolOwnerWS "iqdev/ss/websocket/schoolOwner"
	teacherWS "iqdev/ss/websocket/teacher"

	"net/http"
)

func main() {
	// init sql connection
	// https
	httpgod.HttpGodHanlder()
	httpmanager.HttpManagerHanlder()
	httpdis.HttpDisHanlder()
	httpSo.HttpSOHanlder()
	httpSchool.HttpSchoolHanlder()
	// websocket
	godWS.GodWSHander()
	managerWS.ManagerHandler()
	disWS.DisHandler()
	schoolOwnerWS.SchoolOwnerHandler()
	schoolWS.SchoolHandler()
	teacherWS.TeacherHandler()

	http.HandleFunc("/upload", testuplaod.UploadFile)
	// webrtc

	// webview
	http.Handle("/", http.FileServer(http.Dir("./htmx")))
	// server
	fmt.Println("starting the server at port 3944")
	serverErr := http.ListenAndServe(":3944", nil)
	if serverErr != nil {
		fmt.Println("cannot start the server, server error")
	}

	// stoping connections
	// sql connnections
	fmt.Println(" stoping server")
	//sql.ResetConnection()
}
