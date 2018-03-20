package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/laidingqing/gojobs/accountservice/model"
	"github.com/laidingqing/gojobs/accountservice/storage"
	. "github.com/laidingqing/gojobs/common/controller"
)

//DbStorage instance
var DbStorage storage.Storage

//CreateAccount create account handle.
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, OkResponse{Message: InternalErrorMsg})
		return
	}

	acct := &model.Account{}
	if err := json.Unmarshal(body, acct); err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, OkResponse{Message: InternalErrorMsg})
		return
	}
	rev, err := DbStorage.CreateAccount(*acct)
	if err != nil {
		logrus.Errorf("Some error occured serving : " + err.Error())
		WriteJSONResponse(w, http.StatusBadRequest, OkResponse{Message: "Error create account"})
		return
	}
	WriteJSONResponse(w, http.StatusCreated, OkResponse{Status: "Ok", Message: rev})
}

//GetAccount get account by id
func GetAccount(w http.ResponseWriter, r *http.Request) {
	var accountID = mux.Vars(r)["accountId"]
	account, err := DbStorage.QueryAccount(r.Context(), accountID)
	if err != nil {
		logrus.Errorf("Some error occured serving " + accountID + ": " + err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	account.ServedBy = GetIP()
	WriteJSONResponse(w, http.StatusOK, account)
}

//SeedAccounts seed accounts data.
func SeedAccounts(w http.ResponseWriter, r *http.Request) {
	DbStorage.Seed()
	WriteJSONResponse(w, http.StatusOK, OkResponse{Status: "Ok"})
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
