package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/SystemEngineeringTeam/ryuou-manager/dboperation"
)

func AdminQuestionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		putAdminQuestion(w, r)
	case http.MethodGet:
		getAdminQuestion(w, r)
	}
}

func putAdminQuestion(w http.ResponseWriter, r *http.Request) {
	teamID := strings.Split(r.URL.Path, "/")[3]
	questionID := strings.Split(r.URL.Path, "/")[4]

	numericTeamID, err := strconv.Atoi(teamID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	numericQuestionID, err := strconv.Atoi(questionID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := dboperation.ChangeToCollect(numericTeamID, numericQuestionID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getAdminQuestion(w http.ResponseWriter, r *http.Request) {
	submits, err := dboperation.SelectAllSubmit()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(submits)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(jsonBytes))
}
