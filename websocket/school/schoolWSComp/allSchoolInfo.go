package schoolwscomp

import (
	"encoding/json"
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"reflect"

	"github.com/gorilla/websocket"
)

type EventInstideStruct struct {
}

type AllSchoolInfoObj struct {
	SchoolProfiles       globalObject.SchoolsProfile          `json:"school_profile"`
	TeacherProfiles      []globalObject.TeacherProfile        `json:"teacher_profiles"`
	StudentProfiles      []globalObject.StudentProfile        `json:"student_profile"`
	BusProfiles          []globalObject.BusProfile            `json:"bus_profiles"`
	StudyGroupsProfiles  []globalObject.StudyGroupProfile     `json:"study_groups_profiles"`
	ChatGroupProfiles    []globalObject.StudyChatGroupProfile `json:"study_chat_group_profiles"`
	TeacherGroupsProfile []globalObject.TeacherStudyGroup     `json:"teacher_groups"`
	SchedulesProfile     []globalObject.StudyGroupProfile     `json:"schedules_profile"`
	StoreProfile         []globalObject.StoreProfile          `json:"store_profile"`
	EventProfile         []globalObject.SchoolEvents          `json:"event_profile"`
	ExamsProfile         []globalObject.StudentExams          `json:"exams_profile"`
	StudentHistory       []globalObject.StudentHistory        `json:"student_history"`
	StudentPayment       []globalObject.StudentPaymentProfile `json:"student_payments"`
	TeacherPayment       []globalObject.TeacherPaymentProfile `json:"teacher_payments"`
	TeacherStudyGroup    []globalObject.TeacherStudyGroup     `json:"teacher_study_groups"`
}

type ConnectionsList struct {
	GodConnectionNum   int
	GodConnectionsList []*websocket.Conn
}

type connectionsProfile struct {
}

func assembleInfo(SchoolProfile globalObject.SchoolsProfile,
	connection *websocket.Conn) AllSchoolInfoObj {
	allSchoolInfo := AllSchoolInfoObj{}

	if connection == nil {
		return allSchoolInfo
	}

	schoolsInfo := SchoolProfile
	studentInfo := GetAllStudentInfo(schoolsInfo.SchoolId, connection)
	teachersInfo := GetAllTeacherInfo(schoolsInfo.SchoolId, connection)
	studentHIstoryInfo := GetAllStudentHistoryInfo(schoolsInfo.SchoolId, connection)
	eventInfo := GetAllEventsInfo(schoolsInfo.SchoolId, connection)
	examsInfo := GetAllExamsInfo(schoolsInfo.SchoolId, connection)
	studyGroupsInfo := GetAllStudyGroupsInfo(schoolsInfo.SchoolId, connection)
	busInfo := GetAllBusInfo(schoolsInfo.SchoolId, connection)
	studentChatGroupInfo := GetAllStudyChatGroupsInfo(schoolsInfo.SchoolId, connection)
	studentPaymentInfo := GetAllStudentPaymentsInfo(schoolsInfo.SchoolId, connection)
	teacherPaymentInfo := GetAllTeacherPaymentsInfo(schoolsInfo.SchoolId, connection)
	teacherstudyGroupInfo := GetAllTeacherStudyGroupsInfo(schoolsInfo.SchoolId, connection)

	allSchoolInfo.SchoolProfiles = schoolsInfo
	allSchoolInfo.StudentProfiles = studentInfo
	allSchoolInfo.TeacherProfiles = teachersInfo
	allSchoolInfo.ExamsProfile = examsInfo
	allSchoolInfo.EventProfile = eventInfo
	allSchoolInfo.StudentHistory = studentHIstoryInfo
	allSchoolInfo.StudyGroupsProfiles = studyGroupsInfo
	allSchoolInfo.BusProfiles = busInfo
	allSchoolInfo.ChatGroupProfiles = studentChatGroupInfo
	allSchoolInfo.StudentPayment = studentPaymentInfo
	allSchoolInfo.TeacherPayment = teacherPaymentInfo
	allSchoolInfo.TeacherStudyGroup = teacherstudyGroupInfo

	return allSchoolInfo
}

func GetAllSchoolInfo(connection *websocket.Conn,
	schoolProfile globalObject.SchoolsProfile,
	currCounter int, prevCounter int,
	schoolCurrent globalObject.SchoolsProfile,
	schoolPrev globalObject.SchoolsProfile,
	connCounter int) (bool, int, int, globalObject.SchoolsProfile) {

	var allInfoRow = assembleInfo(schoolProfile, connection)

	schoolCurrent = allInfoRow.SchoolProfiles
	compareschool := reflect.DeepEqual(schoolCurrent, schoolPrev)
	fmt.Println("shit")

	currCounter = len(
		allInfoRow.StudentProfiles,
	) + len(
		allInfoRow.TeacherProfiles,
	) + len(
		allInfoRow.ExamsProfile,
	) + len(
		allInfoRow.EventProfile,
	) + len(
		allInfoRow.StudentHistory,
	) + len(
		allInfoRow.StudyGroupsProfiles,
	) + len(
		allInfoRow.BusProfiles,
	) + len(
		allInfoRow.ChatGroupProfiles,
	) + len(
		allInfoRow.StudentPayment,
	) + len(
		allInfoRow.TeacherPayment,
	) + len(
		allInfoRow.TeacherStudyGroup,
	)

	if connCounter > 3 {
		if currCounter == prevCounter && compareschool {
			return false, currCounter, prevCounter, schoolCurrent
		}
	}

	fmt.Println("shit")
	allInfoJson, ctjsonError := json.Marshal(allInfoRow)
	if ctjsonError != nil {
		fmt.Println(ctjsonError)
		libErrors.ReturnError(connection)
	}

	fmt.Println(string(allInfoJson))

	fmt.Println("shit")
	writeError := connection.WriteMessage(1, allInfoJson)
	fmt.Println(writeError)
	if writeError != nil {
		libErrors.ReturnError(connection)
	}
	prevCounter = currCounter
	schoolPrev = schoolCurrent

	fmt.Println("shit")
	return true, currCounter, prevCounter, schoolPrev
}
