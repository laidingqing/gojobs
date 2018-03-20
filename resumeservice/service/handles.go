package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	. "github.com/laidingqing/gojobs/common/controller"
	"github.com/laidingqing/gojobs/resumeservice/model"
	"github.com/laidingqing/gojobs/resumeservice/storage"
)

//DbStorage instance
var DbStorage storage.Storage

//CreateResume create account resume handle.
func CreateResume(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, OkResponse{Message: InternalErrorMsg})
		return
	}

	resume := &model.Resume{}
	if err := json.Unmarshal(body, resume); err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, OkResponse{Message: InternalErrorMsg})
		return
	}
	rev, err := DbStorage.CreateResume(*resume)
	if err != nil {
		logrus.Errorf("Some error occured serving : " + err.Error())
		WriteJSONResponse(w, http.StatusBadRequest, OkResponse{Message: "Error create Resume"})
		return
	}
	WriteJSONResponse(w, http.StatusCreated, OkResponse{Status: "Ok", Message: rev})
}

//GetResume get account's resume by id
func GetResume(w http.ResponseWriter, r *http.Request) {
	var accountID = mux.Vars(r)["accountId"]
	resume, err := DbStorage.FindResume(accountID)
	if err != nil {
		logrus.Errorf("Some error occured serving " + accountID + ": " + err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	WriteJSONResponse(w, http.StatusOK, resume)
}

//UpdateBasicResume ...
func UpdateBasicResume(w http.ResponseWriter, r *http.Request) {
	var accountID = mux.Vars(r)["accountId"]
	resume := &model.Resume{}
	requestEntryBody(w, r, resume)
	err := DbStorage.UpdateBaseResume(resume.Name, resume.Bio, resume.Intentions, accountID)
	if err != nil {
		logrus.Errorf("Some error occured serving " + accountID + ": " + err.Error())
		WriteJSONResponse(w, http.StatusBadRequest, OkResponse{Message: "Error update basic Resume info"})
		return
	}
}

//UpdateWorkResume ...
func UpdateWorkResume(w http.ResponseWriter, r *http.Request) {
	var accountID = mux.Vars(r)["accountId"]
	work := &model.Work{}
	requestEntryBody(w, r, work)
	err := DbStorage.UpdateWorkResume(*work, accountID)
	if err != nil {
		logrus.Errorf("Some error occured serving " + accountID + ": " + err.Error())
		WriteJSONResponse(w, http.StatusBadRequest, OkResponse{Message: "Error update basic Resume info"})
		return
	}
}

//UpdateEducationResume ...
func UpdateEducationResume(w http.ResponseWriter, r *http.Request) {
	var accountID = mux.Vars(r)["accountId"]
	edu := &model.Education{}
	requestEntryBody(w, r, edu)
	err := DbStorage.UpdateEducationResume(*edu, accountID)
	if err != nil {
		logrus.Errorf("Some error occured serving " + accountID + ": " + err.Error())
		WriteJSONResponse(w, http.StatusBadRequest, OkResponse{Message: "Error update education Resume info"})
		return
	}
}

//UpdateProjectResume ...
func UpdateProjectResume(w http.ResponseWriter, r *http.Request) {
	var accountID = mux.Vars(r)["accountId"]
	project := &model.Project{}
	requestEntryBody(w, r, project)
	err := DbStorage.UpdateProjectResume(*project, accountID)
	if err != nil {
		logrus.Errorf("Some error occured serving " + accountID + ": " + err.Error())
		WriteJSONResponse(w, http.StatusBadRequest, OkResponse{Message: "Error update project Resume info"})
		return
	}
}

//HealthCheck health for service check.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	dbUp := DbStorage.Check()
	if dbUp {
		WriteJSONResponse(w, http.StatusOK, OkResponse{Status: "UP"})
	} else {
		WriteJSONResponse(w, http.StatusServiceUnavailable, OkResponse{Status: "Database unaccessible"})
	}
}

func requestEntryBody(w http.ResponseWriter, r *http.Request, model interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, OkResponse{Message: InternalErrorMsg})
		return
	}
	if err := json.Unmarshal(body, &model); err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, OkResponse{Message: InternalErrorMsg})
		return
	}
}
